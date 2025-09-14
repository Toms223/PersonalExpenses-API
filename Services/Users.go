package Services

import (
	"PersonalExpensesAPI/Model/App"
	"PersonalExpensesAPI/Repositories"
	"errors"

	"github.com/jackc/pgx/v5"
)

type UsersService struct {
	repo *Repositories.UsersRepo
}

func (service UsersService) CreateNewUser(name string, email string) (*App.User, error) {
	_, err := service.repo.GetUserByEmail(email)
	if !errors.Is(err, pgx.ErrNoRows) {
		return nil, UserError.UserAlreadyExists()
	}
	user := App.User{
		Name:  name,
		Email: email,
		Limit: 0,
	}
	err = user.Create(service.repo.Ctx, service.repo.DB)
	if err != nil {
		return nil, UserError.CouldNotCreateUser()
	}
	return &user, nil
}

func (service UsersService) GetUserById(id int) (*App.User, error) {
	user, err := service.repo.GetUserById(id)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, UserError.UserNotFound()
	}
	return user, nil
}

func (service UsersService) GetUserByEmail(email string) (*App.User, error) {
	user, err := service.repo.GetUserByEmail(email)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, UserError.UserNotFound()
	}
	return user, nil
}

func (service UsersService) ChangeUserLimit(id int, limit int) (*App.User, error) {
	if limit < 0 {
		return nil, UserError.LimitMustBeValid()
	}
	user, err := service.GetUserById(id)
	if err != nil {
		return nil, err
	}
	user.Limit = limit
	err = user.Save(service.repo.Ctx, service.repo.DB)
	if err != nil {
		return nil, UserError.CouldNotUpdateUser()
	}
	return user, nil
}

func (service UsersService) ChangeUserName(id int, name string) (*App.User, error) {
	if name == "" {
		return nil, UserError.NameMustBeValid()
	}
	user, err := service.GetUserById(id)
	if err != nil {
		return nil, err
	}
	user.Name = name
	err = user.Save(service.repo.Ctx, service.repo.DB)
	if err != nil {
		return nil, UserError.CouldNotUpdateUser()
	}
	return user, nil
}

func (service UsersService) DeleteUser(id int) error {
	user, err := service.GetUserById(id)
	if err != nil {
		return err
	}
	err = user.Delete(service.repo.Ctx, service.repo.DB)
	if err != nil {
		return UserError.CouldNotDeleteUser()
	}
	return nil
}

func (service UsersService) GetUsers(skip int, limit int) ([]App.User, error) {
	if skip < 0 || limit <= 0 {
		return nil, UserError.InvalidPagination()
	}
	users, err := service.repo.GetUsers(skip, limit)
	if err != nil {
		return nil, UserError.CouldNotGetUsers()
	}
	return users, nil
}

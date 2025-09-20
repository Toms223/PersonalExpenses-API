package Services

import (
	"PersonalExpensesAPI/Errors"
	"PersonalExpensesAPI/Errors/UserError"
	"PersonalExpensesAPI/Model/App"
	"PersonalExpensesAPI/Repositories"
	"errors"

	"github.com/jackc/pgx/v5"
)

type UsersService struct {
	repo *Repositories.UsersRepo
}

func (service UsersService) CreateNewUser(name string, email string) (*App.User, *Errors.APIError) {
	if email == "" {
		return nil, &UserError.InvalidEmail
	}
	if name == "" {
		return nil, &UserError.InvalidName
	}
	_, err := service.repo.GetUserByEmail(email)
	if !errors.Is(err, pgx.ErrNoRows) {
		return nil, &UserError.UserAlreadyExists
	}
	user := App.User{
		Name:  name,
		Email: email,
		Limit: 0,
	}
	err = user.Create(service.repo.Ctx, service.repo.DB)
	if err != nil {
		return nil, &UserError.CouldNotCreateUser
	}
	return &user, nil
}

func (service UsersService) GetUserById(id int) (*App.User, *Errors.APIError) {
	if id < 1 {
		return nil, &UserError.InvalidUserId
	}
	user, err := service.repo.GetUserById(id)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, &UserError.UserNotFound
	}
	return user, nil
}

func (service UsersService) GetUserByEmail(email string) (*App.User, *Errors.APIError) {
	if email == "" {
		return nil, &UserError.InvalidEmail
	}
	user, err := service.repo.GetUserByEmail(email)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, &UserError.UserNotFound
	}
	return user, nil
}

func (service UsersService) ChangeUserLimit(id int, limit int) (*App.User, *Errors.APIError) {
	if limit < 0 {
		return nil, &UserError.InvalidLimit
	}
	user, userError := service.GetUserById(id)
	if userError != nil {
		return nil, userError
	}
	user.Limit = limit
	err := user.Save(service.repo.Ctx, service.repo.DB)
	if err != nil {
		return nil, &UserError.CouldNotUpdateUser
	}
	return user, nil
}

func (service UsersService) ChangeUserName(id int, name string) (*App.User, *Errors.APIError) {
	if name == "" {
		return nil, &UserError.InvalidName
	}
	user, userError := service.GetUserById(id)
	if userError != nil {
		return nil, userError
	}
	user.Name = name
	err := user.Save(service.repo.Ctx, service.repo.DB)
	if err != nil {
		return nil, &UserError.CouldNotUpdateUser
	}
	return user, nil
}

func (service UsersService) DeleteUser(id int) *Errors.APIError {
	user, userError := service.GetUserById(id)
	if userError != nil {
		return userError
	}
	err := user.Delete(service.repo.Ctx, service.repo.DB)
	if err != nil {
		return &UserError.CouldNotDeleteUser
	}
	return nil
}

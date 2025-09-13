package Repositories

import (
	"PersonalExpensesAPI/Model/App"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UsersRepo struct {
	Ctx context.Context
	DB  *pgxpool.Pool
}

func (repo UsersRepo) GetUserById(id int) (*App.User, error) {
	user := App.User{}
	err := repo.DB.QueryRow(repo.Ctx, `SELECT ("Id", "Name","Email","Limit") FROM "Users" WHERE "Id" = $1;`,
		id).Scan(&user.Id, &user.Name, &user.Email, &user.Limit)
	if err != nil {
		return nil, fmt.Errorf("error getting user by id: %s", err.Error())
	}
	return &user, nil
}

func (repo UsersRepo) GetUserByEmail(email string) (*App.User, error) {
	user := App.User{}
	err := repo.DB.QueryRow(repo.Ctx, `SELECT ("Id", "Name","Email","Limit") FROM "Users" WHERE "Email" = $1;`,
		email).Scan(&user.Id, &user.Name, &user.Email, &user.Limit)
	if err != nil {
		return nil, fmt.Errorf("error getting user by id: %s", err.Error())
	}
	return &user, nil
}

func (repo UsersRepo) GetUsers(skip int, limit int) ([]App.User, error) {
	if skip < 0 || limit <= 0 || skip > limit {
		return nil, fmt.Errorf("skip must be less than limit and positive. Limit must be greater than 0")
	}
	users := make([]App.User, limit-skip)
	rows, err := repo.DB.Query(repo.Ctx,
		`SELECT ("Id", "Name","Email","Limit") FROM "Users" ORDER BY "Name" LIMIT $1 OFFSET $2;`,
		limit, skip)
	if err != nil {
		return nil, fmt.Errorf("error getting users: %s", err.Error())
	}
	for rows.Next() {
		var user App.User
		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Limit)
		if err != nil {
			return nil, fmt.Errorf("error getting users: %s", err.Error())
		}
		users = append(users, user)
	}
	return users, nil
}

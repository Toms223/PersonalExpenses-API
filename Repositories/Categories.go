package Repositories

import (
	"PersonalExpensesAPI/Model/App"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type CategoriesRepo struct {
	Ctx context.Context
	DB  *pgxpool.Pool
}

func (repo CategoriesRepo) GetUserCategoryById(id int, userId int) (*App.Category, error) {
	var category App.Category
	err := repo.DB.QueryRow(repo.Ctx, `SELECT ("Id","UserId","Name","Color") FROM "Categories" WHERE "Id" = $1 AND "UserId" = $2`,
		id, userId).Scan(&category.Id, &category.UserId, &category.Name)
	if err != nil {
		return nil, fmt.Errorf("error geting user category by id: %s", err)
	}
	return &category, nil
}

func (repo CategoriesRepo) GetUserCategories(userId int, skip int, limit int) ([]App.Category, error) {
	categories := make([]App.Category, 0)
	rows, err := repo.DB.Query(repo.Ctx,
		`SELECT ("Id","UserId","Name","Color") FROM "Categories" WHERE "UserId" = $1 OFFSET $2 LIMIT $3`,
		userId, skip, limit)
	if err != nil {
		return nil, fmt.Errorf("error getting user categories: %s", err)
	}
	defer rows.Close()
	for rows.Next() {
		var category App.Category
		err := rows.Scan(&category.Id, &category.UserId, &category.Name, &category.Color)
		if err != nil {
			return nil, fmt.Errorf("error getting user categories: %s", err)
		}
		categories = append(categories, category)
	}
	return categories, nil
}

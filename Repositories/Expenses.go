package Repositories

import (
	"PersonalExpensesAPI/Model/App"
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ExpensesReop struct {
	Ctx context.Context
	DB  *pgxpool.Pool
}

func (repo ExpensesReop) GetUserExpenseById(userId int, id int) (*App.Expense, error) {
	var expense App.Expense
	err := repo.DB.QueryRow(repo.Ctx,
		`SELECT ("Id", "Name", "Amount", "Date", "UserId", "CategoryId","Continuous","Fixed","Period") 
			FROM "Expenses" WHERE "Id" = $1 AND "UserId" = $2`, id, userId).Scan(
		&expense.Id,
		&expense.Name,
		&expense.Amount,
		&expense.Date,
		&expense.UserId,
		&expense.CategoryId,
		&expense.Continuous,
		&expense.Fixed,
		&expense.Period)
	if err != nil {
		return nil, fmt.Errorf("error getting user expense by id: %s", err)
	}
	return &expense, nil
}

func (repo ExpensesReop) GetUserExpenses(userId int, skip int, limit int) ([]App.Expense, error) {
	if skip < 0 || limit <= 0 || skip > limit {
		return nil, fmt.Errorf("skip must be less than limit and positive. Limit must be greater than 0")
	}
	rows, err := repo.DB.Query(repo.Ctx,
		`SELECT ("Id", "Name", "Amount", "Date", "UserId", "CategoryId","Continuous","Fixed","Period") 
			FROM "Expenses" WHERE "UserId" = $1 ORDER BY "Date" LIMIT $2 OFFSET $3;`,
		userId, limit, skip)
	if err != nil {
		return nil, fmt.Errorf("error getting expenses: %s", err.Error())
	}
	return rowsToExpensesSlice(rows)
}

func (repo ExpensesReop) GetUserExpensesByDate(userId int, startDate time.Time, endDate time.Time) ([]App.Expense, error) {
	if startDate.Unix() > endDate.Unix() {
		return nil, fmt.Errorf("start date must come before end date")
	}
	rows, err := repo.DB.Query(repo.Ctx,
		`SELECT ("Id", "Name", "Amount", "Date", "UserId", "CategoryId","Continuous","Fixed","Period") 
			FROM "Expenses" WHERE "UserId" = $1 AND "Date" >= $2 AND "Date" <= $3;`,
		userId, startDate, endDate)
	if err != nil {
		return nil, fmt.Errorf("error getting expenses: %s", err.Error())
	}
	return rowsToExpensesSlice(rows)
}

func (repo ExpensesReop) GetYearMonthExpenses(userId int, year int, month int) ([]App.Expense, error) {
	if month < 1 || month > 12 || year < 0 {
		return nil, fmt.Errorf("month value must be between 1 and 12 and year must be positive")
	}
	startDate := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	endDate := startDate.AddDate(0, 1, 0)
	rows, err := repo.DB.Query(repo.Ctx,
		`SELECT ("Id", "Name", "Amount", "Date", "UserId", "CategoryId","Continuous","Fixed","Period") 
			FROM "Expenses" WHERE "UserId" = $1 AND "Date" >= $2 AND "Date" < $3 ;`,
		userId, startDate, endDate)
	if err != nil {
		return nil, fmt.Errorf("error getting expenses: %s", err.Error())
	}
	return rowsToExpensesSlice(rows)
}

func (repo ExpensesReop) GetExpensesByCategoryId(userId int, categoryId int, limit int, skip int) ([]App.Expense, error) {
	if skip < 0 || limit <= 0 || skip > limit {
		return nil, fmt.Errorf("skip must be less than limit and positive. Limit must be greater than 0")
	}
	rows, err := repo.DB.Query(repo.Ctx,
		`SELECT ("Id", "Name", "Amount", "Date", "UserId", "CategoryId","Continuous","Fixed","Period") 
			FROM "Expenses" WHERE "UserId" = $1 AND "CategoryId" = $2 ORDER BY "Date" LIMIT $3 OFFSET $4;`,
		userId, categoryId, limit, skip)
	if err != nil {
		return nil, fmt.Errorf("error getting expenses: %s", err.Error())
	}
	return rowsToExpensesSlice(rows)
}

func (repo ExpensesReop) GetExpensesByCategoryIdAndDate(userId int, categoryId int, startDate time.Time, endDate time.Time) ([]App.Expense, error) {
	if startDate.Unix() > endDate.Unix() {
		return nil, fmt.Errorf("start date must come before end date")
	}
	rows, err := repo.DB.Query(repo.Ctx,
		`SELECT ("Id", "Name", "Amount", "Date", "UserId", "CategoryId","Continuous","Fixed","Period") 
			FROM "Expenses" WHERE "UserId" = $1 AND "CategoryId" = $2 AND "Date" >= $3 AND "Date" <= $4 ;`,
		userId, categoryId, startDate, endDate)
	if err != nil {
		return nil, fmt.Errorf("error getting expenses: %s", err.Error())
	}
	return rowsToExpensesSlice(rows)
}

func (repo ExpensesReop) GetExpensesByCategoryIdAndYearMonth(userId int, categoryId int, year int, month int) ([]App.Expense, error) {
	if month < 1 || month > 12 || year < 0 {
		return nil, fmt.Errorf("month value must be between 1 and 12 and year must be positive")
	}
	startDate := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	endDate := startDate.AddDate(0, 1, 0)
	rows, err := repo.DB.Query(repo.Ctx,
		`SELECT ("Id", "Name", "Amount", "Date", "UserId", "CategoryId","Continuous","Fixed","Period") 
			FROM "Expenses" WHERE "UserId" = $1 AND "CategoryId" = $2 AND "Date" >= $3 AND "Date" < $4 ;`,
		userId, categoryId, startDate, endDate)
	if err != nil {
		return nil, fmt.Errorf("error getting expenses: %s", err.Error())
	}
	return rowsToExpensesSlice(rows)
}

func rowsToExpensesSlice(rows pgx.Rows) ([]App.Expense, error) {
	expenses := make([]App.Expense, 0)
	defer rows.Close()
	for rows.Next() {
		var expense App.Expense
		err := rows.Scan(
			&expense.Id,
			&expense.Name,
			&expense.Amount,
			&expense.Date,
			&expense.UserId,
			&expense.CategoryId,
			&expense.Continuous,
			&expense.Fixed,
			&expense.Period)
		if err != nil {
			return nil, fmt.Errorf("error getting expenses: %s", err.Error())
		}
		expenses = append(expenses, expense)
	}
	return expenses, nil
}

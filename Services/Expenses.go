package Services

import (
	"PersonalExpensesAPI/Errors"
	"PersonalExpensesAPI/Errors/ExpensesError"
	"PersonalExpensesAPI/Model/App"
	"PersonalExpensesAPI/Repositories"
	"time"
)

type ExpensesService struct {
	repo *Repositories.ExpensesRepo
}

func (service ExpensesService) CreateNewExpense(
	name string,
	amount float32,
	date time.Time,
	userId int,
	categoryId int,
	continuous bool,
	fixed bool,
	period int,
) (*App.Expense, *Errors.APIError) {
	expenseError := verifyValidExpense(name, amount, userId, categoryId, continuous, fixed, period)
	if expenseError != nil {
		return nil, expenseError
	}
	expense := &App.Expense{
		Name:       name,
		Amount:     amount,
		Date:       date,
		UserId:     userId,
		CategoryId: categoryId,
		Continuous: continuous,
		Fixed:      fixed,
		Period:     period,
	}
	err := expense.Create(service.repo.Ctx, service.repo.DB)
	if err != nil {
		return nil, &ExpensesError.CouldNotCreateExpense
	}
	return expense, nil
}

func (service ExpensesService) GetExpenseById(userId int, id int) (*App.Expense, *Errors.APIError) {
	if userId < 1 {
		return nil, ExpensesError.InvalidField("userId")
	}
	if id < 1 {
		return nil, ExpensesError.InvalidField("id")
	}
	expense, err := service.repo.GetUserExpenseById(userId, id)
	if err != nil {
		return nil, &ExpensesError.ExpenseNotFound
	}
	return expense, nil
}

func (service ExpensesService) GetExpensesPaginated(userId int, skip int, limit int) ([]App.Expense, *Errors.APIError) {
	if userId < 1 {
		return nil, ExpensesError.InvalidField("userId")
	}
	expenses, err := service.repo.GetUserExpenses(userId, skip, limit)
	if err != nil {
		return nil, &ExpensesError.CouldNotRetrieveExpenses
	}
	return expenses, nil
}

func (service ExpensesService) GetExpensesByDateInterval(userId int, startDate time.Time, endDate time.Time) ([]App.Expense, *Errors.APIError) {
	if userId < 1 {
		return nil, ExpensesError.InvalidField("userId")
	}
	expenses, err := service.repo.GetUserExpensesByDate(userId, startDate, endDate)
	if err != nil {
		return nil, &ExpensesError.CouldNotRetrieveExpenses
	}
	return expenses, nil
}

func (service ExpensesService) GetExpensesByMonthAndYear(userId int, month int, year int) ([]App.Expense, *Errors.APIError) {
	if userId < 1 {
		return nil, ExpensesError.InvalidField("userId")
	}
	expenses, err := service.repo.GetUserExpensesByMonthAndYear(userId, month, year)
	if err != nil {
		return nil, &ExpensesError.CouldNotRetrieveExpenses
	}
	return expenses, nil
}

func (service ExpensesService) GetCategoryExpensesPaginated(userId int, categoryId int, skip int, limit int) ([]App.Expense, *Errors.APIError) {
	if userId < 1 {
		return nil, ExpensesError.InvalidField("userId")
	}
	if categoryId < 1 {
		return nil, ExpensesError.InvalidField("categoryId")
	}
	expenses, err := service.repo.GetUserExpensesByCategoryId(userId, categoryId, skip, limit)
	if err != nil {
		return nil, &ExpensesError.CouldNotRetrieveExpenses
	}
	return expenses, nil
}

func (service ExpensesService) GetCategoryExpensesByDateInterval(userId int, categoryId int, startDate time.Time, endDate time.Time) ([]App.Expense, *Errors.APIError) {
	if userId < 1 {
		return nil, ExpensesError.InvalidField("userId")
	}
	if categoryId < 1 {
		return nil, ExpensesError.InvalidField("categoryId")
	}
	expenses, err := service.repo.GetUserExpensesByCategoryIdAndDate(userId, categoryId, startDate, endDate)
	if err != nil {
		return nil, &ExpensesError.CouldNotRetrieveExpenses
	}
	return expenses, nil
}

func (service ExpensesService) GetCategoryExpensesByMonthAndYear(userId int, categoryId int, month int, year int) ([]App.Expense, *Errors.APIError) {
	if userId < 1 {
		return nil, ExpensesError.InvalidField("userId")
	}
	if categoryId < 1 {
		return nil, ExpensesError.InvalidField("categoryId")
	}
	expenses, err := service.repo.GetUserExpensesByCategoryIdAndMonthAndYear(userId, categoryId, month, year)
	if err != nil {
		return nil, &ExpensesError.CouldNotRetrieveExpenses
	}
	return expenses, nil
}

func (service ExpensesService) UpdateExpense(
	id int,
	name string,
	amount float32,
	date time.Time,
	userId int,
	categoryId int,
	continuous bool,
	fixed bool,
	period int,
) (*App.Expense, *Errors.APIError) {
	expense, expenseError := service.GetExpenseById(userId, id)
	if expenseError != nil {
		return nil, &ExpensesError.ExpenseNotFound
	}
	expenseError = verifyValidExpense(name, amount, userId, categoryId, continuous, fixed, period)
	if expenseError != nil {
		return nil, expenseError
	}
	expense.Name = name
	expense.Amount = amount
	expense.Date = date
	expense.CategoryId = categoryId
	expense.Continuous = continuous
	expense.Fixed = fixed
	expense.Period = period
	err := expense.Save(service.repo.Ctx, service.repo.DB)
	if err != nil {
		return nil, &ExpensesError.CouldNotUpdateExpense
	}
	return expense, nil
}

func (service ExpensesService) DeleteExpense(id int, userId int) *Errors.APIError {
	expense, expenseError := service.GetExpenseById(userId, id)
	if expenseError != nil {
		return &ExpensesError.ExpenseNotFound
	}
	err := expense.Delete(service.repo.Ctx, service.repo.DB)
	if err != nil {
		return &ExpensesError.CouldNotDeleteExpense
	}
	return nil
}

func verifyValidExpense(
	name string,
	amount float32,
	userId int,
	categoryId int,
	continuous bool,
	fixed bool,
	period int) *Errors.APIError {
	if name == "" {
		return ExpensesError.InvalidField("name")
	}
	if amount < 0 {
		return ExpensesError.InvalidField("amount")
	}
	if userId < 1 {
		return ExpensesError.InvalidField("userId")
	}
	if categoryId < 1 {
		return ExpensesError.InvalidField("categoryId")
	}
	if continuous && fixed && period < 1 {
		return ExpensesError.InvalidField("period")
	}
	return nil
}

package ExpensesError

import "PersonalExpensesAPI/Errors"

var (
	CouldNotCreateExpense = Errors.APIError{
		Message: "Could not create expense",
		Code:    500,
	}

	ExpenseNotFound = Errors.APIError{
		Message: "Expense not found",
		Code:    404,
	}

	CouldNotRetrieveExpenses = Errors.APIError{
		Message: "Could not retrieve expenses",
		Code:    500,
	}

	CouldNotUpdateExpense = Errors.APIError{
		Message: "Could not update expense",
		Code:    500,
	}

	CouldNotDeleteExpense = Errors.APIError{
		Message: "Could not delete expense",
		Code:    500,
	}
)

func InvalidField(field string) *Errors.APIError {
	return &Errors.APIError{
		Message: "Invalid " + field,
		Code:    400,
	}
}

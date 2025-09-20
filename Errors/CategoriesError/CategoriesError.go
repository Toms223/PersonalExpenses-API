package CategoriesError

import "PersonalExpensesAPI/Errors"

var (
	InvalidUserId = Errors.APIError{
		Message: "Invalid user id",
		Code:    400,
	}

	InvalidName = Errors.APIError{
		Message: "Invalid name",
		Code:    400,
	}

	InvalidColor = Errors.APIError{
		Message: "Invalid color",
		Code:    400,
	}

	CouldNotCreateCategory = Errors.APIError{
		Message: "Could not create category",
		Code:    500,
	}

	InvalidCategoryId = Errors.APIError{
		Message: "Invalid category id",
		Code:    400,
	}

	CouldNotGetCategory = Errors.APIError{
		Message: "Could not get category",
		Code:    500,
	}

	CouldNotUpdateCategory = Errors.APIError{
		Message: "Could not update category",
		Code:    500,
	}
)

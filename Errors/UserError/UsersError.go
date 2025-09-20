package UserError

import "PersonalExpensesAPI/Errors"

var (
	InvalidEmail = Errors.APIError{
		Message: "Invalid email",
		Code:    400,
	}

	InvalidName = Errors.APIError{
		Message: "Invalid name",
		Code:    400,
	}

	UserAlreadyExists = Errors.APIError{
		Message: "User already exists",
		Code:    409,
	}

	CouldNotCreateUser = Errors.APIError{
		Message: "Could not create user",
		Code:    500,
	}

	InvalidUserId = Errors.APIError{
		Message: "Invalid user id",
		Code:    400,
	}

	UserNotFound = Errors.APIError{
		Message: "User not found",
		Code:    404,
	}

	InvalidLimit = Errors.APIError{
		Message: "Invalid limit",
		Code:    400,
	}

	CouldNotUpdateUser = Errors.APIError{
		Message: "Could not update user",
		Code:    500,
	}

	CouldNotDeleteUser = Errors.APIError{
		Message: "Could not delete user",
		Code:    500,
	}
)

package Services

import (
	"PersonalExpensesAPI/Errors"
	"PersonalExpensesAPI/Errors/CategoriesError"
	"PersonalExpensesAPI/Model/App"
	"PersonalExpensesAPI/Repositories"
	"strings"
)

type CategoriesService struct {
	repo *Repositories.CategoriesRepo
}

func (service CategoriesService) Create(userId int, name string, color string) (*App.Category, *Errors.APIError) {
	if userId < 1 {
		return nil, &CategoriesError.InvalidUserId
	}
	if name == "" {
		return nil, &CategoriesError.InvalidName
	}
	if invalidColor(strings.ToUpper(color)) {
		return nil, &CategoriesError.InvalidColor
	}
	category := App.Category{
		UserId: userId,
		Name:   name,
		Color:  color,
	}
	err := category.Create(service.repo.Ctx, service.repo.DB)
	if err != nil {
		return nil, &CategoriesError.CouldNotCreateCategory
	}
	return &category, nil
}

func (service CategoriesService) GetCategoryById(id int, userId int) (*App.Category, *Errors.APIError) {
	if userId < 1 {
		return nil, &CategoriesError.InvalidUserId
	}
	if id < 1 {
		return nil, &CategoriesError.InvalidCategoryId
	}
	category, err := service.repo.GetUserCategoryById(id, userId)
	if err != nil {
		return nil, &CategoriesError.CouldNotGetCategory
	}
	return category, nil
}

func (service CategoriesService) GetCategoryUserCategories(userId int, skip int, limit int) ([]App.Category, *Errors.APIError) {
	if userId < 1 {
		return nil, &CategoriesError.InvalidUserId
	}
	categories, err := service.repo.GetUserCategories(userId, skip, limit)
	if err != nil {
		return nil, &CategoriesError.CouldNotGetCategory
	}
	return categories, nil
}

func (service CategoriesService) UpdateCategory(id int, userId int, name string, color string) (*App.Category, *Errors.APIError) {
	if userId < 1 {
		return nil, &CategoriesError.InvalidUserId
	}
	if name == "" {
		return nil, &CategoriesError.InvalidName
	}
	if invalidColor(strings.ToUpper(color)) {
		return nil, &CategoriesError.InvalidColor
	}
	category, categoryError := service.GetCategoryById(id, userId)
	if categoryError != nil {
		return nil, &CategoriesError.CouldNotGetCategory
	}
	category.Name = name
	category.Color = color
	err := category.Save(service.repo.Ctx, service.repo.DB)
	if err != nil {
		return nil, &CategoriesError.CouldNotUpdateCategory
	}
	return category, nil
}

func invalidColor(color string) bool {
	if len(color) != 6 {
		return false
	}
	for _, c := range color {
		if !((c >= '0' && c <= '9') ||
			(c >= 'A' && c <= 'F')) {
			return false
		}
	}
	return true
}

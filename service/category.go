package service

import (
	"gostore/entity"
	"gostore/repo"
)

type CategoryService struct {
	Repo repo.CategoryRepository
}

func NewCategoryService(r repo.CategoryRepository) *CategoryService {
	return &CategoryService{
		Repo: r,
	}
}

func (c *CategoryService) GetAllCategories() ([]entity.Category, error) {
	result, err := c.Repo.GetAllCategories()
	return result, err
}

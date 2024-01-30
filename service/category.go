package service

import (
	"context"
	"gostore/entity"
	"gostore/repo"
	categoryError "gostore/utils/helper/domain/errorModel"
)

type categoryService struct {
	Repo repo.CategoryRepository
}

type CategoryService interface {
	GetAllCategories(ctx context.Context) ([]entity.Category, error)
	GetCategoryById(ctx context.Context, id int) (entity.Category, error)
	InsertCategory(ctx context.Context, category *entity.Category) error
	UpdateCategory(ctx context.Context, id int, category *entity.Category) error
	DeleteCategory(ctx context.Context, id int) error
}

func NewCategoryService(r repo.CategoryRepository) CategoryService {
	return &categoryService{
		Repo: r,
	}
}

func (c *categoryService) GetAllCategories(ctx context.Context) ([]entity.Category, error) {
	return c.Repo.GetAllCategories(ctx)
}

func (c *categoryService) GetCategoryById(ctx context.Context, id int) (entity.Category, error) {
	return c.Repo.GetCategoryById(ctx, id)
}

func (c *categoryService) InsertCategory(ctx context.Context, category *entity.Category) error {
	var (
		detailError = make(map[string]any)
	)

	if category.Name == "" {
		detailError["name"] = "This field is missing input"
	}

	if len(detailError) > 0 {
		return categoryError.ErrCategoryInput.AttachDetail(detailError)
	}

	return c.Repo.InsertCategory(ctx, category)
}

func (c *categoryService) UpdateCategory(ctx context.Context, id int, category *entity.Category) error {
	return c.Repo.UpdateCategory(ctx, id, category)
}

func (c *categoryService) DeleteCategory(ctx context.Context, id int) error {
	return c.Repo.DeleteCategory(ctx, id)
}

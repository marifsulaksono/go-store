package repo

import (
	"context"
	"errors"
	"gostore/entity"
	categoryError "gostore/utils/helper/domain/errorModel"

	"gorm.io/gorm"
)

type categoryRepository struct {
	DB *gorm.DB
}

type CategoryRepository interface {
	GetAllCategories(ctx context.Context) ([]entity.Category, error)
	GetCategoryById(ctx context.Context, id int) (entity.Category, error)
	InsertCategory(ctx context.Context, category *entity.Category) error
	UpdateCategory(ctx context.Context, id int, category *entity.Category) error
	DeleteCategory(ctx context.Context, id int) error
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{
		DB: db,
	}
}

func (c *categoryRepository) GetAllCategories(ctx context.Context) ([]entity.Category, error) {
	var result []entity.Category
	err := c.DB.Find(&result).Error
	return result, err
}

func (c *categoryRepository) GetCategoryById(ctx context.Context, id int) (entity.Category, error) {
	var result entity.Category
	err := c.DB.Where("id = ?", id).First(&result).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Category{}, categoryError.ErrCategoryNotFound
		}
		return entity.Category{}, err
	}
	return result, err
}

func (c *categoryRepository) InsertCategory(ctx context.Context, category *entity.Category) error {
	return c.DB.Create(category).Error
}

func (c *categoryRepository) UpdateCategory(ctx context.Context, id int, category *entity.Category) error {
	return c.DB.Model(&entity.Category{}).Where("id = ?", id).Updates(category).Error
}

func (c *categoryRepository) DeleteCategory(ctx context.Context, id int) error {
	return c.DB.Where("id = ?", id).Delete(&entity.Category{}).Error
}

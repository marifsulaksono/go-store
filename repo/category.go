package repo

import (
	"gostore/entity"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{
		DB: db,
	}
}

func (c *CategoryRepository) GetAllCategories() ([]entity.Category, error) {
	var result []entity.Category
	err := c.DB.Find(&result).Error
	return result, err
}

func (c *CategoryRepository) GetCategoryById(id int) (entity.Category, error) {
	var result entity.Category
	err := c.DB.Where("category_id = ?", id).First(&result).Error
	return result, err
}

// err := config.DB.Where("category_id = ?", id).Preload("Category", "id NOT IN (?)", "cancelled").Find(&result).Error

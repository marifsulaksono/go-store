package repo

import (
	"gostore/entity"

	"gorm.io/gorm"
)

type ItemRepository struct {
	DB *gorm.DB
}

func NewItemRepository(db *gorm.DB) *ItemRepository {
	return &ItemRepository{
		DB: db,
	}
}

func (i *ItemRepository) GetAllItems() ([]entity.ItemResponse, error) {
	var items []entity.ItemResponse
	err := i.DB.Preload("Category").Find(&items).Error
	return items, err
}

func (i *ItemRepository) GetItemById(id int) (entity.ItemResponse, error) {
	var item entity.ItemResponse
	err := i.DB.Where("id = ?", id).Preload("Category", "id NOT IN (?)", "cancelled").First(&item).Error
	return item, err
}

func (i *ItemRepository) GetItemByStatus(s int) ([]entity.ItemResponse, error) {
	var items []entity.ItemResponse
	var err error
	if s == 1 {
		err = i.DB.Where("is_sale = ?", s).Preload("Category", "id NOT IN (?)", "cancelled").Find(&items).Error
	} else if s == 0 {
		err = i.DB.Unscoped().Where("is_sale = ? AND delete_at = null", s).Preload("Category", "id NOT IN (?)", "cancelled").Find(&items).Error
	}
	return items, err
}

func (i *ItemRepository) SearchItem(keyword, order, sortBy string, minPrice, maxPrice, limit, offset int) ([]entity.ItemResponse, error) {
	var items []entity.ItemResponse
	err := i.DB.Where("name LIKE ? AND price BETWEEN ? AND ?", "%"+keyword+"%", minPrice, maxPrice).Order(sortBy +
		" " + order).Limit(limit).Offset(offset).Preload("Category").Find(&items).Error
	return items, err
}

func (i *ItemRepository) InsertItem(item *entity.Item) error {
	err := i.DB.Create(item).Error
	return err
}

func (i *ItemRepository) UpdateItem(id int, item *entity.Item) error {
	err := i.DB.Model(&entity.Item{}).Where("id = ?", id).Updates(map[string]interface{}{
		"name":        item.Name,
		"stock":       item.Stock,
		"price":       item.Price,
		"is_sale":     item.IsSale,
		"category_id": item.CategoryId,
	}).Error

	return err
}

func (i *ItemRepository) ChangeStatusItem(id, s int) error {
	err := i.DB.Model(&entity.Item{}).Where("id = ?", id).Update("is_sale = ?", s).Error
	return err
}

func (i *ItemRepository) SoftDeleteItem(id int) error {
	err := i.DB.Model(&entity.Item{}).Where("id = ?", id).Update("is_sale", 0).Error
	if err != nil {
		return err
	}
	err = i.DB.Where("id = ?", id).Delete(&entity.Item{}).Error
	return err
}

func (i *ItemRepository) RestoreDeletedItem(id int) error {
	err := i.DB.Unscoped().Model(&entity.Item{}).Where("id = ?", id).Update("is_sale", 1).Error
	if err != nil {
		return err
	}
	err = i.DB.Unscoped().Model(&entity.Item{}).Where("id = ?", id).Update("delete_at", gorm.DeletedAt{}).Error
	return err
}

func (i *ItemRepository) DeleteItem(id int) error {
	err := i.DB.Unscoped().Where("id = ?", id).Delete(&entity.Item{}).Error
	return err
}

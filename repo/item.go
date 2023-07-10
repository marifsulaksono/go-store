package repo

import (
	"gostore/config"
	"gostore/entity"
)

func GetAllItem() ([]entity.ItemResponse, error) {
	var items []entity.ItemResponse
	err := config.DB.Preload("Category").Find(&items).Error
	return items, err
}

func GetItem(id int64) (entity.ItemResponse, error) {
	var item entity.ItemResponse
	err := config.DB.Where("id = ?", id).Preload("Category", "id NOT IN (?)", "cancelled").First(&item).Error
	return item, err
}

func GetItembyStatus(s int) ([]entity.ItemResponse, error) {
	var item []entity.ItemResponse
	err := config.DB.Where("isSale = ?", s).Preload("Category", "id NOT IN (?)", "cancelled").Find(&item).Error
	return item, err
}
func InsertItem(item entity.Item) error {
	err := config.DB.Create(&item).Error
	return err
}

func UpdateItem(id int64, item entity.Item) (error, error) {
	var itemId entity.Item
	errId := config.DB.Where("id = ?", id).Preload("Category", "id NOT IN (?)", "cancelled").First(&itemId).Error
	err := config.DB.Model(&itemId).Updates(item).Error
	return errId, err
}

func DeleteItem(id int64) error {
	err := config.DB.Where("id = ?", id).Delete(&entity.Item{}).Error
	return err
}

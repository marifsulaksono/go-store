package service

import (
	"gostore/entity"
	"gostore/repo"
)

func GetAllItems() ([]entity.ItemResponse, error) {
	items, err := repo.GetAllItem()
	return items, err
}

func GetItembyId(id int64) (entity.ItemResponse, error) {
	item, err := repo.GetItem(id)
	return item, err
}

func GetItembyStatus(s int) ([]entity.ItemResponse, error) {
	items, err := repo.GetItembyStatus(s)
	return items, err
}

func InsertItem(item entity.Item) error {
	err := repo.InsertItem(item)
	return err
}

func UpdateItem(id int64, item entity.Item) (error, error) {
	errId, err := repo.UpdateItem(id, item)
	return errId, err
}

func DeleteItem(id int64) error {
	err := repo.DeleteItem(id)
	return err
}

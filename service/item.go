package service

import (
	"gostore/entity"
	"gostore/helper"
	"gostore/repo"
)

type ItemService struct {
	Repo repo.ItemRepository
}

func NewItemService(r repo.ItemRepository) *ItemService {
	return &ItemService{
		Repo: r,
	}
}

func (i *ItemService) GetAllItems() ([]entity.ItemResponse, error) {
	items, err := i.Repo.GetAllItems()
	return items, err
}

func (i *ItemService) GetItembyId(id int) (entity.ItemResponse, error) {
	item, err := i.Repo.GetItemById(id)
	return item, err
}

func (i *ItemService) GetItembyStatus(s int) ([]entity.ItemResponse, error) {
	items, err := i.Repo.GetItemByStatus(s)
	return items, err
}

func (i *ItemService) SearchItem(keyword, order, sortBy string, minPrice, maxPrice, limit, page int) ([]entity.ItemResponse, error) {
	if order != "DESC" {
		order = "ASC"
	}

	if sortBy != "name" && sortBy != "stock" && sortBy != "price" {
		sortBy = "name"
	}

	offset := (page - 1) * limit

	items, err := i.Repo.SearchItem(keyword, order, sortBy, minPrice, maxPrice, limit, offset)
	return items, err

	// DB.Where("name LIKE ? AND price BETWEEN ? AND ?", "%"+keyword+"%", minPrice, maxPrice).Order(sortBy +
	// " " + order).Limit(limit).Offset(page).Preload("Category").Find(&items)
	// DB.Where("name LIKE ? AND price BETWEEN ? AND ?", "%"+keyword+"%", minPrice, maxPrice).Order(sortBy + " " + order).Limit(limit).Offset(page).Preload("Category").Find(&items)
}

func (i *ItemService) InsertItem(item *entity.Item) error {
	err := i.Repo.InsertItem(item)
	return err
}

func (i *ItemService) UpdateItem(id int, item *entity.Item) error {
	_, err := i.Repo.GetItemById(id)
	if err != nil {
		return err
	}

	err = i.Repo.UpdateItem(id, item)
	return err
}

func (i *ItemService) ChangeStatusItem(id, s int) error {
	itemCheck, err := i.Repo.GetItemById(id)
	if err != nil {
		return err
	} else if itemCheck.IsSale == s {
		return helper.ErrChangeStatusItem
	}

	err = i.Repo.ChangeStatusItem(id, s)
	return err
}

func (i *ItemService) SoftDeleteItem(id int) error {
	_, err := i.Repo.GetItemById(id)
	if err != nil {
		return helper.ErrRecDeleted
	}

	err = i.Repo.SoftDeleteItem(id)
	return err
}

func (i *ItemService) RestoreDeletedItem(id int) error {
	itemCheck, err := i.Repo.GetItemById(id)
	if err == nil || itemCheck.IsSale == 1 {
		return helper.ErrRecRestored
	}

	err = i.Repo.RestoreDeletedItem(id)
	return err
}

func (i *ItemService) DeleteItem(id int) error {
	err := i.Repo.DeleteItem(id)
	return err
}

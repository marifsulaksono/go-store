package controller

import (
	"encoding/json"
	"fmt"
	"gostore/entity"
	"gostore/helper"
	"gostore/service"
	"net/http"
)

type ItemController struct {
	Service service.ItemService
}

func NewItemContoller(s service.ItemService) *ItemController {
	return &ItemController{
		Service: s,
	}
}

func (i *ItemController) GetItems(w http.ResponseWriter, r *http.Request) {
	items, err := i.Service.GetAllItems()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helper.ResponseWrite(w, items, "Success get all items")
}

func (i *ItemController) GetItembyId(w http.ResponseWriter, r *http.Request) {
	if id, s := helper.IdVarsMux(w, r); s {
		item, err := i.Service.GetItembyId(id)
		if err != nil {
			helper.RecordNotFound(w, err)
			return
		}

		message := fmt.Sprintf("Success get item %d", id)
		helper.ResponseWrite(w, item, message)
	}
}

func (i *ItemController) InsertItem(w http.ResponseWriter, r *http.Request) {
	var item entity.Item
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	item.IsSale = 1
	if err := i.Service.InsertItem(&item); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	message := fmt.Sprintf("Success create item %v", item.Name)
	helper.ResponseWrite(w, &item, message)
}

func (i *ItemController) UpdateItem(w http.ResponseWriter, r *http.Request) {
	if id, s := helper.IdVarsMux(w, r); s {
		var item entity.Item
		err := json.NewDecoder(r.Body).Decode(&item)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		err = i.Service.UpdateItem(id, &item)
		if err != nil {
			helper.RecordNotFound(w, err)
			return
		}

		message := fmt.Sprintf("Success update item %d", id)
		helper.ResponseWrite(w, item, message)
	}
}

func (i *ItemController) SalesItem(w http.ResponseWriter, r *http.Request) {
	result, err := i.Service.GetItembyStatus(1)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if result == nil || len(result) < 1 {
		w.Write([]byte("No item found!"))
		return
	}

	helper.ResponseWrite(w, result, "Success get all sales item")
}

func (i *ItemController) ChangeItemtoSale(w http.ResponseWriter, r *http.Request) {
	if id, s := helper.IdVarsMux(w, r); s {
		err := i.Service.ChangeStatusItem(id, 1)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		message := fmt.Sprintf("Item %d is sale now", id)
		helper.ResponseWrite(w, id, message)
	}
}

func (i *ItemController) SoldoutsItem(w http.ResponseWriter, r *http.Request) {
	result, err := i.Service.GetItembyStatus(0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if result == nil || len(result) < 1 {
		w.Write([]byte("No item found!"))
		return
	}

	helper.ResponseWrite(w, result, "Success get all soldouts item")
}

func (i *ItemController) ChangeItemtoSoldout(w http.ResponseWriter, r *http.Request) {
	if id, s := helper.IdVarsMux(w, r); s {
		err := i.Service.ChangeStatusItem(id, 0)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		message := fmt.Sprintf("Item %d is soldout now", id)
		helper.ResponseWrite(w, id, message)
	}
}

func (i *ItemController) SoftDeleteItem(w http.ResponseWriter, r *http.Request) {
	if id, s := helper.IdVarsMux(w, r); s {
		if err := i.Service.SoftDeleteItem(id); err != nil {
			helper.RecordNotFound(w, err)
			return
		}

		message := fmt.Sprintf("Success delete item %d", id)
		helper.ResponseWrite(w, id, message)
	}
}

func (i *ItemController) RestoreDeletedItem(w http.ResponseWriter, r *http.Request) {
	if id, s := helper.IdVarsMux(w, r); s {
		err := i.Service.RestoreDeletedItem(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		message := fmt.Sprintf("Success restore item %d", id)
		helper.ResponseWrite(w, id, message)
	}
}

func (i *ItemController) DeleteItem(w http.ResponseWriter, r *http.Request) {
	if id, s := helper.IdVarsMux(w, r); s {
		if err := i.Service.DeleteItem(id); err != nil {
			helper.RecordNotFound(w, err)
			return
		}

		message := fmt.Sprintf("Success delete permanently item %d", id)
		helper.ResponseWrite(w, id, message)
	}
}

package controller

import (
	"encoding/json"
	"fmt"
	"gostore/entity"
	"gostore/helper"
	"gostore/repo"
	"gostore/service"
	"net/http"
)

// ==================== Function Data Items ====================

func GetAllItems(w http.ResponseWriter, r *http.Request) {
	items, err := service.GetAllItems()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helper.ResponseWrite(w, items, "Success get all items")
}

func GetItembyId(w http.ResponseWriter, r *http.Request) {
	if id, s := helper.IdVarsMux(w, r); s {
		item, err := service.GetItembyId(id)
		if err != nil {
			helper.RecordNotFound(w, err)
			return
		}

		message := fmt.Sprintf("Success get item %d", id)
		helper.ResponseWrite(w, item, message)
	}
}

func InsertItem(w http.ResponseWriter, r *http.Request) {
	var item entity.Item
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if err := service.InsertItem(item); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	message := fmt.Sprintf("Success create item %v", item.Name)
	helper.ResponseWrite(w, &item, message)
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	if id, s := helper.IdVarsMux(w, r); s {
		var item entity.Item
		err := json.NewDecoder(r.Body).Decode(&item)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		if errId, err := repo.UpdateItem(id, item); errId != nil {
			helper.RecordNotFound(w, err)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			return
		}

		message := fmt.Sprintf("Success update item %d", id)
		helper.ResponseWrite(w, item, message)
	}
}

func SalesItem(w http.ResponseWriter, r *http.Request) {
	result, err := service.GetItembyStatus(1)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if result == nil || len(result) < 1 {
		w.Write([]byte("No item found!"))
		return
	}

	helper.ResponseWrite(w, result, "Success get all sale items")
}

func SoldoutsItem(w http.ResponseWriter, r *http.Request) {
	result, err := service.GetItembyStatus(0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if result == nil || len(result) < 1 {
		w.Write([]byte("No item found!"))
		return
	}

	helper.ResponseWrite(w, result, "Success get all soldout items")
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	if id, s := helper.IdVarsMux(w, r); s {
		if _, err := service.GetItembyId(id); err != nil {
			helper.RecordNotFound(w, err)
			return
		}

		if err := service.DeleteItem(id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		message := fmt.Sprintf("Success delete item %d", id)
		helper.ResponseWrite(w, id, message)
	}
}

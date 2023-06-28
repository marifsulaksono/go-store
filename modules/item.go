package modules

import (
	"encoding/json"
	"fmt"
	"gostore/config"
	"gostore/entity"
	"gostore/helper"
	"gostore/middleware"
	"net/http"
)

// ==================== Function Data Items ====================

func GetAllItems(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var result []entity.ItemResponse
		if err := config.DB.Preload("Category").Find(&result).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		helper.ResponseWrite(w, result, "Success get all items data!")
		return
	}
	http.Error(w, "Method isn't valid!", http.StatusBadRequest)
}

func GetItembyId(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		ctx := r.Context()
		userId := ctx.Value(middleware.GOSTORE_USERID)
		fmt.Println(userId)

		id, s := helper.IdVarsMux(w, r)
		if !s {
			return
		}

		var result entity.ItemResponse
		err := config.DB.Where("id = ?", id).Preload("Category", "id NOT IN (?)", "cancelled").First(&result).Error
		if err != nil {
			helper.RecordNotFound(w, err)
			return
		}

		helper.ResponseWrite(w, result, "Success get the item data!")
		return
	}
	http.Error(w, "Method isn't valid!", http.StatusBadRequest)
}

func InsertItem(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var item entity.Item
		err := json.NewDecoder(r.Body).Decode(&item)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		if err := config.DB.Create(&item).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		helper.ResponseWrite(w, &item, "Success insert item data!")
		return
	}
	http.Error(w, "Method isn't valid!", http.StatusBadRequest)
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		id, s := helper.IdVarsMux(w, r)
		if !s {
			return
		}

		var item entity.Item
		err := json.NewDecoder(r.Body).Decode(&item)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		itemId := entity.Item{}
		if err := config.DB.Where("id = ?", id).First(&itemId).Error; err != nil {
			helper.RecordNotFound(w, err)
			return
		}

		if err := config.DB.Model(&itemId).Updates(item).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		helper.ResponseWrite(w, itemId, "Success update item data!")
		return
	}
	http.Error(w, "Method isn't valid!", http.StatusBadRequest)
}

func SalesItem(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var result []entity.ItemResponse
		err := config.DB.Where("isSale = ?", 1).Preload("Category", "id NOT IN (?)", "cancelled").Find(&result).Error
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if result == nil || len(result) < 1 {
			w.Write([]byte("No item found!"))
			return
		}

		helper.ResponseWrite(w, result, "Success get all sale items data!")
		return
	}
	http.Error(w, "Method isn't valid!", http.StatusBadRequest)
}

func SaleItem(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		id, s := helper.IdVarsMux(w, r)
		if !s {
			return
		}

		itemId := entity.Item{}
		if err := config.DB.Where("id = ?", id).First(&itemId).Error; err != nil {
			helper.RecordNotFound(w, err)
			return
		}

		if err := config.DB.Model(&itemId).Update("isSale", 1).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		helper.ResponseWrite(w, itemId.Id, "Update success, item is sale now!")
		return
	}
	http.Error(w, "Method isn't valid!", http.StatusBadRequest)
}

func SoldoutsItem(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var result []entity.ItemResponse
		err := config.DB.Where("isSale = ?", 0).Preload("Category", "id NOT IN (?)", "cancelled").Find(&result).Error
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if result == nil || len(result) < 1 {
			w.Write([]byte("No item found!"))
			return
		}

		helper.ResponseWrite(w, result, "Success get all soldout items data!")
		return
	}
	http.Error(w, "Method isn't valid!", http.StatusBadRequest)
}

func SoldoutItem(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		id, s := helper.IdVarsMux(w, r)
		if !s {
			return
		}

		var itemId entity.Item
		if err := config.DB.Where("id = ?", id).First(&itemId).Error; err != nil {
			helper.RecordNotFound(w, err)
			return
		}

		if err := config.DB.Model(&itemId).Update("isSale", 0).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		helper.ResponseWrite(w, itemId.Id, "Update success, item is sold out!")
		return
	}
	http.Error(w, "Method isn't valid!", http.StatusBadRequest)
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		id, s := helper.IdVarsMux(w, r)
		if !s {
			return
		}

		var itemId entity.Item
		if err := config.DB.Where("id = ?", id).First(&itemId).Error; err != nil {
			helper.RecordNotFound(w, err)
			return
		}

		if err := config.DB.Where("id = ?", id).Delete(&entity.Item{}).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		helper.ResponseWrite(w, nil, "Success delete item data!")
		return
	}
	http.Error(w, "Method isn't valid!", http.StatusBadRequest)
}

func CategoryItem(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		id, s := helper.IdVarsMux(w, r)
		if !s {
			return
		}

		var result []entity.Item
		if err := config.DB.Where("category_id = ?", id).Preload("Category", "id NOT IN (?)", "cancelled").Find(&result).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if result == nil || len(result) < 1 {
			w.Write([]byte("No item found!"))
			return
		}

		helper.ResponseWrite(w, result, "Success get all categories!")
		return
	}
	http.Error(w, "Method isn't valid!", http.StatusBadRequest)
}

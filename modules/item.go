package modules

import (
	"encoding/json"
	"errors"
	"gostore/config"
	"gostore/entity"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// ==================== Function Data Items ====================

func GetAllItems(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var result []entity.ItemResponse
		if err := config.DB.Preload("Category").Find(&result).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		ResponseWrite(w, result, "Success get all items data!")
		return
	}
	http.Error(w, "Method isn't valid!", http.StatusBadRequest)
}

func GetItembyId(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		vars := mux.Vars(r)
		id, err := strconv.ParseInt(vars["id"], 10, 0)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var result entity.ItemResponse
		if err := config.DB.Where("id = ?", id).Preload("Category", "id NOT IN (?)", "cancelled").First(&result).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				w.Write([]byte("id not found!"))
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		ResponseWrite(w, result, "Success get the item data!")
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

		ResponseWrite(w, &item, "Success insert item data!")
		return
	}
	http.Error(w, "Method isn't valid!", http.StatusBadRequest)
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		vars := mux.Vars(r)
		id, errV := strconv.ParseInt(vars["id"], 10, 0)
		if errV != nil {
			http.Error(w, errV.Error(), http.StatusBadRequest)
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
			if errors.Is(err, gorm.ErrRecordNotFound) {
				w.Write([]byte("Id not found"))
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := config.DB.Model(&itemId).Updates(item).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		ResponseWrite(w, itemId, "Success update item data!")
		return
	}
	http.Error(w, "Method isn't valid!", http.StatusBadRequest)
}

func SalesItem(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var result []entity.ItemResponse
		if err := config.DB.Where("isSale = ?", 1).Preload("Category", "id NOT IN (?)", "cancelled").Find(&result).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if result == nil || len(result) < 1 {
			w.Write([]byte("No item found!"))
			return
		}

		ResponseWrite(w, result, "Success get all sale items data!")
		return
	}
	http.Error(w, "Method isn't valid!", http.StatusBadRequest)
}

func SaleItem(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		vars := mux.Vars(r)
		id, errV := strconv.ParseInt(vars["id"], 10, 0)
		if errV != nil {
			http.Error(w, errV.Error(), http.StatusBadRequest)
			return
		}

		itemId := entity.Item{}
		if err := config.DB.Where("id = ?", id).First(&itemId).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				http.Error(w, "id not found!", http.StatusBadRequest)
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := config.DB.Model(&itemId).Update("isSale", 1).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		ResponseWrite(w, itemId.Id, "Update success, item is sale now!")
		return
	}
	http.Error(w, "Method isn't valid!", http.StatusBadRequest)
}

func SoldoutsItem(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var result []entity.ItemResponse
		if err := config.DB.Where("isSale = ?", 2).Preload("Category", "id NOT IN (?)", "cancelled").Find(&result).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if result == nil || len(result) < 1 {
			w.Write([]byte("No item found!"))
			return
		}

		ResponseWrite(w, result, "Success get all soldout items data!")
		return
	}
	http.Error(w, "Method isn't valid!", http.StatusBadRequest)
}

func SoldoutItem(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		vars := mux.Vars(r)
		id, errV := strconv.ParseInt(vars["id"], 10, 0)
		if errV != nil {
			http.Error(w, errV.Error(), http.StatusNotFound)
			return
		}

		var itemId entity.Item
		if err := config.DB.Where("id = ?", id).First(&itemId).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				http.Error(w, "id not found!", http.StatusBadRequest)
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := config.DB.Model(&itemId).Update("isSale", 2).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		ResponseWrite(w, itemId.Id, "Update success, item is sold out!")
		return
	}
	http.Error(w, "Method isn't valid!", http.StatusBadRequest)
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		vars := mux.Vars(r)
		id, errV := strconv.ParseInt(vars["id"], 10, 0)
		if errV != nil {
			http.Error(w, errV.Error(), http.StatusBadRequest)
			return
		}

		var itemId entity.Item
		if err := config.DB.Where("id = ?", id).First(&itemId).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				http.Error(w, "id not found!", http.StatusBadRequest)
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := config.DB.Where("id = ?", id).Delete(&entity.Item{}).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		ResponseWrite(w, nil, "Success delete item data!")
		return
	}
	http.Error(w, "Method isn't valid!", http.StatusBadRequest)
}

func CategoryItem(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		vars := mux.Vars(r)
		id, errV := strconv.ParseInt(vars["id"], 10, 0)
		if errV != nil {
			http.Error(w, errV.Error(), http.StatusNotFound)
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

		ResponseWrite(w, result, "Success get all categories!")
		return
	}
	http.Error(w, "Method isn't valid!", http.StatusBadRequest)
}

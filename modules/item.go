package modules

import (
	"encoding/json"
	"fmt"
	"gostore/config"
	"gostore/entity"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// ==================== Function Data Items ====================

func GetAllItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		var result []entity.ItemResponse
		if result := config.DB.Preload("Category").Find(&result); result.Error != nil {
			http.Error(w, result.Error.Error(), http.StatusInternalServerError)
			return
		}

		var jsonData, errJ = json.Marshal(result)
		if errJ != nil {
			http.Error(w, errJ.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(jsonData)
		return
	}
	http.Error(w, "Data not found", http.StatusNotFound)
}

func GetItembyId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		vars := mux.Vars(r)
		id, err := strconv.ParseInt(vars["id"], 10, 0)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		var result entity.ItemResponse
		if result := config.DB.Where("id = ?", id).Preload("Category", "id NOT IN (?)", "cancelled").Find(&result); result.Error != nil {
			http.Error(w, result.Error.Error(), http.StatusInternalServerError)
			return
		}

		var jsonData, errJ = json.Marshal(result)
		if errJ != nil {
			fmt.Println()
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(jsonData)
		return
	}
	http.Error(w, "Data not found", http.StatusNotFound)
}

func InsertItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		var item entity.Item
		err := json.NewDecoder(r.Body).Decode(&item)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		if errC := config.DB.Create(&item).Error; errC != nil {
			fmt.Println(errC.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Add item success!"))
		return
	}
	http.Error(w, "", http.StatusNotFound)
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "PUT" {
		vars := mux.Vars(r)
		id, errV := strconv.ParseInt(vars["id"], 10, 0)
		if errV != nil {
			http.Error(w, errV.Error(), http.StatusInternalServerError)
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
		config.DB.Where("id = ?", id).First(&itemId)
		if errU := config.DB.Model(&itemId).Updates(item).Error; errU != nil {
			w.Write([]byte("Id not found"))
			fmt.Println(errU.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Update item success!"))
		return
	}
	http.Error(w, "", http.StatusNotFound)
}

func SalesItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		var result []entity.ItemResponse
		if result := config.DB.Where("isSale = ?", 1).Preload("Category", "id NOT IN (?)", "cancelled").Find(&result); result.Error != nil {
			http.Error(w, result.Error.Error(), http.StatusInternalServerError)
			return
		}

		var jsonData, errJ = json.Marshal(result)
		if errJ != nil {
			fmt.Println(errJ.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(jsonData)
		return
	}
	http.Error(w, "Data not found", http.StatusNotFound)
}

func SaleItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "PUT" {
		vars := mux.Vars(r)
		id, errV := strconv.ParseInt(vars["id"], 10, 0)
		if errV != nil {
			fmt.Println(errV.Error())
			return
		}

		itemId := entity.Item{}
		config.DB.Where("id = ?", id).First(&itemId)
		config.DB.Model(&itemId).Update("isSale", 1)

		w.WriteHeader(http.StatusOK)
		var message = fmt.Sprintf("Item %d is sale now!", id)
		w.Write([]byte(message))
		return
	}
	http.Error(w, "Data not found", http.StatusNotFound)
}

func SoldoutsItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		var result []entity.ItemResponse
		if result := config.DB.Where("isSale = ?", 2).Preload("Category", "id NOT IN (?)", "cancelled").Find(&result); result.Error != nil {
			http.Error(w, result.Error.Error(), http.StatusInternalServerError)
			return
		}

		var jsonData, errJ = json.Marshal(result)
		if errJ != nil {
			fmt.Println(errJ.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(jsonData)
		return
	}
	http.Error(w, "Data not found", http.StatusNotFound)
}

func SoldoutItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "PUT" {
		vars := mux.Vars(r)
		id, errV := strconv.ParseInt(vars["id"], 10, 0)
		if errV != nil {
			fmt.Println(errV.Error())
			return
		}

		itemId := entity.Item{}
		config.DB.Where("id = ?", id).First(&itemId)
		config.DB.Model(&itemId).Update("isSale", 2)

		w.WriteHeader(http.StatusOK)
		var message = fmt.Sprintf("Item %d is sold out now!", id)
		w.Write([]byte(message))
		return
	}
	http.Error(w, "", http.StatusNotFound)
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "DELETE" {
		vars := mux.Vars(r)
		id, errV := strconv.ParseInt(vars["id"], 10, 0)
		if errV != nil {
			fmt.Println(errV.Error())
			return
		}

		if errD := config.DB.Where("id = ?", id).Delete(&entity.Item{}).Error; errD != nil {
			w.Write([]byte("Id not found"))
			fmt.Println(errD.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Delete success!"))
		return
	}
	http.Error(w, "", http.StatusNotFound)
}

func CategoryItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		vars := mux.Vars(r)
		id, errV := strconv.ParseInt(vars["id"], 10, 0)
		if errV != nil {
			http.Error(w, errV.Error(), http.StatusInternalServerError)
			return
		}

		var result []entity.Item
		if result := config.DB.Where("category_id = ?", id).Preload("Category", "id NOT IN (?)", "cancelled").Find(&result); result.Error != nil {
			http.Error(w, result.Error.Error(), http.StatusInternalServerError)
			return
		}

		var jsonData, errJ = json.Marshal(result)
		if errJ != nil {
			fmt.Println(errJ.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(jsonData)
		return
	}
	http.Error(w, "Data not found", http.StatusNotFound)
}

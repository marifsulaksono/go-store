package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Item struct {
	Id         int      `json:"id"`
	Name       string   `json:"name"`
	Stock      int      `json:"stock"`
	Price      int      `json:"price"`
	Sale       int      `gorm:"column:isSale" json:"isSale"`
	CategoryId int      `json:"categoryId"`
	Category   Category `json:"category"`
	// Category   Category `gorm:"foreignKey:IdCategory" json:"category"` // inisialisasi foreignkey pada gorm
}

type Category struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func connect() (*gorm.DB, error) {
	newDB := "root:@tcp(127.0.0.1:3306)/db_store"
	db, err := gorm.Open(mysql.Open(newDB), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

// ==================== Fungsi jSON Data Items ====================

func getAllItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		db, err := connect()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		var result []Item
		if result := db.Preload("Category").Find(&result); result.Error != nil {
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

func getItembyId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		vars := mux.Vars(r)
		id, err := strconv.ParseInt(vars["id"], 10, 0)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		db, err := connect()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		var result Item
		if result := db.Where("id = ?", id).Preload("Category", "id NOT IN (?)", "cancelled").Find(&result); result.Error != nil {
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

func insertItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		var item Item
		err := json.NewDecoder(r.Body).Decode(&item)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		db, errC := connect()
		if errC != nil {
			fmt.Println(errC.Error())
			return
		}

		if errC := db.Create(&item).Error; errC != nil {
			fmt.Println(errC.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Add item success!"))
		return
	}
	http.Error(w, "", http.StatusNotFound)
}

func updateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "PUT" {
		vars := mux.Vars(r)
		id, errV := strconv.ParseInt(vars["id"], 10, 0)
		if errV != nil {
			http.Error(w, errV.Error(), http.StatusInternalServerError)
			return
		}

		var item Item
		err := json.NewDecoder(r.Body).Decode(&item)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		db, err := connect()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		itemId := Item{}
		db.Where("id = ?", id).First(&itemId)
		if errU := db.Model(&itemId).Updates(item).Error; errU != nil {
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

func salesItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		db, err := connect()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		var result []Item
		if result := db.Where("isSale = ?", 1).Preload("Category", "id NOT IN (?)", "cancelled").Find(&result); result.Error != nil {
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

func saleItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "PUT" {
		vars := mux.Vars(r)
		id, errV := strconv.ParseInt(vars["id"], 10, 0)
		if errV != nil {
			fmt.Println(errV.Error())
			return
		}

		db, err := connect()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		itemId := Item{}
		db.Where("id = ?", id).First(&itemId)
		db.Model(&itemId).Update("isSale", 1)

		w.WriteHeader(http.StatusOK)
		var message = fmt.Sprintf("Item %d is sale now!", id)
		w.Write([]byte(message))
		return
	}
	http.Error(w, "Data not found", http.StatusNotFound)
}

func soldoutsItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		db, err := connect()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		var result []Item
		if result := db.Where("isSale = ?", 2).Preload("Category", "id NOT IN (?)", "cancelled").Find(&result); result.Error != nil {
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

func soldoutItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "PUT" {
		vars := mux.Vars(r)
		id, errV := strconv.ParseInt(vars["id"], 10, 0)
		if errV != nil {
			fmt.Println(errV.Error())
			return
		}

		db, err := connect()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		itemId := Item{}
		db.Where("id = ?", id).First(&itemId)
		db.Model(&itemId).Update("isSale", 2)

		w.WriteHeader(http.StatusOK)
		var message = fmt.Sprintf("Item %d is sold out now!", id)
		w.Write([]byte(message))
		return
	}
	http.Error(w, "", http.StatusNotFound)
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "DELETE" {
		vars := mux.Vars(r)
		id, errV := strconv.ParseInt(vars["id"], 10, 0)
		if errV != nil {
			fmt.Println(errV.Error())
			return
		}

		db, err := connect()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		if errD := db.Where("id = ?", id).Delete(&Item{}).Error; errD != nil {
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

func categoryItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		vars := mux.Vars(r)
		id, errV := strconv.ParseInt(vars["id"], 10, 0)
		if errV != nil {
			http.Error(w, errV.Error(), http.StatusInternalServerError)
			return
		}

		db, err := connect()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		var result []Item
		if result := db.Where("category_id = ?", id).Preload("Category", "id NOT IN (?)", "cancelled").Find(&result); result.Error != nil {
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

// ==================== Fungsi main ====================

func main() {
	r := mux.NewRouter()
	// ===================== Handler Item ======================
	r.HandleFunc("/items", getAllItems).Methods("GET")                // get all sales & soldout item data
	r.HandleFunc("/item/{id}", getItembyId).Methods("GET")            // get item data by id
	r.HandleFunc("/items", insertItem).Methods("POST")                // create new item data
	r.HandleFunc("/items/{id}", updateItem).Methods("PUT")            // update item data by id
	r.HandleFunc("/items/sale", salesItem).Methods("GET")             // get all sale item data
	r.HandleFunc("/items/sale/{id}", saleItem).Methods("PUT")         // update item to sale
	r.HandleFunc("/items/soldout", soldoutsItem).Methods("GET")       // get all soldout item data
	r.HandleFunc("/items/soldout/{id}", soldoutItem).Methods("PUT")   // update item to soldout
	r.HandleFunc("/items/{id}", deleteItem).Methods("DELETE")         // delete item data by id
	r.HandleFunc("/items/category/{id}", categoryItem).Methods("GET") // filter items by category

	// ==================== Start Server ====================
	fmt.Println("Server started at localhost:8080")
	http.ListenAndServe(":8080", r)
}

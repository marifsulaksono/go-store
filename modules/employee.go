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

// ==================== Function Data Employee ====================

func GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		result := []entity.Employee{}
		config.DB.Find(&result)

		var jsonData, err = json.Marshal(result)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(jsonData)
		return
	}
	http.Error(w, "Data not found", http.StatusNotFound)
}

func GetEmployeebyId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		vars := mux.Vars(r)
		id, err := strconv.ParseInt(vars["id"], 10, 0)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		result := entity.Employee{}
		if errG := config.DB.Where("id = ?", id).First(&result).Error; errG != nil {
			http.Error(w, errG.Error(), http.StatusInternalServerError)
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

func InsertEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		var employee entity.Employee
		err := json.NewDecoder(r.Body).Decode(&employee)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		defer r.Body.Close()

		if errC := config.DB.Create(&employee).Error; errC != nil {
			fmt.Println(errC.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Add success!"))
		return
	}
	http.Error(w, "", http.StatusNotFound)
}

func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "PUT" {
		vars := mux.Vars(r)
		id, errV := strconv.ParseInt(vars["id"], 10, 0)
		if errV != nil {
			http.Error(w, errV.Error(), http.StatusInternalServerError)
			return
		}

		var employee entity.Employee
		err := json.NewDecoder(r.Body).Decode(&employee)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		emply := entity.Employee{}
		config.DB.Where("id = ?", id).First(&emply)
		if errU := config.DB.Model(&emply).Updates(employee).Error; errU != nil {
			fmt.Println(errU.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Update success!"))
		return
	}
	http.Error(w, "", http.StatusNotFound)
}

func ActivedEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		result := []entity.Employee{}
		config.DB.Where("isActive = ?", "Aktif").Find(&result)

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

func ActiveEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "PUT" {
		vars := mux.Vars(r)
		id, err := strconv.ParseInt(vars["id"], 10, 0)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		employee := entity.Employee{}
		config.DB.Where("id = ?", id).First(&employee)
		config.DB.Model(&employee).Update("isActive", "Aktif")

		w.WriteHeader(http.StatusOK)
		var message = fmt.Sprintf("Employee %d active now!", id)
		w.Write([]byte(message))
		return
	}
	http.Error(w, "", http.StatusNotFound)
}

func InactivedEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		result := []entity.Employee{}
		config.DB.Where("isActive = ?", "Nonaktif").Find(&result)

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

func InactiveEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "PUT" {
		vars := mux.Vars(r)
		id, err := strconv.ParseInt(vars["id"], 10, 0)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		employee := entity.Employee{}
		config.DB.Where("id = ?", id).First(&employee)
		config.DB.Model(&employee).Update("isActive", "Nonaktif")

		w.WriteHeader(http.StatusOK)
		var message = fmt.Sprintf("Employee %d non-active now!", id)
		w.Write([]byte(message))
		return
	}
	http.Error(w, "", http.StatusNotFound)
}

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "DELETE" {
		vars := mux.Vars(r)
		id, err := strconv.ParseInt(vars["id"], 10, 0)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		if errD := config.DB.Where("id = ?", id).Delete(&entity.Employee{}).Error; errD != nil {
			w.Write([]byte("Id not found"))
			fmt.Println(errD.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Delete success!"))
		return
	}
	http.Error(w, "Data not found", http.StatusNotFound)
}

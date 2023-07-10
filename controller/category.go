package controller

import (
	"gostore/config"
	"gostore/entity"
	"gostore/helper"

	"net/http"
)

func GetAllCategory(w http.ResponseWriter, r *http.Request) {
	if id, s := helper.IdVarsMux(w, r); s {
		var result []entity.Item
		if err := config.DB.Where("category_id = ?", id).Preload("Category", "id NOT IN (?)", "cancelled").Find(&result).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if result == nil || len(result) < 1 {
			w.Write([]byte("No category found!"))
			return
		}

		helper.ResponseWrite(w, result, "Success get all categories!")
	}
}

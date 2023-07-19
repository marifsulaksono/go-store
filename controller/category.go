package controller

import (
	"gostore/helper"
	"gostore/service"

	"net/http"
)

type CategoryController struct {
	Service service.CategoryService
}

func NewCategoryController(s service.CategoryService) *CategoryController {
	return &CategoryController{
		Service: s,
	}
}

func (c *CategoryController) GetAllCategories(w http.ResponseWriter, r *http.Request) {
	result, err := c.Service.GetAllCategories()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if result == nil || len(result) < 1 {
		w.Write([]byte("No category found!"))
		return
	}

	helper.ResponseWrite(w, result, "Success get all categories!")
}

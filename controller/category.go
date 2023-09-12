package controller

import (
	"encoding/json"
	"gostore/entity"
	"gostore/helper"
	"gostore/helper/response"
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
	var (
		ctx = r.Context()
	)

	result, err := c.Service.GetAllCategories(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if result == nil || len(result) < 1 {
		w.Write([]byte("No category found!"))
		return
	}

	response.BuildSuccesResponse(w, result, nil, "Success get all categories!")
}

func (c *CategoryController) GetCategoryById(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = r.Context()
	)

	id, err := helper.ParamIdChecker(w, r)
	if err != nil {
		response.BuildErorResponse(w, err)
		return
	}

	result, err := c.Service.GetCategoryById(ctx, id)
	if err != nil {
		response.BuildErorResponse(w, err)
		return
	}

	response.BuildSuccesResponse(w, result, nil, "")
}

func (c *CategoryController) InsertCategory(w http.ResponseWriter, r *http.Request) {
	var (
		ctx      = r.Context()
		category entity.Category
	)

	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		response.BuildErorResponse(w, err)
		return
	}
	defer r.Body.Close()

	if err := c.Service.InsertCategory(ctx, &category); err != nil {
		response.BuildErorResponse(w, err)
		return
	}

	response.BuildSuccesResponse(w, nil, nil, "Success create new category")
}

func (c *CategoryController) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	var (
		ctx      = r.Context()
		category entity.Category
	)

	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		response.BuildErorResponse(w, err)
		return
	}
	defer r.Body.Close()

	id, err := helper.ParamIdChecker(w, r)
	if err != nil {
		response.BuildErorResponse(w, err)
		return
	}

	if err := c.Service.UpdateCategory(ctx, id, &category); err != nil {
		response.BuildErorResponse(w, err)
		return
	}

	response.BuildSuccesResponse(w, nil, nil, "Success update category")
}

func (c *CategoryController) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = r.Context()
	)

	id, err := helper.ParamIdChecker(w, r)
	if err != nil {
		response.BuildErorResponse(w, err)
		return
	}

	if err := c.Service.DeleteCategory(ctx, id); err != nil {
		response.BuildErorResponse(w, err)
		return
	}

	response.BuildSuccesResponse(w, nil, nil, "Success delete category")
}

package controller

import (
	"encoding/json"
	"fmt"
	"gostore/entity"
	"gostore/helper"
	"gostore/helper/response"
	"gostore/service"
	"net/http"
	"strconv"
	"strings"
)

type ProductController struct {
	Service service.ProductService
}

func NewProductController(s service.ProductService) *ProductController {
	return &ProductController{
		Service: s,
	}
}

func (p *ProductController) GetProducts(w http.ResponseWriter, r *http.Request) {
	var (
		ctx     = r.Context()
		message string
	)

	keyword := r.URL.Query().Get("keyword")
	status := r.URL.Query().Get("status")
	sortBy := r.URL.Query().Get("sortBy")
	order := strings.ToUpper(r.URL.Query().Get("order"))
	minPrice, _ := strconv.Atoi(r.URL.Query().Get("minPrice"))
	maxPrice, _ := strconv.Atoi(r.URL.Query().Get("maxPrice"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	categoryId, _ := strconv.Atoi(r.URL.Query().Get("categoryId"))
	storeId, _ := strconv.Atoi(r.URL.Query().Get("storeId"))

	products, pagination, err := p.Service.GetAllProducts(ctx, keyword, status, order, sortBy, minPrice, maxPrice, categoryId, storeId, limit, page)
	if err != nil {
		response.BuildErorResponse(w, err)
		return
	}

	if products == nil || len(products) < 1 {
		products = nil
		pagination = helper.Page{}
		message = "No results found"
	}

	response.BuildSuccesResponse(w, products, pagination, message)
	// localhost:49999/product/search?keyword=tas&sortBy=name&order=desc&minPrice=0&maxPrice=1000000&limit=5&page=0
}

func (p *ProductController) GetProductbyId(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = r.Context()
	)

	id, err := helper.ParamIdChecker(w, r)
	if err != nil {
		response.BuildErorResponse(w, err)
		return
	}

	result, err := p.Service.GetProductbyId(ctx, id)
	if err != nil {
		response.BuildErorResponse(w, err)
		return
	}

	response.BuildSuccesResponse(w, result, nil, "")
}

func (p *ProductController) InsertProduct(w http.ResponseWriter, r *http.Request) {
	var (
		ctx     = r.Context()
		product entity.Product
	)

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		response.BuildErorResponse(w, err)
		return
	}
	defer r.Body.Close()

	if err := p.Service.InsertProduct(ctx, &product); err != nil {
		response.BuildErorResponse(w, err)
		return
	}

	message := fmt.Sprintf("Success create product %v", product.Name)
	response.BuildSuccesResponse(w, nil, nil, message)
}

func (p *ProductController) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = r.Context()
	)

	id, err := helper.ParamIdChecker(w, r)
	if err != nil {
		response.BuildErorResponse(w, err)
		return
	}
	var product entity.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		response.BuildErorResponse(w, err)
		return
	}
	defer r.Body.Close()

	if err := p.Service.UpdateProduct(ctx, id, &product); err != nil {
		response.BuildErorResponse(w, err)
		return
	}

	response.BuildSuccesResponse(w, nil, nil, "Success update product")
}

func (p *ProductController) SoftDeleteProduct(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = r.Context()
	)

	id, err := helper.ParamIdChecker(w, r)
	if err != nil {
		response.BuildErorResponse(w, err)
		return
	}

	if err := p.Service.SoftDeleteProduct(ctx, id); err != nil {
		response.BuildErorResponse(w, err)
		return
	}

	response.BuildSuccesResponse(w, nil, nil, "Success delete product")
}

func (p *ProductController) RestoreDeletedProduct(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = r.Context()
	)

	id, err := helper.ParamIdChecker(w, r)
	if err != nil {
		response.BuildErorResponse(w, err)
		return
	}

	if err := p.Service.RestoreDeletedProduct(ctx, id); err != nil {
		response.BuildErorResponse(w, err)
		return
	}

	response.BuildSuccesResponse(w, nil, nil, "Success restore product")
}

func (p *ProductController) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = r.Context()
	)

	id, err := helper.ParamIdChecker(w, r)
	if err != nil {
		response.BuildErorResponse(w, err)
		return
	}

	if err := p.Service.DeleteProduct(ctx, id); err != nil {
		response.BuildErorResponse(w, err)
		return
	}

	message := fmt.Sprintf("Success delete permanently product %d", id)
	response.BuildSuccesResponse(w, nil, nil, message)
}

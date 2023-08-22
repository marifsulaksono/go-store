package controller

import (
	"encoding/json"
	"fmt"
	"gostore/entity"
	"gostore/helper"
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
	ctx := r.Context()
	// keyword := ""
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

	products, err := p.Service.GetAllProducts(ctx, keyword, status, order, sortBy, minPrice, maxPrice, categoryId, storeId, limit, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if products == nil || len(products) < 1 {
		w.Write([]byte("No product found!"))
		return
	}

	helper.ResponseWrite(w, products, "Success search products")
	// localhost:49999/product/search?keyword=tas&sortBy=name&order=desc&minPrice=0&maxPrice=1000000&limit=5&page=0
}

func (p *ProductController) GetProductbyId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	if id, s := helper.IdVarsMux(w, r); s {
		result, err := p.Service.GetProductbyId(ctx, id)
		if err != nil {
			helper.BuildError(w, err)
			return
		}

		helper.BuildResponseSuccess(w, result, "")
	}
}

func (p *ProductController) InsertProduct(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var product entity.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if err := p.Service.InsertProduct(ctx, &product); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	message := fmt.Sprintf("Success create product %v", product.Name)
	helper.ResponseWrite(w, product, message)
}

func (p *ProductController) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	if id, s := helper.IdVarsMux(w, r); s {
		var product entity.Product
		err := json.NewDecoder(r.Body).Decode(&product)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		err = p.Service.UpdateProduct(ctx, id, &product)
		if err != nil {
			helper.RecordNotFound(w, err)
			return
		}

		message := fmt.Sprintf("Success update product %d", id)
		helper.ResponseWrite(w, product, message)
	}
}

func (p *ProductController) SoftDeleteProduct(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	if id, s := helper.IdVarsMux(w, r); s {
		if err := p.Service.SoftDeleteProduct(ctx, id); err != nil {
			helper.RecordNotFound(w, err)
			return
		}

		message := fmt.Sprintf("Success delete product %d", id)
		helper.ResponseWrite(w, id, message)
	}
}

func (p *ProductController) RestoreDeletedProduct(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	if id, s := helper.IdVarsMux(w, r); s {
		err := p.Service.RestoreDeletedProduct(ctx, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		message := fmt.Sprintf("Success restore product %d", id)
		helper.ResponseWrite(w, id, message)
	}
}

func (p *ProductController) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	if id, s := helper.IdVarsMux(w, r); s {
		if err := p.Service.DeleteProduct(ctx, id); err != nil {
			helper.RecordNotFound(w, err)
			return
		}

		// message := fmt.Sprintf("Success delete permanently product %d", id)
		helper.ResponseWrite(w, id, fmt.Sprintf("Success delete permanently product %d", id))
	}
}

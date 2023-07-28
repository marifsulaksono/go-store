package controller

import (
	"encoding/json"
	"fmt"
	"gostore/entity"
	"gostore/helper"
	"gostore/service"
	"math"
	"net/http"
	"strconv"
	"strings"
)

type ProductController struct {
	Service service.ProductService
}

func NewProductContoller(s service.ProductService) *ProductController {
	return &ProductController{
		Service: s,
	}
}

func (p *ProductController) GetProducts(w http.ResponseWriter, r *http.Request) {
	result, err := p.Service.GetAllProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helper.ResponseWrite(w, result, "Success get all products")
}

func (p *ProductController) GetProductbyId(w http.ResponseWriter, r *http.Request) {
	if id, s := helper.IdVarsMux(w, r); s {
		result, err := p.Service.GetProductbyId(id)
		if err != nil {
			helper.RecordNotFound(w, err)
			return
		}

		message := fmt.Sprintf("Success get product %d", id)
		helper.ResponseWrite(w, result, message)
	}
}

func (p *ProductController) SearchProduct(w http.ResponseWriter, r *http.Request) {
	keyword := ""
	keyword = r.URL.Query().Get("keyword")
	sortBy := r.URL.Query().Get("sortBy")
	order := strings.ToUpper(r.URL.Query().Get("order"))
	minPrice, err := strconv.ParseFloat(r.URL.Query().Get("minPrice"), 64)
	if err != nil || minPrice < 0 {
		minPrice = 0
		fmt.Println(err)
	}
	maxPrice, err := strconv.ParseFloat(r.URL.Query().Get("maxPrice"), 64)
	if err != nil {
		maxPrice = math.MaxFloat64
	} else if maxPrice < minPrice {
		http.Error(w, "max price must be higher than min price", http.StatusBadRequest)
		return
	}
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil || limit <= 0 {
		limit = 10
	}
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page <= 0 {
		page = 1
	}
	fmt.Printf("%v, %v, %v, %v, %v, %v, %v | ", keyword, sortBy, order, minPrice, maxPrice, limit, page)

	products, err := p.Service.SearchProduct(keyword, order, sortBy, minPrice, maxPrice, limit, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	message := fmt.Sprintf("Success search product with keyword %v", keyword)
	helper.ResponseWrite(w, products, message)
	// localhost:49999/product/search?keyword=tas&sortBy=name&order=desc&minPrice=0&maxPrice=1000000&limit=5&page=0
}

func (p *ProductController) InsertProduct(w http.ResponseWriter, r *http.Request) {
	var product entity.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if err := p.Service.InsertProduct(&product); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	message := fmt.Sprintf("Success create product %v", product.Name)
	helper.ResponseWrite(w, product, message)
}

func (p *ProductController) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	if id, s := helper.IdVarsMux(w, r); s {
		var product entity.Product
		err := json.NewDecoder(r.Body).Decode(&product)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		result, err := p.Service.UpdateProduct(id, &product)
		if err != nil {
			helper.RecordNotFound(w, err)
			return
		}

		message := fmt.Sprintf("Success update product %d", id)
		helper.ResponseWrite(w, result, message)
	}
}

func (p *ProductController) SoldoutsProduct(w http.ResponseWriter, r *http.Request) {
	result, err := p.Service.GetProductbyStatus("soldout")
	if err != nil {
		helper.RecordNotFound(w, err)
		return
	}

	if result == nil || len(result) < 1 {
		w.Write([]byte("No product found!"))
		return
	}

	helper.ResponseWrite(w, result, "Success get all soldouts product")
}

func (p *ProductController) ChangeProducttoSale(w http.ResponseWriter, r *http.Request) {
	if id, s := helper.IdVarsMux(w, r); s {
		err := p.Service.ChangeStatusProduct(id, "sale")
		if err != nil {
			helper.RecordNotFound(w, err)
			return
		}

		message := fmt.Sprintf("Product %d is sale now", id)
		helper.ResponseWrite(w, id, message)
	}
}

func (p *ProductController) ChangeProducttoSoldout(w http.ResponseWriter, r *http.Request) {
	if id, s := helper.IdVarsMux(w, r); s {
		err := p.Service.ChangeStatusProduct(id, "soldout")
		if err != nil {
			helper.RecordNotFound(w, err)
			return
		}

		message := fmt.Sprintf("Product %d is soldout now", id)
		helper.ResponseWrite(w, id, message)
	}
}

func (p *ProductController) SoftDeleteProduct(w http.ResponseWriter, r *http.Request) {
	if id, s := helper.IdVarsMux(w, r); s {
		if err := p.Service.SoftDeleteProduct(id); err != nil {
			helper.RecordNotFound(w, err)
			return
		}

		message := fmt.Sprintf("Success delete product %d", id)
		helper.ResponseWrite(w, id, message)
	}
}

func (p *ProductController) RestoreDeletedProduct(w http.ResponseWriter, r *http.Request) {
	if id, s := helper.IdVarsMux(w, r); s {
		err := p.Service.RestoreDeletedProduct(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		message := fmt.Sprintf("Success restore product %d", id)
		helper.ResponseWrite(w, id, message)
	}
}

func (p *ProductController) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	if id, s := helper.IdVarsMux(w, r); s {
		if err := p.Service.DeleteProduct(id); err != nil {
			helper.RecordNotFound(w, err)
			return
		}

		message := fmt.Sprintf("Success delete permanently product %d", id)
		helper.ResponseWrite(w, id, message)
	}
}

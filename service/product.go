package service

import (
	"context"
	"fmt"
	"gostore/entity"
	"gostore/helper"
	productError "gostore/helper/domain/errorModel"
	"gostore/repo"
)

type productService struct {
	Repo repo.ProductRepository
}

type ProductService interface {
	GetAllProducts(ctx context.Context, keyword, status, order, sortBy string, minPrice, maxPrice, categoryId, storeId, limit, page int) ([]entity.Product, helper.Page, error)
	GetProductbyId(ctx context.Context, id int) (entity.Product, error)
	InsertProduct(ctx context.Context, product *entity.Product) error
	UpdateProduct(ctx context.Context, id int, product *entity.Product) error
	SoftDeleteProduct(ctx context.Context, id int) error
	RestoreDeletedProduct(ctx context.Context, id int) error
	DeleteProduct(ctx context.Context, id int) error
}

func NewProductService(r repo.ProductRepository) ProductService {
	return &productService{
		Repo: r,
	}
}

func (p *productService) GetAllProducts(ctx context.Context, keyword, status, order, sortBy string,
	minPrice, maxPrice, categoryId, storeId, limit, page int) ([]entity.Product, helper.Page, error) {
	var (
		totalPage int
	)

	// order by default is ASC
	if order != "DESC" {
		order = "ASC"
	}

	// sorting by default is id
	if sortBy != "name" && sortBy != "stock" && sortBy != "price" && sortBy != "sold" {
		sortBy = "id"
	}

	// limit page by default is 25
	if limit <= 0 {
		limit = 25
	}

	// page view by default is 1
	if page <= 0 {
		page = 1
	}

	// pagination formula
	offset := (page - 1) * limit

	result, count, err := p.Repo.GetAllProducts(ctx, keyword, status, order, sortBy, minPrice, maxPrice, categoryId, storeId, limit, offset)
	if err != nil {
		return []entity.Product{}, helper.Page{}, err
	}

	if int(count)%limit == 0 {
		totalPage = int(count) / limit
	} else {
		totalPage = (int(count) / limit) + 1
	}

	pagination := helper.Page{
		Limit:     limit,
		Total:     int(count),
		Page:      page,
		TotalPage: totalPage,
	}

	return result, pagination, err
}

func (p *productService) GetProductbyId(ctx context.Context, id int) (entity.Product, error) {
	return p.Repo.GetProductById(ctx, id)
}

func (p *productService) InsertProduct(ctx context.Context, product *entity.Product) error {
	var (
		detailError = make(map[string]any)
		status      = "sale"
	)

	if product.Name == "" {
		detailError["name"] = "this field is missing input"
	}

	if product.Stock == nil {
		detailError["stock"] = "this field is missing input"
	} else if *product.Stock < 0 {
		detailError["stock"] = "this field must not negative"
	}

	if product.Price == nil {
		detailError["price"] = "this field is missing input"
	} else if *product.Price < 1 {
		detailError["price"] = "this field must be higher than 0"
	}

	if product.CategoryId == nil {
		detailError["category_id"] = "this field is missing input"
	}

	fmt.Println(detailError)
	if len(detailError) > 0 {
		return productError.ErrProductInput.AttachDetail(detailError)
	}

	if *product.Stock == 0 {
		status = "soldout"
	}

	product.Status = status
	return p.Repo.InsertProduct(ctx, product)
}

func (p *productService) UpdateProduct(ctx context.Context, id int, product *entity.Product) error {
	var (
		detailError = make(map[string]any)
	)

	_, err := p.Repo.GetProductById(ctx, id)
	if err != nil {
		return err
	}

	if product.Name == "" {
		detailError["name"] = "this field is missng input"
	}

	if product.Stock == nil {
		detailError["stock"] = "this field is missng input"
	} else if *product.Stock < 0 {
		detailError["stock"] = "this field must not negative"
	}

	if product.Price == nil {
		detailError["price"] = "this field is missng input"
	} else if *product.Price < 1 {
		detailError["price"] = "this field must be higher than 0"
	}

	if product.CategoryId == nil {
		detailError["price"] = "this field is missng input"
	}

	fmt.Println(detailError)
	if len(detailError) > 0 {
		return productError.ErrProductInput.AttachDetail(detailError)
	}

	// update status product to "sale" if stock more than 0 after updated
	if *product.Stock > 0 {
		product.Status = "sale"
	} else if *product.Stock == 0 {
		product.Status = "soldout"
	}

	return p.Repo.UpdateProduct(ctx, id, product)
}

func (p *productService) SoftDeleteProduct(ctx context.Context, id int) error {
	_, err := p.Repo.GetProductById(ctx, id)
	if err != nil {
		return productError.ErrProductDeleted
	}

	return p.Repo.SoftDeleteProduct(ctx, id)
}

func (p *productService) RestoreDeletedProduct(ctx context.Context, id int) error {
	productCheck, err := p.Repo.GetDeletedProduct(ctx, id)
	if err != nil {
		return err
	} else if !productCheck.DeleteAt.Valid {
		return productError.ErrProductRestored
	}

	return p.Repo.RestoreDeletedProduct(ctx, id)
}

func (p *productService) DeleteProduct(ctx context.Context, id int) error {
	return p.Repo.DeleteProduct(ctx, id)
}

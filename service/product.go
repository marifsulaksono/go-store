package service

import (
	"context"
	"gostore/entity"
	"gostore/helper"
	"gostore/repo"
)

type productService struct {
	Repo repo.ProductRepository
}

type ProductService interface {
	GetAllProducts(ctx context.Context, keyword, status, order, sortBy string, minPrice, maxPrice, categoryId, storeId, limit, page int) ([]entity.Product, error)
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
	minPrice, maxPrice, categoryId, storeId, limit, page int) ([]entity.Product, error) {
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

	return p.Repo.GetAllProducts(ctx, keyword, status, order, sortBy, minPrice, maxPrice, categoryId, storeId, limit, offset)
}

func (p *productService) GetProductbyId(ctx context.Context, id int) (entity.Product, error) {
	return p.Repo.GetProductById(ctx, id)
}

func (p *productService) InsertProduct(ctx context.Context, product *entity.Product) error {
	product.Status = "sale"
	return p.Repo.InsertProduct(ctx, product)
}

func (p *productService) UpdateProduct(ctx context.Context, id int, product *entity.Product) error {
	_, err := p.Repo.GetProductById(ctx, id)
	if err != nil {
		return err
	}
	// update status product to "sale" if stock more than 0 after updated
	if product.Stock > 0 {
		product.Status = "sale"
	} else if product.Stock == 0 {
		product.Status = "soldout"
	}

	return p.Repo.UpdateProduct(ctx, id, product)
}

func (p *productService) SoftDeleteProduct(ctx context.Context, id int) error {
	_, err := p.Repo.GetProductById(ctx, id)
	if err != nil {
		return helper.ErrRecDeleted
	}

	return p.Repo.SoftDeleteProduct(ctx, id)
}

func (p *productService) RestoreDeletedProduct(ctx context.Context, id int) error {
	productCheck, err := p.Repo.GetDeletedProduct(ctx, id)
	if err != nil {
		return err
	} else if !productCheck.DeleteAt.Valid {
		return helper.ErrRecRestored
	}

	return p.Repo.RestoreDeletedProduct(ctx, id)
}

func (p *productService) DeleteProduct(ctx context.Context, id int) error {
	return p.Repo.DeleteProduct(ctx, id)
}

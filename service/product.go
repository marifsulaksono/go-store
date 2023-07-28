package service

import (
	"gostore/entity"
	"gostore/helper"
	"gostore/repo"
)

type ProductService struct {
	Repo repo.ProductRepository
}

func NewProductService(r repo.ProductRepository) *ProductService {
	return &ProductService{
		Repo: r,
	}
}

func (p *ProductService) GetAllProducts() ([]entity.Product, error) {
	products, err := p.Repo.GetAllProducts()
	return products, err
}

func (p *ProductService) GetProductbyId(id int) (entity.Product, error) {
	product, err := p.Repo.GetProductById(id)
	return product, err
}

func (p *ProductService) GetProductbyStatus(s string) ([]entity.Product, error) {
	products, err := p.Repo.GetProductByStatus(s)
	return products, err
}

func (p *ProductService) SearchProduct(keyword, order, sortBy string, minPrice, maxPrice float64, limit, page int) ([]entity.Product, error) {
	if order != "DESC" {
		order = "ASC"
	}

	if sortBy != "name" && sortBy != "stock" && sortBy != "price" && sortBy != "sold" {
		sortBy = "name"
	}

	offset := (page - 1) * limit

	products, err := p.Repo.SearchProduct(keyword, order, sortBy, minPrice, maxPrice, limit, offset)
	return products, err

	// DB.Where("name LIKE ? AND price BETWEEN ? AND ?", "%"+keyword+"%", minPrice, maxPrice).Order(sortBy +
	// 	" " + order).Limit(limit).Offset(offset).Preload("Category").Find(&products).Error
}

func (p *ProductService) InsertProduct(product *entity.Product) error {
	product.Status = "sale"
	err := p.Repo.InsertProduct(product)
	return err
}

func (p *ProductService) UpdateProduct(id int, product *entity.Product) (entity.Product, error) {
	existProduct, err := p.Repo.GetProductById(id)
	if err != nil {
		return entity.Product{}, err
	}
	err = p.Repo.UpdateProduct(id, &existProduct, product)
	if product.Stock > 0 {
		err = p.Repo.ChangeStatusProduct(id, "sale")
		if err != nil {
			return entity.Product{}, err
		} else if existProduct, err = p.Repo.GetProductById(id); err != nil {
			return entity.Product{}, err
		}
	}
	return existProduct, err
}

func (p *ProductService) ChangeStatusProduct(id int, s string) error {
	productCheck, err := p.Repo.GetProductById(id)
	if err != nil {
		return err
	} else if productCheck.Status == s {
		return helper.ErrChangeStatusProduct
	}

	err = p.Repo.ChangeStatusProduct(id, s)
	return err
}

func (p *ProductService) SoftDeleteProduct(id int) error {
	_, err := p.Repo.GetProductById(id)
	if err != nil {
		return helper.ErrRecDeleted
	}

	err = p.Repo.SoftDeleteProduct(id)
	return err
}

func (p *ProductService) RestoreDeletedProduct(id int) error {
	productCheck, err := p.Repo.GetDeletedProduct(id)
	if err != nil {
		return err
	} else if !productCheck.DeleteAt.Valid {
		return helper.ErrRecRestored
	}
	err = p.Repo.RestoreDeletedProduct(id)
	return err
}

func (p *ProductService) DeleteProduct(id int) error {
	err := p.Repo.DeleteProduct(id)
	return err
}

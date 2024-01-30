package repo

import (
	"context"
	"errors"
	"gostore/entity"
	productError "gostore/utils/helper/domain/errorModel"

	"gorm.io/gorm"
)

type productRepository struct {
	DB *gorm.DB
}

type ProductRepository interface {
	GetAllProducts(ctx context.Context, keyword, status, order, sortBy string, minPrice, maxPrice, categoryId, storeId, limit, offset int) ([]entity.Product, int64, error)
	GetProductById(ctx context.Context, id int) (entity.Product, error)
	GetDeletedProduct(ctx context.Context, id int) (entity.Product, error)
	InsertProduct(ctx context.Context, product *entity.Product) error
	UpdateProduct(ctx context.Context, id int, product *entity.Product) error
	ChangeStatusProduct(ctx context.Context, id int, s string) error
	SoftDeleteProduct(ctx context.Context, id int) error
	RestoreDeletedProduct(ctx context.Context, id int) error
	DeleteProduct(ctx context.Context, id int) error
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{
		DB: db,
	}
}

func (p *productRepository) GetAllProducts(ctx context.Context, keyword, status, order, sortBy string,
	minPrice, maxPrice, categoryId, storeId, limit, offset int) ([]entity.Product, int64, error) {
	var (
		products []entity.Product
		db       = p.DB
		count    int64
	)

	if status != "" {
		db = db.Where("status = ?", status)
	}

	if keyword != "" {
		db = db.Where("name LIKE ?", "%"+keyword+"%")
	}

	if categoryId > 0 {
		db = db.Where("category_id = ?", categoryId)
	}

	if storeId > 0 {
		db = db.Where("store_id = ?", storeId)
	}

	if minPrice > 0 {
		db = db.Where("price >= ?", minPrice)
	}

	if maxPrice > 0 {
		db = db.Where("price <= ?", maxPrice)
	}

	err := db.Order(sortBy +
		" " + order).Limit(limit).Offset(offset).Preload("Category").Preload("Store").Find(&products).Count(&count).Error
	return products, count, err
}

func (p *productRepository) GetProductById(ctx context.Context, id int) (entity.Product, error) {
	var product entity.Product
	err := p.DB.Where("id = ?", id).Preload("Category").Preload("Store").First(&product).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Product{}, productError.ErrProductNotFound
		}
		return entity.Product{}, err
	}
	return product, nil
}

func (p *productRepository) GetDeletedProduct(ctx context.Context, id int) (entity.Product, error) {
	var product entity.Product
	err := p.DB.Unscoped().Where("id = ?", id).First(&product).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Product{}, productError.ErrProductNotFound
		}
		return entity.Product{}, err
	}
	return product, nil
}

func (p *productRepository) GetProductByStatus(ctx context.Context, s string) ([]entity.Product, error) {
	var products []entity.Product
	err := p.DB.Where("status = ?", s).Preload("Category").Find(&products).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, productError.ErrProductNotFound
		}
		return nil, err
	}
	return products, nil
}

func (p *productRepository) InsertProduct(ctx context.Context, product *entity.Product) error {
	return p.DB.Create(&product).Error
}

func (p *productRepository) UpdateProduct(ctx context.Context, id int, product *entity.Product) error {
	updateProducts := map[string]interface{}{
		"Name":       product.Name,
		"Stock":      product.Stock,
		"Price":      product.Price,
		"Desc":       product.Desc,
		"Status":     product.Status,
		"CategoryId": product.CategoryId,
	}
	return p.DB.Model(&entity.Product{}).Where("id = ?", id).Updates(updateProducts).Error
}

func (p *productRepository) ChangeStatusProduct(ctx context.Context, id int, s string) error {
	return p.DB.Model(&entity.Product{}).Where("id = ?", id).Update("status", s).Error
}

func (p *productRepository) SoftDeleteProduct(ctx context.Context, id int) error {
	return p.DB.Where("id = ?", id).Delete(&entity.Product{}).Error
}

func (p *productRepository) RestoreDeletedProduct(ctx context.Context, id int) error {
	return p.DB.Unscoped().Model(&entity.Product{}).Where("id = ?", id).Update("delete_at", gorm.DeletedAt{}).Error
}

func (p *productRepository) DeleteProduct(ctx context.Context, id int) error {
	return p.DB.Unscoped().Where("id = ?", id).Delete(&entity.Product{}).Error
}

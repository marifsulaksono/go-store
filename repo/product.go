package repo

import (
	"gostore/entity"

	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		DB: db,
	}
}

func (p *ProductRepository) GetAllProducts() ([]entity.Product, error) {
	var products []entity.Product
	err := p.DB.Preload("Category").Preload("Store").Find(&products).Error
	return products, err
}

func (p *ProductRepository) GetProductById(id int) (entity.Product, error) {
	var product entity.Product
	err := p.DB.Where("id = ?", id).Preload("Category").Preload("Store").First(&product).Error
	return product, err
}

func (p *ProductRepository) GetDeletedProduct(id int) (entity.Product, error) {
	var product entity.Product
	err := p.DB.Unscoped().Where("id = ?", id).Preload("Category", "id NOT IN (?)", "cancelled").First(&product).Error
	return product, err
}

func (p *ProductRepository) GetProductByStatus(s string) ([]entity.Product, error) {
	var products []entity.Product
	err := p.DB.Where("status = ?", s).Preload("Category").Find(&products).Error
	return products, err
}

func (p *ProductRepository) GetProductOnStore(id int) ([]entity.ProductResponse, error) {
	var products []entity.ProductResponse
	err := p.DB.Where("store_id = ?", id).Preload("Category").Find(&products).Error
	return products, err
}

func (p *ProductRepository) SearchProduct(keyword, order, sortBy string, minPrice, maxPrice, categoryId, storeId, limit, offset int) ([]entity.Product, error) {
	var (
		products []entity.Product
		db       = p.DB
	)

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
		" " + order).Limit(limit).Offset(offset).Preload("Category").Preload("Store").Find(&products).Error

	// err := p.DB.Where("name LIKE ? AND price BETWEEN ? AND ?", "%"+keyword+"%", minPrice, maxPrice).Order(sortBy +
	// 	" " + order).Limit(limit).Offset(offset).Preload("Category").Find(&products).Error
	return products, err
}

func (p *ProductRepository) InsertProduct(product *entity.Product) error {
	err := p.DB.Create(&product).Error
	return err
}

func (p *ProductRepository) UpdateProduct(id int, model, product *entity.Product) error {
	err := p.DB.Model(model).Where("id = ?", id).Updates(product).Error
	return err
}

func (p *ProductRepository) UpdateStockandSold(id int, product *entity.Product) error {
	err := p.DB.Save(&product).Error
	// err := p.DB.Model(&entity.Product{}).Select("Stock", "Sold").Where("id = ?", id).Updates(product).Error
	return err
}

func (p *ProductRepository) ChangeStatusProduct(id int, s string) error {
	err := p.DB.Model(&entity.Product{}).Where("id = ?", id).Update("status", s).Error
	return err
}

func (p *ProductRepository) SoftDeleteProduct(id int) error {
	err := p.DB.Where("id = ?", id).Delete(&entity.Product{}).Error
	return err
}

func (p *ProductRepository) RestoreDeletedProduct(id int) error {
	err := p.DB.Unscoped().Model(&entity.Product{}).Where("id = ?", id).Update("delete_at", gorm.DeletedAt{}).Error
	return err
}

func (p *ProductRepository) DeleteProduct(id int) error {
	err := p.DB.Unscoped().Where("id = ?", id).Delete(&entity.Product{}).Error
	return err
}

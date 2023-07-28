package repo_test

import (
	"fmt"
	"gostore/entity"
	"testing"

	_ "github.com/stretchr/testify"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func createProduct(t *testing.T) entity.Product {
	arg := entity.Product{
		Name:       "HP Iphone",
		Stock:      10,
		Price:      15000000,
		Status:     "sale",
		CategoryId: 3,
	}

	err := testDB.InsertProduct(&arg)
	var product entity.Product
	_ = testDB.DB.Last(&product)
	fmt.Println(product)

	require.NoError(t, err)
	require.NotEmpty(t, product)

	require.Equal(t, arg.Name, product.Name)
	require.Equal(t, arg.Stock, product.Stock)
	require.Equal(t, arg.Price, product.Price)
	require.Equal(t, arg.Status, product.Status)
	require.Equal(t, arg.CategoryId, product.CategoryId)

	require.NotZero(t, product.Id)

	return product
}

func TestInsertProduct(t *testing.T) {
	createProduct(t)
}

func TestGetProduct(t *testing.T) {
	product1 := createProduct(t)
	product2, err := testDB.GetProductById(product1.Id)

	require.NoError(t, err)
	require.NotEmpty(t, product2)

	require.Equal(t, product1.Id, product2.Id)
	require.Equal(t, product1.Name, product2.Name)
	require.Equal(t, product1.Stock, product2.Stock)
	require.Equal(t, product1.Price, product2.Price)
	require.Equal(t, product1.Status, product2.Status)
	require.Equal(t, product1.CategoryId, product2.CategoryId)
}

func TestUpdateProduct(t *testing.T) {
	newProduct := createProduct(t)
	arg := entity.Product{
		// get the newProduct prop if no changes
		Id:         newProduct.Id, // no change
		Name:       "Cleo Botol",
		Stock:      100,
		Price:      4000,
		Status:     newProduct.Status,     // no change
		CategoryId: newProduct.CategoryId, // no change
	}

	err := testDB.UpdateProduct(newProduct.Id, &arg)
	require.NoError(t, err)

	product, err := testDB.GetProductById(newProduct.Id)
	require.NoError(t, err)
	require.NotEmpty(t, product)

	require.Equal(t, arg.Id, product.Id)
	require.Equal(t, arg.Name, product.Name)
	require.Equal(t, arg.Stock, product.Stock)
	require.Equal(t, arg.Price, product.Price)
	require.Equal(t, newProduct.Status, product.Status)
	require.Equal(t, newProduct.CategoryId, product.CategoryId)
}

func TestDeleteProduct(t *testing.T) {
	newProduct := createProduct(t)
	err := testDB.DeleteProduct(newProduct.Id)
	require.NoError(t, err)

	checkProduct, err := testDB.GetProductById(newProduct.Id)
	require.Error(t, err)
	require.EqualError(t, err, gorm.ErrRecordNotFound.Error())
	require.Empty(t, checkProduct)
}

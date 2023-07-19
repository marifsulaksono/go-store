package repo_test

import (
	"fmt"
	"gostore/entity"
	"testing"

	_ "github.com/stretchr/testify"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func createItem(t *testing.T) entity.Item {
	arg := entity.Item{
		Name:       "HP Iphone",
		Stock:      10,
		Price:      15000000,
		IsSale:     1,
		CategoryId: 3,
	}

	err := testDB.InsertItem(&arg)
	var item entity.Item
	_ = testDB.DB.Last(&item)
	fmt.Println(item)

	require.NoError(t, err)
	require.NotEmpty(t, item)

	require.Equal(t, arg.Name, item.Name)
	require.Equal(t, arg.Stock, item.Stock)
	require.Equal(t, arg.Price, item.Price)
	require.Equal(t, arg.IsSale, item.IsSale)
	require.Equal(t, arg.CategoryId, item.CategoryId)

	require.NotZero(t, item.Id)

	return item
}

func TestInsertItem(t *testing.T) {
	createItem(t)
}

func TestGetItem(t *testing.T) {
	item1 := createItem(t)
	item2, err := testDB.GetItemById(item1.Id)

	require.NoError(t, err)
	require.NotEmpty(t, item2)

	require.Equal(t, item1.Id, item2.Id)
	require.Equal(t, item1.Name, item2.Name)
	require.Equal(t, item1.Stock, item2.Stock)
	require.Equal(t, item1.Price, item2.Price)
	require.Equal(t, item1.IsSale, item2.IsSale)
	require.Equal(t, item1.CategoryId, item2.CategoryId)
}

func TestUpdateITem(t *testing.T) {
	newItem := createItem(t)
	arg := entity.Item{
		// get the newItem prop if no changes
		Id:         newItem.Id, // no change
		Name:       "Cleo Botol",
		Stock:      100,
		Price:      4000,
		IsSale:     newItem.IsSale,     // no change
		CategoryId: newItem.CategoryId, // no change
	}

	err := testDB.UpdateItem(newItem.Id, &arg)
	require.NoError(t, err)

	item, err := testDB.GetItemById(newItem.Id)
	require.NoError(t, err)
	require.NotEmpty(t, item)

	require.Equal(t, arg.Id, item.Id)
	require.Equal(t, arg.Name, item.Name)
	require.Equal(t, arg.Stock, item.Stock)
	require.Equal(t, arg.Price, item.Price)
	require.Equal(t, newItem.IsSale, item.IsSale)
	require.Equal(t, newItem.CategoryId, item.CategoryId)
}

func TestDeleteItem(t *testing.T) {
	newItem := createItem(t)
	err := testDB.DeleteItem(newItem.Id)
	require.NoError(t, err)

	checkItem, err := testDB.GetItemById(newItem.Id)
	require.Error(t, err)
	require.EqualError(t, err, gorm.ErrRecordNotFound.Error())
	require.Empty(t, checkItem)
}

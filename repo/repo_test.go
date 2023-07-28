package repo_test

import (
	"gostore/repo"
	"os"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var testDB repo.ProductRepository

func TestMain(m *testing.M) {
	conn, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/db_store?parseTime=True"), &gorm.Config{})
	if err != nil {
		panic("Connection failed!")
	}

	testDB = *repo.NewProductRepository(conn)
	os.Exit(m.Run())
}

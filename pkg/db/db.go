package db

import (
	"log"

	"github.com/14jasimmtp/product-svc/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Handler struct {
    DB *gorm.DB
}

func Connection(DB_URL string) Handler{
	db,err:=gorm.Open(postgres.Open(DB_URL),&gorm.Config{})

	if err != nil{
		log.Fatal(err)
	}

	db.AutoMigrate(&models.StockDecreaseLog{})
	db.AutoMigrate(&models.Product{})


	return Handler{db}
}
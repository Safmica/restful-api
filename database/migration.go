package database

import (
	"GDSC-PROJECT/models/entity"
	"fmt"
	"log"
)

func DBMigration() {
	err := DB.AutoMigrate(&entity.Product{}, &entity.Category{}, &entity.Stock{}, &entity.Category{})
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Database migrated successfully")
}

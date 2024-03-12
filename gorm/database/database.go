package database

import (
	"backend/gorm/entities"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPostgres() gorm.DB {
	dsn := "host=localhost user=postgres password=password dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entities.Product{}, &entities.User{})

	fmt.Println("Koneksi database sukses")

	return *db
}

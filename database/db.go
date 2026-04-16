package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "host=localhost user=postgres password=postgres " +
		"dbname=postgres " +
		"port=5432 " +
		"sslmode=disable"

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Не удалось подключиться к базе данных!")
	}

	DB = database
	fmt.Println("Успешное подключение к PostgreSQL")
}

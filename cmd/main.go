package main

import (
	"log"

	"github.com/gui-costads/carteira-app/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=carteira port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	err = db.AutoMigrate(&models.Usuario{}, &models.Categoria{}, &models.Transacao{}, &models.Orcamento{})
	if err != nil {
		log.Fatalf("Error migrating models: %v", err)
	}

}

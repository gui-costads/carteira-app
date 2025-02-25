package main

import (
	"log"

	"github.com/gui-costads/carteira-app/internal/config"
	"github.com/gui-costads/carteira-app/internal/models"
)

func main() {
	db, err := config.DatabaseConnection()
	if err != nil {
		log.Fatal("Falha ao conectar ao banco de dados: %v", err)
	}

	err = db.AutoMigrate(
		&models.Usuario{},
		&models.Orcamento{},
		&models.Categoria{},
		&models.Orcamento{},
	)
	if err != nil {
		log.Fatal("Falha ao gerar modelos: %v", err)
	}
	log.Println("Modelos gerados com sucesso")

}

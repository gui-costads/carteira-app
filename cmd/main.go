package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/gui-costads/carteira-app/internal/config"
	"github.com/gui-costads/carteira-app/internal/controller"
	"github.com/gui-costads/carteira-app/internal/models"
	"github.com/gui-costads/carteira-app/internal/repository/orcamentorepository"
	"github.com/gui-costads/carteira-app/internal/repository/usuariorepository"
	"github.com/gui-costads/carteira-app/internal/routes"
	orcamentoservice "github.com/gui-costads/carteira-app/internal/service/orcamento"
	usuarioservice "github.com/gui-costads/carteira-app/internal/service/usuario"
)

func main() {
	db, err := config.DatabaseConnection()
	if err != nil {
		log.Fatalf("Falha ao conectar ao banco de dados: %v", err)
	}

	err = db.AutoMigrate(
		&models.Usuario{},
		&models.Orcamento{},
		&models.Categoria{},
		&models.Orcamento{},
	)
	if err != nil {
		log.Fatalf("Falha ao gerar modelos: %v", err)
	}
	log.Println("Modelos gerados com sucesso")

	usuarioRepo := usuariorepository.NewUsuarioRepositoryImpl(db)
	usuarioService := usuarioservice.NewUsuarioService(usuarioRepo)
	usuarioController := controller.NewUsuarioController(usuarioService)

	orcamentoRepo := orcamentorepository.NewUsuarioRepositoryImpl(db)
	orcamentoService := orcamentoservice.NewOrcamentoService(orcamentoRepo)
	orcamentoController := controller.NewOrcamentoController(orcamentoService)

	router := gin.Default()

	api := router.Group("/api/v1")

	{
		routes.SetupUsuarioRoutes(api, usuarioController)
		routes.SetupOrcamentoRoutes(api, orcamentoController)
	}
	log.Println("🚀 Servidor iniciado na porta 8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Falha ao iniciar servidor: %v", err)
	}
}

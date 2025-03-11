package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/gui-costads/carteira-app/internal/auth"
	"github.com/gui-costads/carteira-app/internal/config"
	"github.com/gui-costads/carteira-app/internal/controller"
	"github.com/gui-costads/carteira-app/internal/models"
	"github.com/gui-costads/carteira-app/internal/repository/categoriarepository"
	"github.com/gui-costads/carteira-app/internal/repository/orcamentorepository"
	"github.com/gui-costads/carteira-app/internal/repository/transacaorepository"
	"github.com/gui-costads/carteira-app/internal/repository/usuariorepository"
	"github.com/gui-costads/carteira-app/internal/routes"
	categoriaservice "github.com/gui-costads/carteira-app/internal/service/categoria"
	orcamentoservice "github.com/gui-costads/carteira-app/internal/service/orcamento"
	transacaoservice "github.com/gui-costads/carteira-app/internal/service/transacao"
	usuarioservice "github.com/gui-costads/carteira-app/internal/service/usuario"
)

func main() {
	cfg := config.Load()
	db, err := cfg.DatabaseConnection()
	authService := auth.NewAuthService(cfg)
	if err != nil {
		log.Fatalf("Falha ao conectar ao banco de dados: %v", err)
	}

	err = db.AutoMigrate(
		&models.Usuario{},
		&models.Transacao{},
		&models.Categoria{},
		&models.Orcamento{},
	)
	if err != nil {
		log.Fatalf("Falha ao gerar modelos: %v", err)
	}
	log.Println("Modelos gerados com sucesso")

	usuarioRepo := usuariorepository.NewUsuarioRepository(db)
	usuarioService := usuarioservice.NewUsuarioService(usuarioRepo)
	usuarioController := controller.NewUsuarioController(usuarioService, authService)

	orcamentoRepo := orcamentorepository.NewOrcamentoRepository(db)
	orcamentoService := orcamentoservice.NewOrcamentoService(orcamentoRepo)
	orcamentoController := controller.NewOrcamentoController(orcamentoService)

	categoriaRepo := categoriarepository.NewCategoriaRepository(db)
	categoriaService := categoriaservice.NewCategoriaService(categoriaRepo)
	categoriaController := controller.NewCategoriaController(categoriaService)

	transacaoRepo := transacaorepository.NewTransacaoRepository(db)
	transacaoService := transacaoservice.NewTransacaoService(transacaoRepo)
	transacaoController := controller.NewTransacaoController(transacaoService)

	router := gin.Default()

	api := router.Group("/api/v1")

	{
		router.Use(authService.AuthenticationMiddleware())
		routes.SetupUsuarioRoutes(api, usuarioController, authService)
		routes.SetupOrcamentoRoutes(api, orcamentoController)
		routes.SetupCategoriaRoutes(api, categoriaController)
		routes.SetupTransacaoRoutes(api, transacaoController)
	}
	log.Println("🚀 Servidor iniciado na porta 8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Falha ao iniciar servidor: %v", err)
	}
}

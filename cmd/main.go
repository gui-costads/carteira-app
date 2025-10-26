package main

import (
	"log"

	"github.com/gin-contrib/cors"
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

	// Configurar CORS
	configCors := cors.DefaultConfig()
	configCors.AllowOrigins = []string{"http://localhost:4200", "http://localhost:5173"}
	configCors.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	configCors.AllowHeaders = []string{"Content-Type", "Authorization"}

	r := gin.Default()
	r.Use(cors.New(configCors))

	// Rotas públicas (sem autenticação)
	public := r.Group("/usuarios")
	public.POST("/login", usuarioController.Login)
	public.POST("", usuarioController.CriarUsuario)

	// Rotas privadas (precisam de autenticação)
	private := r.Group("/usuarios")
	private.Use(authService.AuthenticationMiddleware())

	private.GET("", usuarioController.BuscarTodosUsuarios)
	private.GET("/:id", usuarioController.BuscarPorId)
	private.PUT("/:id", usuarioController.AtualizarUsuario)
	private.DELETE("/:id", usuarioController.DeletarUsuario)

	api := r.Group("/api/v1")
	{
		routes.SetupOrcamentoRoutes(api, orcamentoController, authService)
		routes.SetupCategoriaRoutes(api, categoriaController, authService)
		routes.SetupTransacaoRoutes(api, transacaoController, authService)
	}
	log.Println("🚀 Servidor iniciado na porta 8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Falha ao iniciar servidor: %v", err)
	}
}

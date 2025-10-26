package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gui-costads/carteira-app/internal/auth"
	"github.com/gui-costads/carteira-app/internal/controller"
)

func SetupUsuarioRoutes(router *gin.RouterGroup, usuarioController *controller.UsuarioController, authService *auth.AuthService) {
	// Rotas públicas (sem autenticação)
	public := router.Group("/usuarios")
	public.POST("/login", usuarioController.Login)
	public.POST("", usuarioController.CriarUsuario)

	// Rotas privadas (precisam de autenticação)
	private := router.Group("/usuarios")
	private.Use(authService.AuthenticationMiddleware())

	private.GET("", usuarioController.BuscarTodosUsuarios)
	private.GET("/:id", usuarioController.BuscarPorId)
	private.PUT("/:id", usuarioController.AtualizarUsuario)
	private.DELETE("/:id", usuarioController.DeletarUsuario)
}

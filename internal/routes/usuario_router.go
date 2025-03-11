package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gui-costads/carteira-app/internal/auth"
	"github.com/gui-costads/carteira-app/internal/controller"
)

func SetupUsuarioRoutes(router *gin.RouterGroup, usuarioController *controller.UsuarioController, authService *auth.AuthService) {
	router.POST("/usuarios/login", usuarioController.Login)
	router.POST("/usuarios", usuarioController.CriarUsuario)
	{
		usuarios := router.Group("/usuarios")

		usuarios.Use(authService.AuthenticationMiddleware())

		usuarios.GET("", usuarioController.BuscarTodosUsuarios)

		usuarios.GET("/:id", usuarioController.BuscarPorId)

		usuarios.PUT("/:id", usuarioController.AtualizarUsuario)

		usuarios.DELETE("/:id", usuarioController.DeletarUsuario)
	}
}

package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gui-costads/carteira-app/internal/controller"
)

func SetupUsuarioRoutes(router *gin.RouterGroup, usuarioController *controller.UsuarioController) {
	usuarios := router.Group("/usuarios")
	{
		usuarios.GET("", usuarioController.BuscarTodosUsuarios)

		usuarios.GET("/:id", usuarioController.BuscarPorId)

		usuarios.POST("", usuarioController.CriarUsuario)

		usuarios.PUT("/:id", usuarioController.AtualizarUsuario)

		usuarios.DELETE("/:id", usuarioController.DeletarUsuario)
	}
}

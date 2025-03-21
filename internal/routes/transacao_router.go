package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gui-costads/carteira-app/internal/auth"
	"github.com/gui-costads/carteira-app/internal/controller"
)

func SetupTransacaoRoutes(router *gin.RouterGroup, transacaoController *controller.TransacaoController, authService *auth.AuthService) {
	transacoes := router.Group("/transacoes")

	transacoes.Use(authService.AuthenticationMiddleware())
	{
		transacoes.GET("", transacaoController.BuscarTodasTransacoes)

		transacoes.GET("/:id", transacaoController.BuscarTransacaoPorID)

		transacoes.POST("", transacaoController.CriarTransacao)

		transacoes.PUT("/:id", transacaoController.AtualizarTransacao)

		transacoes.DELETE("/:id", transacaoController.DeletarTransacao)
	}
}

package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gui-costads/carteira-app/internal/controller"
)

func SetupOrcamentoRoutes(router *gin.RouterGroup, orcamentoController *controller.OrcamentoController) {
	orcamentos := router.Group("/orcamentos")
	{
		orcamentos.GET("", orcamentoController.BuscarTodosOrcamentos)

		orcamentos.GET("/:id", orcamentoController.BuscarOrcamentoPorID)

		orcamentos.POST("", orcamentoController.CriarOrcamento)

		orcamentos.PUT("/:id", orcamentoController.AtualizarOrcamento)

		orcamentos.DELETE("/:id", orcamentoController.DeletarOrcamento)
	}
}

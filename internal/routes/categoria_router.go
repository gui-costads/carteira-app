package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gui-costads/carteira-app/internal/controller"
)

func SetupCategoriaRoutes(router *gin.RouterGroup, categoriaController *controller.CategoriaController) {
	categorias := router.Group("/categorias")
	{
		categorias.GET("", categoriaController.BuscarTodasCategorias)

		categorias.GET("/:id", categoriaController.BuscarPorID)

		categorias.POST("", categoriaController.CriarCategoria)

		categorias.PUT("/:id", categoriaController.AtualizarCategoria)

		categorias.DELETE("/:id", categoriaController.DeletarCategoria)
	}
}

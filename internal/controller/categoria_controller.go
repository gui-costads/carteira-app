package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gui-costads/carteira-app/internal/data/categoriadto"
	response "github.com/gui-costads/carteira-app/internal/data/http"
	categoriaService "github.com/gui-costads/carteira-app/internal/service/categoria"
)

type CategoriaController struct {
	categoriaService categoriaService.CategoriaService
}

func NewCategoriaController(categoriaService categoriaService.CategoriaService) *CategoriaController {
	return &CategoriaController{categoriaService: categoriaService}
}

func (controller *CategoriaController) BuscarTodasCategorias(ctx *gin.Context) {
	data, err := controller.categoriaService.BuscarTodasCategorias()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:    500,
			Message: err.Error(),
		})
		return
	}
	res := response.Response{
		Code:   200,
		Status: "OK",
		Data:   data,
	}

	ctx.JSON(http.StatusOK, res)
}

func (controller *CategoriaController) BuscarPorID(ctx *gin.Context) {
	id := ctx.Param("id")
	uid, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}
	idu := uint(uid)

	data, err := controller.categoriaService.BuscarCategoriaPorID(idu)
	if err != nil {
		ctx.JSON(http.StatusNotFound, response.ErrorResponse{
			Code:    404,
			Message: err.Error(),
		})
		return
	}

	res := response.Response{
		Code:   200,
		Status: "OK",
		Data:   data,
	}
	ctx.JSON(http.StatusOK, res)
}

func (controller *CategoriaController) CriarCategoria(ctx *gin.Context) {
	req := categoriadto.CriarCategoriaRequest{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    400,
			Message: "Dados inválidos",
		})
		return
	}

	data, err := controller.categoriaService.CriarCategoria(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	res := response.Response{
		Code:   201,
		Status: "Created",
		Data:   data,
	}

	ctx.JSON(http.StatusCreated, res)
}

func (controller *CategoriaController) AtualizarCategoria(ctx *gin.Context) {
	req := categoriadto.AtualizarCategoriaRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    400,
			Message: "Dados inválidos",
		})
		return
	}

	categoriaId := ctx.Param("id")
	id, err := strconv.Atoi(categoriaId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}
	uid := uint(id)

	data, err := controller.categoriaService.AtualizarCategoria(uid, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:    400,
			Message: "Dados inválidos",
		})
		return
	}

	res := response.Response{
		Code:   200,
		Status: "OK",
		Data:   data,
	}
	ctx.JSON(http.StatusOK, res)
}

func (controller *CategoriaController) DeletarCategoria(ctx *gin.Context) {
	categoriaId := ctx.Param("id")
	id, err := strconv.Atoi(categoriaId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}
	uid := uint(id)

	err = controller.categoriaService.DeletarCategoria(uid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	res := response.Response{
		Code:   204,
		Status: "No Content",
		Data:   nil,
	}
	ctx.JSON(http.StatusNoContent, res)
}

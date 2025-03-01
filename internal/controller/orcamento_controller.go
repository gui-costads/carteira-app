package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	response "github.com/gui-costads/carteira-app/internal/data/http"

	"github.com/gui-costads/carteira-app/internal/data/orcamentodto"
	orcamentoservice "github.com/gui-costads/carteira-app/internal/service/orcamento"
)

type OrcamentoController struct {
	orcamentoservice orcamentoservice.OrcamentoService
}

func NewOrcamentoController(orcamentoService orcamentoservice.OrcamentoService) *OrcamentoController {
	return &OrcamentoController{orcamentoservice: orcamentoService}
}

func (controller *OrcamentoController) BuscarTodosOrcamentos(ctx *gin.Context) {
	data, err := controller.orcamentoservice.BuscarTodosOrcamentos()

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

func (controller *OrcamentoController) BuscarOrcamentoPorID(ctx *gin.Context) {
	orcamentoID := ctx.Param("id")
	id, err := strconv.Atoi(orcamentoID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}
	uid := uint(id)

	data, err := controller.orcamentoservice.BuscarOrcamentoPorID(uid)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "orçamento não encontrado" {
			statusCode = http.StatusNotFound
		}
		ctx.JSON(statusCode, response.ErrorResponse{
			Code:    statusCode,
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

func (controller *OrcamentoController) CriarOrcamento(ctx *gin.Context) {
	req := orcamentodto.CriarOrcamentoRequest{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    400,
			Message: "Dados inválidos",
		})
		return
	}

	data, err := controller.orcamentoservice.CriarOrcamento(req)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "já existe um orçamento para esta combinação de usuário, período e categoria" {
			statusCode = http.StatusConflict
		}
		ctx.JSON(statusCode, response.ErrorResponse{
			Code:    statusCode,
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

func (controller *OrcamentoController) AtualizarOrcamento(ctx *gin.Context) {
	req := orcamentodto.AtualizarOrcamentoRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    400,
			Message: "Dados inválidos",
		})
		return
	}

	orcamentoID := ctx.Param("id")
	id, err := strconv.Atoi(orcamentoID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}
	uid := uint(id)

	data, err := controller.orcamentoservice.AtualizarOrcamento(uid, req)
	if err != nil {
		statusCode := http.StatusInternalServerError
		switch err.Error() {
		case "orçamento não encontrado":
			statusCode = http.StatusNotFound
		case "combinação única de usuário/período/categoria já existe":
			statusCode = http.StatusConflict
		}
		ctx.JSON(statusCode, response.ErrorResponse{
			Code:    statusCode,
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

func (controller *OrcamentoController) DeletarOrcamento(ctx *gin.Context) {
	orcamentoID := ctx.Param("id")
	id, err := strconv.Atoi(orcamentoID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}
	uid := uint(id)

	err = controller.orcamentoservice.DeletarOrcamento(uid)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "orçamento não encontrado" {
			statusCode = http.StatusNotFound
		}
		ctx.JSON(statusCode, response.ErrorResponse{
			Code:    statusCode,
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

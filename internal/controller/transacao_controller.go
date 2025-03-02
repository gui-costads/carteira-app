package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	response "github.com/gui-costads/carteira-app/internal/data/http"

	"github.com/gui-costads/carteira-app/internal/data/transacaodto"
	transacaoservice "github.com/gui-costads/carteira-app/internal/service/transacao"
)

type TransacaoController struct {
	transacaoservice transacaoservice.TransacaoService
}

func NewTransacaoController(transacaoService transacaoservice.TransacaoService) *TransacaoController {
	return &TransacaoController{transacaoservice: transacaoService}
}

func (controller *TransacaoController) BuscarTodasTransacoes(ctx *gin.Context) {
	data, err := controller.transacaoservice.BuscarTodasTransacoes()

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

func (controller *TransacaoController) BuscarTransacaoPorID(ctx *gin.Context) {
	transacaoID := ctx.Param("id")
	id, err := strconv.Atoi(transacaoID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}
	uid := uint(id)

	data, err := controller.transacaoservice.BuscarTransacaoPorID(uid)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "transacao não encontrada" {
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

func (controller *TransacaoController) CriarTransacao(ctx *gin.Context) {
	req := transacaodto.CriarTransacaoRequest{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    400,
			Message: "Dados inválidos",
		})
		return
	}

	data, err := controller.transacaoservice.CriarTransacao(req)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "já existe uma transacao para esta combinação de usuário, período e categoria" {
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

func (controller *TransacaoController) AtualizarTransacao(ctx *gin.Context) {
	req := transacaodto.AtualizarTransacaoRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    400,
			Message: "Dados inválidos",
		})
		return
	}

	transacaoID := ctx.Param("id")
	id, err := strconv.Atoi(transacaoID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}
	uid := uint(id)

	data, err := controller.transacaoservice.AtualizarTransacao(uid, req)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "transacao não encontrada" {
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

func (controller *TransacaoController) DeletarTransacao(ctx *gin.Context) {
	transacaoID := ctx.Param("id")
	id, err := strconv.Atoi(transacaoID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}
	uid := uint(id)

	err = controller.transacaoservice.DeletarTransacao(uid)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "transacao não encontrada" {
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

package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	response "github.com/gui-costads/carteira-app/internal/data/http"

	usuarioservice "github.com/gui-costads/carteira-app/internal/service/usuario"

	"github.com/gui-costads/carteira-app/internal/data/usuariodto"
)

type UsuarioController struct {
	usuarioservice usuarioservice.UsuarioService
}

func NewUsuarioController(usuarioService usuarioservice.UsuarioService) *UsuarioController {
	return &UsuarioController{usuarioservice: usuarioService}
}

func (controller *UsuarioController) BuscarTodosUsuarios(ctx *gin.Context) {
	data, err := controller.usuarioservice.BuscarTodosUsuarios()

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

func (controller *UsuarioController) BuscarPorId(ctx *gin.Context) {
	usuarioId := ctx.Param("id")
	id, err := strconv.Atoi(usuarioId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}
	uid := uint(id)

	data, err := controller.usuarioservice.BuscarUsuarioPorID(uid)
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

func (controller *UsuarioController) CriarUsuario(ctx *gin.Context) {
	req := usuariodto.CriarUsuarioRequest{}

	ctx.ShouldBindJSON(&req)

	data, err := controller.usuarioservice.CriarUsuario(req)
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
	ctx.JSON(http.StatusOK, res)
}

func (controller *UsuarioController) AtualizarUsuario(ctx *gin.Context) {
	req := usuariodto.AtualizarUsuarioRequest{}
	ctx.ShouldBindJSON(&req)

	usuarioId := ctx.Param("id")
	id, err := strconv.Atoi(usuarioId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	uid := uint(id)

	data, err := controller.usuarioservice.AtualizarUsuario(uid, req)
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

func (controller *UsuarioController) DeletarUsuario(ctx *gin.Context) {
	usuarioId := ctx.Param("id")
	id, err := strconv.Atoi(usuarioId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}
	uid := uint(id)

	err = controller.usuarioservice.DeletarUsuario(uid)
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
		Data:   nil,
	}
	ctx.JSON(http.StatusOK, res)
}

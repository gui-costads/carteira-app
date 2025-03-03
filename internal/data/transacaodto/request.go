package transacaodto

import "time"

type CriarTransacaoRequest struct {
	Descricao       string    `json:"descricao,omitempty"`
	Valor           float64   `json:"valor" binding:"required,min=0"`
	Data            time.Time `json:"data" binding:"required"`
	TipoDeTransacao string    `json:"tipoDeTransacao" binding:"required,oneof=acrescimo decrescimo"`
	UsuarioID       uint      `json:"UsuarioID" binding:"required"`
	CategoriaID     uint      `json:"CategoriaID" binding:"required"`
}

type AtualizarTransacaoRequest struct {
	Descricao       *string    `json:"descricao,omitempty"`
	Valor           *float64   `json:"valor,omitempty" binding:"omitempty,min=0"`
	Data            *time.Time `json:"data,omitempty"`
	TipoDeTransacao *string    `json:"tipoDeTransacao,omitempty" binding:"omitempty,oneof=acrescimo decrescimo"`
	UsuarioID       *uint      `json:"UsuarioID,omitempty"`
	CategoriaID     *uint      `json:"CategoriaID,omitempty"`
}

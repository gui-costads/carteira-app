package transacaodto

import "time"

type CriarTransacaoRequest struct {
	Descricao       string    `json:"descricao,omitempty"`
	Valor           float64   `json:"valor" binding:"required,min=0"`
	Data            time.Time `json:"data" binding:"required"`
	TipoDeTransacao string    `json:"tipo_de_transacao" binding:"required,oneof=acrescimo decrescimo"`
	UsuarioID       uint      `json:"usuario_id" binding:"required"`
	CategoriaID     uint      `json:"categoria_id" binding:"required"`
}

type AtualizarTransacaoRequest struct {
	Descricao       *string    `json:"descricao,omitempty"`
	Valor           *float64   `json:"valor,omitempty" binding:"omitempty,min=0"`
	Data            *time.Time `json:"data,omitempty"`
	TipoDeTransacao *string    `json:"tipo_de_transacao,omitempty" binding:"omitempty,oneof=acrescimo decrescimo"`
	UsuarioID       *uint      `json:"usuario_id,omitempty"`
	CategoriaID     *uint      `json:"categoria_id,omitempty"`
}

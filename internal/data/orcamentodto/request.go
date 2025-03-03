package orcamentodto

type CriarOrcamentoRequest struct {
	Saldo       float64 `json:"saldo" validate:"required,min=0"`
	Periodo     string  `json:"periodo" validate:"required,oneof=semanal mensal anual"`
	UsuarioID   uint    `json:"UsuarioID" validate:"required"`
	CategoriaID uint    `json:"CategoriaID" validate:"required"`
}

type AtualizarOrcamentoRequest struct {
	Saldo       *float64 `json:"saldo,omitempty" validate:"omitempty,min=0"`
	Periodo     *string  `json:"periodo,omitempty" validate:"omitempty,oneof=semanal mensal anual"`
	CategoriaID *uint    `json:"CategoriaID,omitempty" validate:"omitempty"`
}

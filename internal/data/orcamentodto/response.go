package orcamentodto

type ResponseOrcamento struct {
	ID          uint    `json:"id"`
	Saldo       float64 `json:"saldo"`
	Periodo     string  `json:"periodo"`
	UsuarioID   uint    `json:"usuarioID"`
	CategoriaID uint    `json:"categoriaID"`
}

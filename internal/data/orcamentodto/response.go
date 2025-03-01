package orcamentodto

type ResponseOrcamento struct {
	ID          uint    `json:"id"`
	Saldo       float64 `json:"saldo"`
	Periodo     string  `json:"periodo"`
	UsuarioID   uint    `json:"usuario_id"`
	CategoriaID uint    `json:"categoria_id"`
}

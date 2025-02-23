package models

type Orcamento struct {
	Saldo     float64
	Periodo   string
	Usuario   Usuario
	Categoria Categoria
}

package models

type Categoria struct {
	Nome          string
	TipoDeReceita string
	Transacao     []Transacao
	Orcamento     []Orcamento
}

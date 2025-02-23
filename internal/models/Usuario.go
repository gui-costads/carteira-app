package models

type Usuario struct {
	Nome      string
	Sobrenome string
	Email     string
	Senha     string
	Orcamento []Orcamento
	Transacao []Transacao
	Categoria []Categoria
}

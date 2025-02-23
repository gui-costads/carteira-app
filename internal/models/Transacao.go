package models

import "time"

type Transacao struct {
	Descricao       string
	Valor           float64
	Data            time.Time
	TipoDeTransacao string
	Usuario         Usuario
	Categoria       Categoria
}

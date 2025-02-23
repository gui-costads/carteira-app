package models

import "time"

type Transacao struct {
	Descricao       string    `gorm:"type:text"`
	Valor           float64   `gorm:"type:decimal(10,2);not null; check:valor >= 0"`
	Data            time.Time `gorm:"not null"`
	TipoDeTransacao string    `gorm:"not null; check:tipo_de_transacao IN ('acrescimo', 'decrescimo')"`
	UsuarioID       uint      `gorm:"not null"`
	Usuario         Usuario   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CategoriaID     uint      `gorm:"not null"`
	Categoria       Categoria `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

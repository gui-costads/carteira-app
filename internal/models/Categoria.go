package models

import "gorm.io/gorm"

type Categoria struct {
	gorm.Model
	Nome          string      `gorm:"not null; unique"`
	TipoDeReceita string      `gorm:"not null;check:tipo_de_receita IN ('renda', 'despesa')"`
	Transacoes    []Transacao `gorm:"foreignKey:CategoriaID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	Orcamentos    []Orcamento `gorm:"foreignKey:CategoriaID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
}

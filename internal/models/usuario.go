package models

import "gorm.io/gorm"

type Usuario struct {
	gorm.Model
	Nome       string      `gorm:"not null"`
	Sobrenome  string      `gorm:"not null"`
	Email      string      `gorm:"not null;unique"`
	Senha      string      `gorm:"type:varchar(255);not null"`
	Orcamentos []Orcamento `gorm:"foreignKey:UsuarioID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Transacoes []Transacao `gorm:"foreignKey:UsuarioID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

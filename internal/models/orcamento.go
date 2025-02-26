package models

import "gorm.io/gorm"

type Orcamento struct {
	gorm.Model
	Saldo       float64   `gorm:"type:decimal(10,2);not null; check:saldo >= 0"`
	Periodo     string    `gorm:"not null; check:periodo IN ('semanal','mensal', 'anual'); uniqueIndex:idx_orcamento_usuario_periodo_categoria"`
	UsuarioID   uint      `gorm:"not null; uniqueIndex:idx_orcamento_usuario_periodo_categoria"`
	Usuario     Usuario   `gorm:"foreignKey:UsuarioID"`
	CategoriaID uint      `gorm:"not null; uniqueIndex:idx_orcamento_usuario_periodo_categoria"`
	Categoria   Categoria `gorm:"foreignKey:CategoriaID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
}

package models

type Orcamento struct {
	Saldo       float64   `gorm:"type:decimal(10,2);not null; check:saldo >= 0"`
	Periodo     string    `gorm:"not null; check:periodo IN ('semanal','mensal', 'anual'); uniqueIndex:idx_orcamento_usuario_periodo_categoria"`
	UsuarioID   uint      `gorm:"not null; uniqueIndex:idx_orcamento_usuario_periodo_categoria"`
	Usuario     Usuario   `gorm:"constraint:OnUpdate:CASCADE;"`
	CategoriaID uint      `gorm:"not null; uniqueIndex:idx_orcamento_usuario_periodo_categoria"`
	Categoria   Categoria `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
}

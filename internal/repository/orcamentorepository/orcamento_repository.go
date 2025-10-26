package orcamentorepository

import (
	"github.com/gui-costads/carteira-app/internal/models"
)

type OrcamentoRepository interface {
	Criar(orcamento models.Orcamento) (models.Orcamento, error)
	Atualizar(orcamento models.Orcamento) (models.Orcamento, error)
	Deletar(orcamento models.Orcamento) error
	BuscarPorID(id uint) (models.Orcamento, error)
	BuscarTodos() ([]models.Orcamento, error)
	BuscarPorUsuarioId(id uint) ([]models.Orcamento, error)
}

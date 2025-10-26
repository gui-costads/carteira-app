package orcamentoservice

import "github.com/gui-costads/carteira-app/internal/data/orcamentodto"

type OrcamentoService interface {
	CriarOrcamento(orcament orcamentodto.CriarOrcamentoRequest) (orcamentodto.ResponseOrcamento, error)
	AtualizarOrcamento(id uint, orcamentodto orcamentodto.AtualizarOrcamentoRequest) (orcamentodto.ResponseOrcamento, error)
	DeletarOrcamento(id uint) error
	BuscarOrcamentoPorID(id uint) (orcamentodto.ResponseOrcamento, error)
	BuscarTodosOrcamentos() ([]orcamentodto.ResponseOrcamento, error)
	BuscarOrcamentoPorUsuarioId(id uint) ([]orcamentodto.ResponseOrcamento, error)
}

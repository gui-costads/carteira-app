package orcamentoservice

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/gui-costads/carteira-app/internal/data/orcamentodto"
	"github.com/gui-costads/carteira-app/internal/models"
	"github.com/gui-costads/carteira-app/internal/repository/orcamentorepository"
	"gorm.io/gorm"
)

type orcamentoServiceImpl struct {
	orcamentoRepo orcamentorepository.OrcamentoRepository
	validate      *validator.Validate
}

func NewOrcamentoService(orcamentoRepo orcamentorepository.OrcamentoRepository) OrcamentoService {
	return &orcamentoServiceImpl{
		orcamentoRepo: orcamentoRepo,
		validate:      validator.New(),
	}
}

func (orcamentoService *orcamentoServiceImpl) CriarOrcamento(request orcamentodto.CriarOrcamentoRequest) (orcamentodto.ResponseOrcamento, error) {
	if err := orcamentoService.validate.Struct(request); err != nil {
		return orcamentodto.ResponseOrcamento{}, err
	}

	orcamento := models.Orcamento{
		Saldo:       request.Saldo,
		Periodo:     request.Periodo,
		UsuarioID:   request.UsuarioID,
		CategoriaID: request.CategoriaID,
	}

	orcamentoModel, err := orcamentoService.orcamentoRepo.Criar(orcamento)
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return orcamentodto.ResponseOrcamento{}, err
		}
		return orcamentodto.ResponseOrcamento{}, err
	}

	orcamentoResponse := orcamentodto.ResponseOrcamento{
		ID:          orcamentoModel.ID,
		Saldo:       orcamentoModel.Saldo,
		Periodo:     orcamentoModel.Periodo,
		UsuarioID:   orcamentoModel.UsuarioID,
		CategoriaID: orcamentoModel.CategoriaID,
	}

	return orcamentoResponse, nil
}

func (orcamentoService *orcamentoServiceImpl) AtualizarOrcamento(id uint, request orcamentodto.AtualizarOrcamentoRequest) (orcamentodto.ResponseOrcamento, error) {
	orcamento, err := orcamentoService.orcamentoRepo.BuscarPorID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return orcamentodto.ResponseOrcamento{}, errors.New("orçamento não encontrado")
		}
		return orcamentodto.ResponseOrcamento{}, err
	}

	if err := orcamentoService.validate.Struct(request); err != nil {
		return orcamentodto.ResponseOrcamento{}, err
	}

	if request.Saldo != nil {
		orcamento.Saldo = *request.Saldo
	}
	if request.Periodo != nil {
		orcamento.Periodo = *request.Periodo
	}
	if request.CategoriaID != nil {
		orcamento.CategoriaID = *request.CategoriaID
	}

	updated, err := orcamentoService.orcamentoRepo.Atualizar(orcamento)
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return orcamentodto.ResponseOrcamento{}, err
		}
		return orcamentodto.ResponseOrcamento{}, err
	}

	orcamentoResponse := orcamentodto.ResponseOrcamento{
		ID:          updated.ID,
		Saldo:       updated.Saldo,
		Periodo:     updated.Periodo,
		UsuarioID:   updated.UsuarioID,
		CategoriaID: updated.CategoriaID,
	}

	return orcamentoResponse, nil
}

func (orcamentoService *orcamentoServiceImpl) DeletarOrcamento(id uint) error {
	orcamento, err := orcamentoService.orcamentoRepo.BuscarPorID(id)
	if err != nil {
		return err
	}

	orcamentoService.orcamentoRepo.Deletar(orcamento)
	return nil
}

func (orcamentoService *orcamentoServiceImpl) BuscarOrcamentoPorID(id uint) (orcamentodto.ResponseOrcamento, error) {
	orcamento, err := orcamentoService.orcamentoRepo.BuscarPorID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return orcamentodto.ResponseOrcamento{}, errors.New("orçamento não encontrado")
		}
		return orcamentodto.ResponseOrcamento{}, err
	}

	return orcamentodto.ResponseOrcamento{
		ID:          orcamento.ID,
		Saldo:       orcamento.Saldo,
		Periodo:     orcamento.Periodo,
		UsuarioID:   orcamento.UsuarioID,
		CategoriaID: orcamento.CategoriaID,
	}, nil
}

func (orcamentoService *orcamentoServiceImpl) BuscarTodosOrcamentos() ([]orcamentodto.ResponseOrcamento, error) {
	orcamentos, err := orcamentoService.orcamentoRepo.BuscarTodos()
	if err != nil {
		return nil, err
	}

	var response []orcamentodto.ResponseOrcamento
	for _, o := range orcamentos {
		response = append(response, orcamentodto.ResponseOrcamento{
			ID:    o.ID,
			Saldo: o.Saldo, // Atualiza no repositório

			Periodo:     o.Periodo,
			UsuarioID:   o.UsuarioID,
			CategoriaID: o.CategoriaID,
		})
	}

	return response, nil
}

func (orcamentoService *orcamentoServiceImpl) BuscarOrcamentoPorUsuarioId(id uint) ([]orcamentodto.ResponseOrcamento, error) {
	orcamentos, err := orcamentoService.orcamentoRepo.BuscarPorUsuarioId(id)
	if err != nil {
		return nil, err
	}

	var response []orcamentodto.ResponseOrcamento
	for _, o := range orcamentos {
		response = append(response, orcamentodto.ResponseOrcamento{
			ID:    o.ID,
			Saldo: o.Saldo, // Atualiza no repositório

			Periodo:     o.Periodo,
			UsuarioID:   o.UsuarioID,
			CategoriaID: o.CategoriaID,
		})
	}

	return response, nil
}

package orcamentorepository

import (
	"errors"
	"fmt"

	"github.com/gui-costads/carteira-app/internal/models"
	"gorm.io/gorm"
)

type orcamentoRepositoryImpl struct {
	db *gorm.DB
}

func NewOrcamentoRepository(db *gorm.DB) OrcamentoRepository {
	return &orcamentoRepositoryImpl{db: db}
}

func (repo *orcamentoRepositoryImpl) Criar(orcamento models.Orcamento) (models.Orcamento, error) {
	if err := repo.db.Create(&orcamento).Error; err != nil {
		return models.Orcamento{}, err
	}
	return orcamento, nil
}

func (repo *orcamentoRepositoryImpl) Atualizar(orcamento models.Orcamento) (models.Orcamento, error) {
	if err := repo.db.Save(&orcamento).Error; err != nil {
		return models.Orcamento{}, err
	}
	return orcamento, nil
}

func (repo *orcamentoRepositoryImpl) Deletar(orcamento models.Orcamento) error {
	if err := repo.db.Delete(&orcamento).Error; err != nil {
		return err
	}
	return nil
}

func (repo *orcamentoRepositoryImpl) BuscarPorID(id uint) (models.Orcamento, error) {
	var orcamento models.Orcamento
	if err := repo.db.First(&orcamento, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return orcamento, fmt.Errorf("orcamento com ID %d não encontrado", id)
		}
		return orcamento, err
	}
	return orcamento, nil
}

func (repo *orcamentoRepositoryImpl) BuscarTodos() ([]models.Orcamento, error) {
	var orcamentos []models.Orcamento
	if err := repo.db.Find(&orcamentos).Error; err != nil {
		return nil, err
	}
	return orcamentos, nil
}

func (repo *orcamentoRepositoryImpl) BuscarPorUsuarioId(id uint) ([]models.Orcamento, error) {
	var orcamentos []models.Orcamento
	if err := repo.db.Where("usuario_id = ?", id).Find(&orcamentos).Error; err != nil {
		return nil, err
	}
	return orcamentos, nil
}

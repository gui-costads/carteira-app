package orcamentorepository

import (
	"errors"
	"fmt"

	"github.com/gui-costads/carteira-app/internal/models"
	"gorm.io/gorm"
)

type OrcamentoRepositoryImpl struct {
	Db *gorm.DB
}

func NewUsuarioRepositoryImpl(db *gorm.DB) *OrcamentoRepositoryImpl {
	return &OrcamentoRepositoryImpl{Db: db}
}

func (repo *OrcamentoRepositoryImpl) Criar(orcamento models.Orcamento) (models.Orcamento, error) {
	if err := repo.Db.Create(&orcamento).Error; err != nil {
		return models.Orcamento{}, err
	}
	return orcamento, nil
}

func (repo *OrcamentoRepositoryImpl) Atualizar(orcamento models.Orcamento) (models.Orcamento, error) {
	if err := repo.Db.Save(&orcamento).Error; err != nil {
		return models.Orcamento{}, err
	}
	return orcamento, nil
}

func (repo *OrcamentoRepositoryImpl) Deletar(orcamento models.Orcamento) error {
	if err := repo.Db.Delete(&orcamento).Error; err != nil {
		return err
	}
	return nil
}

func (repo *OrcamentoRepositoryImpl) BuscarPorID(id uint) (models.Orcamento, error) {
	var orcamento models.Orcamento
	if err := repo.Db.First(&orcamento, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return orcamento, fmt.Errorf("orcamento com ID %d não encontrado", id)
		}
		return orcamento, err
	}
	return orcamento, nil
}

func (repo *OrcamentoRepositoryImpl) BuscarTodos() ([]models.Orcamento, error) {
	var orcamentos []models.Orcamento
	if err := repo.Db.Find(&orcamentos).Error; err != nil {
		return nil, err
	}
	return orcamentos, nil
}

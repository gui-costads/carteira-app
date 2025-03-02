package transacaorepository

import (
	"errors"
	"fmt"

	"github.com/gui-costads/carteira-app/internal/models"
	"gorm.io/gorm"
)

type transacaoRepositoryImpl struct {
	db *gorm.DB
}

func NewTransacaoRepository(db *gorm.DB) TransacaoRepository {
	return &transacaoRepositoryImpl{
		db: db,
	}
}

func (repo *transacaoRepositoryImpl) Criar(transacao models.Transacao) (models.Transacao, error) {
	if err := repo.db.Create(&transacao).Error; err != nil {
		return models.Transacao{}, err
	}
	return transacao, nil
}

func (repo *transacaoRepositoryImpl) Atualizar(transacao models.Transacao) (models.Transacao, error) {
	if err := repo.db.Save(&transacao).Error; err != nil {
		return models.Transacao{}, err
	}
	return transacao, nil
}

func (repo *transacaoRepositoryImpl) Deletar(transacao models.Transacao) error {
	if err := repo.db.Delete(&transacao).Error; err != nil {
		return err
	}
	return nil
}

func (repo *transacaoRepositoryImpl) BuscarPorID(id uint) (models.Transacao, error) {
	var transacao models.Transacao
	if err := repo.db.First(&transacao, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return transacao, fmt.Errorf("transacao com ID %d não encontrado", id)
		}
		return transacao, err
	}
	return transacao, nil
}

func (repo *transacaoRepositoryImpl) BuscarTodos() ([]models.Transacao, error) {
	var transacaos []models.Transacao
	if err := repo.db.Find(&transacaos).Error; err != nil {
		return nil, err
	}
	return transacaos, nil
}

package categoriarepository

import (
	"fmt"

	"github.com/gui-costads/carteira-app/internal/models"
	"gorm.io/gorm"
)

type CategoriaRepositoryImpl struct {
	db *gorm.DB
}

func NewCategoriaRepositoryImpl(db *gorm.DB) CategoriaRepository {
	return &CategoriaRepositoryImpl{
		db: db,
	}
}

func (categoriaRepository *CategoriaRepositoryImpl) Criar(categoria models.Categoria) (models.Categoria, error) {
	if err := categoriaRepository.db.Create(&categoria).Error; err != nil {
		return models.Categoria{}, err
	}
	return categoria, nil
}

func (categoriaRepository *CategoriaRepositoryImpl) Atualizar(categoria models.Categoria) (models.Categoria, error) {
	if err := categoriaRepository.db.Save(&categoria).Error; err != nil {
		return models.Categoria{}, err
	}
	return categoria, nil
}

func (categoriaRepository *CategoriaRepositoryImpl) Deletar(categoria models.Categoria) error {
	if err := categoriaRepository.db.Delete(&categoria).Error; err != nil {
		return err
	}
	return nil
}

func (categoriaRepository *CategoriaRepositoryImpl) BuscarTodos() ([]models.Categoria, error) {
	var categorias []models.Categoria
	if err := categoriaRepository.db.Find(&categorias).Error; err != nil {
		return nil, err
	}
	return categorias, nil
}

func (categoriaRepository *CategoriaRepositoryImpl) BuscarPorID(id uint) (models.Categoria, error) {
	var categoria models.Categoria
	if err := categoriaRepository.db.First(&categoria, id).Error; err != nil {
		return categoria, fmt.Errorf("categoria com ID %d não encontrado", id)
	}
	return categoria, nil
}

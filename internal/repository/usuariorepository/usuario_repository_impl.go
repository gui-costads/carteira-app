package usuariorepository

import (
	"errors"

	"fmt"

	"github.com/gui-costads/carteira-app/internal/models"
	"gorm.io/gorm"
)

type usuarioRepositoryImpl struct {
	db *gorm.DB
}

func NewUsuarioRepository(db *gorm.DB) UsuarioRepository {
	return &usuarioRepositoryImpl{
		db: db,
	}
}

func (repo *usuarioRepositoryImpl) Criar(usuario models.Usuario) (models.Usuario, error) {
	if err := repo.db.Create(&usuario).Error; err != nil {
		return models.Usuario{}, err
	}
	return usuario, nil
}

func (repo *usuarioRepositoryImpl) Atualizar(usuario models.Usuario) (models.Usuario, error) {
	if err := repo.db.Save(&usuario).Error; err != nil {
		return models.Usuario{}, err
	}
	return usuario, nil
}

func (repo *usuarioRepositoryImpl) Deletar(usuario models.Usuario) error {
	if err := repo.db.Delete(&usuario).Error; err != nil {
		return err
	}
	return nil
}

func (repo *usuarioRepositoryImpl) BuscarPorID(id uint) (models.Usuario, error) {
	var usuario models.Usuario
	if err := repo.db.First(&usuario, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return usuario, fmt.Errorf("usuário com ID %d não encontrado", id)
		}
		return usuario, err
	}
	return usuario, nil
}

func (repo *usuarioRepositoryImpl) BuscarTodos() ([]models.Usuario, error) {
	var usuarios []models.Usuario
	if err := repo.db.Find(&usuarios).Error; err != nil {
		return nil, err
	}
	return usuarios, nil
}

func (repo *usuarioRepositoryImpl) BuscarPorEmail(email string) (models.Usuario, error) {
	var usuario models.Usuario
	if err := repo.db.Where("email = ?", email).First(&usuario).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return usuario, fmt.Errorf("usuário com email %s não encontrado", email)
		}
		return usuario, err
	}
	return usuario, nil
}

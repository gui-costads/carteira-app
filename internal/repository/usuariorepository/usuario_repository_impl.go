package usuariorepository

import (
	"errors"

	"github.com/gui-costads/carteira-app/internal/models"
	"gorm.io/gorm"
)

type UsuarioRepositoryImpl struct {
	Db *gorm.DB
}

func NewUsuarioRepositoryImpl(db *gorm.DB) *UsuarioRepositoryImpl {
	return &UsuarioRepositoryImpl{
		Db: db,
	}
}

func (repo *UsuarioRepositoryImpl) Criar(usuario models.Usuario) (models.Usuario, error) {
	if err := repo.Db.Create(&usuario).Error; err != nil {
		return models.Usuario{}, err
	}
	return usuario, nil
}

func (repo *UsuarioRepositoryImpl) Atualizar(usuario models.Usuario) (models.Usuario, error) {
	if err := repo.Db.Save(&usuario).Error; err != nil {
		return models.Usuario{}, err
	}
	return usuario, nil
}

func (repo *UsuarioRepositoryImpl) Deletar(usuario models.Usuario) error {
	if err := repo.Db.Delete(&usuario).Error; err != nil {
		return err
	}
	return nil
}

func (repo *UsuarioRepositoryImpl) BuscarPorID(id uint) (models.Usuario, error) {
	var usuario models.Usuario
	if err := repo.Db.First(&usuario, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return usuario, nil
		}
		return usuario, err
	}
	return usuario, nil
}

func (repo *UsuarioRepositoryImpl) BuscarTodos() ([]models.Usuario, error) {
	var usuarios []models.Usuario
	if err := repo.Db.Find(&usuarios).Error; err != nil {
		return nil, err
	}
	return usuarios, nil
}

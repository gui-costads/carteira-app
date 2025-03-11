package usuariorepository

import (
	"github.com/gui-costads/carteira-app/internal/models"
)

type UsuarioRepository interface {
	Criar(usuario models.Usuario) (models.Usuario, error)
	Atualizar(usuario models.Usuario) (models.Usuario, error)
	Deletar(usuario models.Usuario) error
	BuscarPorID(id uint) (usuario models.Usuario, err error)
	BuscarTodos() ([]models.Usuario, error)
	BuscarPorEmail(email string) (usuario models.Usuario, err error)
}

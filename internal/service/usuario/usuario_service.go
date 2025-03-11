package usuarioservice

import (
	"github.com/gui-costads/carteira-app/internal/data/usuariodto"
)

type UsuarioService interface {
	CriarUsuario(usuario usuariodto.CriarUsuarioRequest) (usuariodto.UsuarioResponse, error)
	AtualizarUsuario(id uint, usuario usuariodto.AtualizarUsuarioRequest) (usuariodto.UsuarioResponse, error)
	DeletarUsuario(id uint) error
	BuscarUsuarioPorID(id uint) (usuariodto.UsuarioResponse, error)
	BuscarTodosUsuarios() ([]usuariodto.UsuarioResponse, error)
	AutenticarUsuario(usuario usuariodto.LoginRequest) (usuariodto.UsuarioResponse, error)
	BuscarUsuarioPorEmail(email string) (usuariodto.UsuarioResponse, error)
}

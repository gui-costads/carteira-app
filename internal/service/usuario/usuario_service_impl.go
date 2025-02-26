package usuarioservice

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/gui-costads/carteira-app/internal/data/usuariodto"
	"github.com/gui-costads/carteira-app/internal/models"
	"github.com/gui-costads/carteira-app/internal/repository/usuariorepository"
	"gorm.io/gorm"
)

type UsuarioServiceImpl struct {
	usuarioRepo usuariorepository.UsuarioRepository
	validate    *validator.Validate
}

func NewUsuarioService(usuarioRepo usuariorepository.UsuarioRepository) *UsuarioServiceImpl {
	return &UsuarioServiceImpl{
		usuarioRepo: usuarioRepo,
		validate:    validator.New(),
	}
}

func (usarioService *UsuarioServiceImpl) CriarUsuario(request usuariodto.CriarUsuarioRequest) (usuariodto.UsuarioResponse, error) {
	if err := usarioService.validate.Struct(request); err != nil {
		return usuariodto.UsuarioResponse{}, err
	}

	usuarioModel := models.Usuario{
		Nome:      request.Nome,
		Sobrenome: request.Sobrenome,
		Email:     request.Email,
		Senha:     request.Senha,
	}

	usuarioModel, err := usarioService.usuarioRepo.Criar(usuarioModel)
	if err != nil {
		return usuariodto.UsuarioResponse{}, err
	}

	usuarioResponse := usuariodto.UsuarioResponse{
		ID:        usuarioModel.ID,
		Nome:      usuarioModel.Nome,
		Sobrenome: usuarioModel.Sobrenome,
		Email:     usuarioModel.Email,
	}
	return usuarioResponse, nil
}

func (usarioService *UsuarioServiceImpl) AtualizarUsuario(id uint, usuarioRequest usuariodto.AtualizarUsuarioRequest) (usuariodto.UsuarioResponse, error) {
	usuario, err := usarioService.usuarioRepo.BuscarPorID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return usuariodto.UsuarioResponse{}, errors.New("usuario not found")
		}
		return usuariodto.UsuarioResponse{}, err
	}

	if err := usarioService.validate.Struct(usuarioRequest); err != nil {
		return usuariodto.UsuarioResponse{}, err
	}

	if usuarioRequest.Nome != "" {
		usuario.Nome = usuarioRequest.Nome
	}
	if usuarioRequest.Sobrenome != "" {
		usuario.Sobrenome = usuarioRequest.Sobrenome
	}
	if usuarioRequest.Senha != "" {
		usuario.Senha = usuarioRequest.Senha // You might want to hash the password here
	}

	usuario, err = usarioService.usuarioRepo.Atualizar(usuario)
	if err != nil {
		return usuariodto.UsuarioResponse{}, err
	}

	usuarioResponse := usuariodto.UsuarioResponse{
		ID:        usuario.ID,
		Nome:      usuario.Nome,
		Sobrenome: usuario.Sobrenome,
		Email:     usuario.Email,
	}
	return usuarioResponse, nil
}

func (usarioService *UsuarioServiceImpl) DeletarUsuario(id uint) error {
	usuario, err := usarioService.usuarioRepo.BuscarPorID(id)
	if err != nil {
		return err
	}

	usarioService.usuarioRepo.Deletar(usuario)
	return nil
}

func (usarioService *UsuarioServiceImpl) BuscarUsuarioPorID(id uint) (usuariodto.UsuarioResponse, error) {
	usuario, err := usarioService.usuarioRepo.BuscarPorID(id)
	if err != nil {
		return usuariodto.UsuarioResponse{}, err
	}

	usuarioResponse := usuariodto.UsuarioResponse{
		ID:        usuario.ID,
		Nome:      usuario.Nome,
		Sobrenome: usuario.Sobrenome,
		Email:     usuario.Email,
	}

	return usuarioResponse, nil
}

func (usarioService *UsuarioServiceImpl) BuscarTodosUsuarios() ([]usuariodto.UsuarioResponse, error) {
	usuarios := usarioService.usuarioRepo.BuscarTodos()

	var usuariosResponse []usuariodto.UsuarioResponse
	for _, usuario := range usuarios {
		usuarioResponse := usuariodto.UsuarioResponse{
			ID:        usuario.ID,
			Nome:      usuario.Nome,
			Sobrenome: usuario.Sobrenome,
			Email:     usuario.Email,
		}
		usuariosResponse = append(usuariosResponse, usuarioResponse)
	}

	return usuariosResponse, nil
}

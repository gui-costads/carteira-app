package usuarioservice

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/gui-costads/carteira-app/internal/data/usuariodto"
	"github.com/gui-costads/carteira-app/internal/models"
	"github.com/gui-costads/carteira-app/internal/repository/usuariorepository"
	"golang.org/x/crypto/bcrypt"
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

func (usuarioService *UsuarioServiceImpl) CriarUsuario(request usuariodto.CriarUsuarioRequest) (usuariodto.UsuarioResponse, error) {
	if err := usuarioService.validate.Struct(request); err != nil {
		return usuariodto.UsuarioResponse{}, err
	}

	hashSenha, _ := encriptarSenhar(request.Senha)

	usuarioModel := models.Usuario{
		Nome:      request.Nome,
		Sobrenome: request.Sobrenome,
		Email:     request.Email,
		Senha:     hashSenha,
	}

	usuarioModel, err := usuarioService.usuarioRepo.Criar(usuarioModel)
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

func (usuarioService *UsuarioServiceImpl) AtualizarUsuario(id uint, usuarioRequest usuariodto.AtualizarUsuarioRequest) (usuariodto.UsuarioResponse, error) {
	usuario, err := usuarioService.usuarioRepo.BuscarPorID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return usuariodto.UsuarioResponse{}, errors.New("usuario not found")
		}
		return usuariodto.UsuarioResponse{}, err
	}

	if err := usuarioService.validate.Struct(usuarioRequest); err != nil {
		return usuariodto.UsuarioResponse{}, err
	}

	if usuarioRequest.Nome != "" {
		usuario.Nome = usuarioRequest.Nome
	}
	if usuarioRequest.Sobrenome != "" {
		usuario.Sobrenome = usuarioRequest.Sobrenome
	}
	if usuarioRequest.Senha != "" {
		usuario.Senha = usuarioRequest.Senha
	}

	usuario, err = usuarioService.usuarioRepo.Atualizar(usuario)
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

func (usuarioService *UsuarioServiceImpl) DeletarUsuario(id uint) error {
	usuario, err := usuarioService.usuarioRepo.BuscarPorID(id)
	if err != nil {
		return err
	}

	usuarioService.usuarioRepo.Deletar(usuario)
	return nil
}

func (usuarioService *UsuarioServiceImpl) BuscarUsuarioPorID(id uint) (usuariodto.UsuarioResponse, error) {
	usuario, err := usuarioService.usuarioRepo.BuscarPorID(id)
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

func encriptarSenhar(senha string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(senha), 14)
	return string(bytes), err
}

func verificarSenha(senha, hashSenha string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashSenha), []byte(senha))
	return err == nil
}

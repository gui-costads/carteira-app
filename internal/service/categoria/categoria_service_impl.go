package categoriaservice

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/gui-costads/carteira-app/internal/data/categoriadto"
	"github.com/gui-costads/carteira-app/internal/models"
	"github.com/gui-costads/carteira-app/internal/repository/categoriarepository"
	"gorm.io/gorm"
)

type categoriaServiceImpl struct {
	categoriaRepository categoriarepository.CategoriaRepository
	validate            *validator.Validate
}

func NewCategoriaService(categoriaRepository categoriarepository.CategoriaRepository) CategoriaService {
	return &categoriaServiceImpl{
		categoriaRepository: categoriaRepository,
		validate:            validator.New(),
	}
}

func (categoriaService *categoriaServiceImpl) CriarCategoria(request categoriadto.CriarCategoriaRequest) (categoriadto.ResponseCategoria, error) {
	if err := categoriaService.validate.Struct(request); err != nil {
		return categoriadto.ResponseCategoria{}, err
	}

	categoria := models.Categoria{
		Nome:          request.Nome,
		TipoDeReceita: request.TipoDeReceita,
	}

	categoriaModel, err := categoriaService.categoriaRepository.Criar(categoria)

	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return categoriadto.ResponseCategoria{}, err
		}
		return categoriadto.ResponseCategoria{}, err
	}

	response := categoriadto.ResponseCategoria{
		ID:            categoriaModel.ID,
		Nome:          categoriaModel.Nome,
		TipoDeReceita: categoriaModel.TipoDeReceita,
	}

	return response, nil
}

func (categoriaService *categoriaServiceImpl) AtualizarCategoria(id uint, request categoriadto.AtualizarCategoriaRequest) (categoriadto.ResponseCategoria, error) {
	categoria, err := categoriaService.categoriaRepository.BuscarPorID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return categoriadto.ResponseCategoria{}, err
		}
		return categoriadto.ResponseCategoria{}, err
	}

	if err := categoriaService.validate.Struct(request); err != nil {
		return categoriadto.ResponseCategoria{}, err
	}

	if request.Nome != nil {
		categoria.Nome = *request.Nome
	}
	if request.TipoDeReceita != nil {
		categoria.TipoDeReceita = *request.TipoDeReceita
	}

	updated, err := categoriaService.categoriaRepository.Atualizar(categoria)
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return categoriadto.ResponseCategoria{}, err
		}
		return categoriadto.ResponseCategoria{}, err
	}

	response := categoriadto.ResponseCategoria{
		ID:            updated.ID,
		Nome:          updated.Nome,
		TipoDeReceita: updated.TipoDeReceita,
	}

	return response, nil
}

func (categoriaService *categoriaServiceImpl) DeletarCategoria(id uint) error {
	categoria, err := categoriaService.categoriaRepository.BuscarPorID(id)
	if err != nil {
		return err
	}

	categoriaService.categoriaRepository.Deletar(categoria)
	return nil
}

func (categoriaService *categoriaServiceImpl) BuscarCategoriaPorID(id uint) (categoriadto.ResponseCategoria, error) {
	categoria, err := categoriaService.categoriaRepository.BuscarPorID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return categoriadto.ResponseCategoria{}, err
		}
		return categoriadto.ResponseCategoria{}, err
	}

	response := categoriadto.ResponseCategoria{
		ID:            categoria.ID,
		Nome:          categoria.Nome,
		TipoDeReceita: categoria.TipoDeReceita,
	}

	return response, nil
}

func (categoriaService *categoriaServiceImpl) BuscarTodasCategorias() ([]categoriadto.ResponseCategoria, error) {
	categorias, err := categoriaService.categoriaRepository.BuscarTodos()
	if err != nil {
		return nil, err
	}

	var response []categoriadto.ResponseCategoria
	for _, c := range categorias {
		response = append(response, categoriadto.ResponseCategoria{
			ID:            c.ID,
			Nome:          c.Nome,
			TipoDeReceita: c.TipoDeReceita,
		})
	}

	return response, nil
}

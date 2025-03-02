package categoriaservice

import (
	"github.com/gui-costads/carteira-app/internal/data/categoriadto"
)

type CategoriaService interface {
	CriarCategoria(categoria categoriadto.CriarCategoriaRequest) (categoriadto.ResponseCategoria, error)
	AtualizarCategoria(id uint, categoria categoriadto.AtualizarCategoriaRequest) (categoriadto.ResponseCategoria, error)
	DeletarCategoria(id uint) error
	BuscarCategoriaPorID(id uint) (categoriadto.ResponseCategoria, error)
	BuscarTodasCategorias() ([]categoriadto.ResponseCategoria, error)
}

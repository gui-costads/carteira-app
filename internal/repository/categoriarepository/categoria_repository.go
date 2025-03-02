package categoriarepository

import "github.com/gui-costads/carteira-app/internal/models"

type CategoriaRepository interface {
	Criar(categoria models.Categoria) (models.Categoria, error)
	Atualizar(categoria models.Categoria) (models.Categoria, error)
	Deletar(categoria models.Categoria) error
	BuscarTodos() ([]models.Categoria, error)
	BuscarPorID(id uint) (models.Categoria, error)
}

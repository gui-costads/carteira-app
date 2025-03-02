package transacaorepository

import "github.com/gui-costads/carteira-app/internal/models"

type TransacaoRepository interface {
	Criar(transacao models.Transacao) (models.Transacao, error)
	Atualizar(transacao models.Transacao) (models.Transacao, error)
	Deletar(transacao models.Transacao) error
	BuscarPorID(id uint) (models.Transacao, error)
	BuscarTodos() ([]models.Transacao, error)
}

package transacaoservice

import (
	"github.com/gui-costads/carteira-app/internal/data/transacaodto"
)

type TransacaoService interface {
	CriarTransacao(transacao transacaodto.CriarTransacaoRequest) (transacaodto.ResponseTransacao, error)
	AtualizarTransacao(id uint, transacao transacaodto.AtualizarTransacaoRequest) (transacaodto.ResponseTransacao, error)
	DeletarTransacao(id uint) error
	BuscarTransacaoPorID(id uint) (transacaodto.ResponseTransacao, error)
	BuscarTodasTransacoes() ([]transacaodto.ResponseTransacao, error)
}

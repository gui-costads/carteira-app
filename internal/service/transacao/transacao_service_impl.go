package transacaoservice

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/gui-costads/carteira-app/internal/data/transacaodto"
	"github.com/gui-costads/carteira-app/internal/models"
	"github.com/gui-costads/carteira-app/internal/repository/transacaorepository"
	"gorm.io/gorm"
)

type transacaoServiceImpl struct {
	transacaoRepo transacaorepository.TransacaoRepository
	validate      *validator.Validate
}

func NewTransacaoService(transacaoRepo transacaorepository.TransacaoRepository) TransacaoService {
	return &transacaoServiceImpl{
		transacaoRepo: transacaoRepo,
		validate:      validator.New(),
	}
}

func (transacaoService *transacaoServiceImpl) CriarTransacao(request transacaodto.CriarTransacaoRequest) (transacaodto.ResponseTransacao, error) {
	if err := transacaoService.validate.Struct(request); err != nil {
		return transacaodto.ResponseTransacao{}, err
	}

	transacao := models.Transacao{
		Descricao:       request.Descricao,
		Valor:           request.Valor,
		Data:            request.Data,
		TipoDeTransacao: request.TipoDeTransacao,
		UsuarioID:       request.UsuarioID,
		CategoriaID:     request.CategoriaID,
	}

	transacaoModel, err := transacaoService.transacaoRepo.Criar(transacao)
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return transacaodto.ResponseTransacao{}, err
		}
		return transacaodto.ResponseTransacao{}, err
	}

	response := transacaodto.ResponseTransacao{
		ID:              transacaoModel.ID,
		Descricao:       transacaoModel.Descricao,
		Valor:           transacaoModel.Valor,
		Data:            transacaoModel.Data.Format("2006-01-02"),
		TipoDeTransacao: transacaoModel.TipoDeTransacao,
		UsuarioID:       transacaoModel.UsuarioID,
		CategoriaID:     transacaoModel.CategoriaID,
	}

	return response, nil
}

func (transacaoService *transacaoServiceImpl) AtualizarTransacao(id uint, request transacaodto.AtualizarTransacaoRequest) (transacaodto.ResponseTransacao, error) {
	transacao, err := transacaoService.transacaoRepo.BuscarPorID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return transacaodto.ResponseTransacao{}, errors.New("transacao não encontrada")
		}
		return transacaodto.ResponseTransacao{}, err
	}

	if err := transacaoService.validate.Struct(request); err != nil {
		return transacaodto.ResponseTransacao{}, err
	}

	if request.Descricao != nil {
		transacao.Descricao = *request.Descricao
	}
	if request.Valor != nil {
		transacao.Valor = *request.Valor
	}
	if request.Data != nil {
		transacao.Data = *request.Data
	}
	if request.TipoDeTransacao != nil {
		transacao.TipoDeTransacao = *request.TipoDeTransacao
	}
	if request.UsuarioID != nil {
		transacao.UsuarioID = *request.UsuarioID
	}
	if request.CategoriaID != nil {
		transacao.CategoriaID = *request.CategoriaID
	}

	updated, err := transacaoService.transacaoRepo.Atualizar(transacao)
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return transacaodto.ResponseTransacao{}, err
		}
		return transacaodto.ResponseTransacao{}, err
	}

	response := transacaodto.ResponseTransacao{
		ID:              updated.ID,
		Descricao:       updated.Descricao,
		Valor:           updated.Valor,
		Data:            updated.Data.Format("2006-01-02"),
		TipoDeTransacao: updated.TipoDeTransacao,
		UsuarioID:       updated.UsuarioID,
		CategoriaID:     updated.CategoriaID,
	}

	return response, nil
}

func (transacaoService *transacaoServiceImpl) DeletarTransacao(id uint) error {
	transacao, err := transacaoService.transacaoRepo.BuscarPorID(id)
	if err != nil {
		return err
	}

	transacaoService.transacaoRepo.Deletar(transacao)
	return nil
}

func (transacaoService *transacaoServiceImpl) BuscarTransacaoPorID(id uint) (transacaodto.ResponseTransacao, error) {
	transacao, err := transacaoService.transacaoRepo.BuscarPorID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return transacaodto.ResponseTransacao{}, errors.New("transacao não encontrada")
		}
		return transacaodto.ResponseTransacao{}, err
	}

	response := transacaodto.ResponseTransacao{
		ID:              transacao.ID,
		Descricao:       transacao.Descricao,
		Valor:           transacao.Valor,
		Data:            transacao.Data.Format("2006-01-02"),
		TipoDeTransacao: transacao.TipoDeTransacao,
		UsuarioID:       transacao.UsuarioID,
		CategoriaID:     transacao.CategoriaID,
	}

	return response, nil
}

func (transacaoService *transacaoServiceImpl) BuscarTodasTransacoes() ([]transacaodto.ResponseTransacao, error) {
	transacoes, err := transacaoService.transacaoRepo.BuscarTodos()
	if err != nil {
		return nil, err
	}

	var response []transacaodto.ResponseTransacao
	for _, t := range transacoes {
		response = append(response, transacaodto.ResponseTransacao{
			ID:              t.ID,
			Descricao:       t.Descricao,
			Valor:           t.Valor,
			Data:            t.Data.Format("2006-01-02"),
			TipoDeTransacao: t.TipoDeTransacao,
			UsuarioID:       t.UsuarioID,
			CategoriaID:     t.CategoriaID,
		})
	}

	return response, nil
}

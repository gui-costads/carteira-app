package usuarioservice

import (
	"testing"

	"github.com/gui-costads/carteira-app/internal/data/usuariodto"
)

func TestUsuarioServiceImpl_CriarUsuario(t *testing.T) {
	t.Run("Criar usuario com sucesso", func(t *testing.T) {
		usuarioService := NewUsuarioService(nil)
		request := usuariodto.CriarUsuarioRequest{
			Nome:      "Guilherme",
			Sobrenome: "Costa",
			Email:     "guilherme@email.com",
			Senha:     "123456",
		}

		response, err := usuarioService.CriarUsuario(request)
		if err != nil {
			t.Errorf("Erro ao criar usuario: %v", err)
		}

		if response.Nome != request.Nome {
			t.Errorf("Nome do usuario não corresponde")
		}
		if response.Sobrenome != request.Sobrenome {
			t.Errorf("Sobrenome do usuario não corresponde")
		}
		if response.Email != request.Email {
			t.Errorf("Email do usuario não corresponde")
		}
	})

}

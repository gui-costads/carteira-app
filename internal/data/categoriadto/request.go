package categoriadto

type CriarCategoriaRequest struct {
	Nome          string `json:"nome" binding:"required,min=1"`
	TipoDeReceita string `json:"tipoDeReceita" binding:"required,oneof=renda despesa"`
}

type AtualizarCategoriaRequest struct {
	Nome          *string `json:"nome,omitempty" binding:"omitempty,min=1"`
	TipoDeReceita *string `json:"tipoDeReceita,omitempty" binding:"omitempty,oneof=renda despesa"`
}

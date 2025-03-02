package categoriadto

type CriarCategoriaRequest struct {
	Nome          string `json:"nome" binding:"required,min=1"`
	TipoDeReceita string `json:"tipo_de_receita" binding:"required,oneof=renda despesa"`
}

type AtualizarCategoriaRequest struct {
	Nome          *string `json:"nome,omitempty" binding:"omitempty,min=1"`
	TipoDeReceita *string `json:"tipo_de_receita,omitempty" binding:"omitempty,oneof=renda despesa"`
}

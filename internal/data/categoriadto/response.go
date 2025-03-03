package categoriadto

type ResponseCategoria struct {
	ID            uint   `json:"id"`
	Nome          string `json:"nome"`
	TipoDeReceita string `json:"tipoDeReceita"`
}

package transacaodto

type ResponseTransacao struct {
	ID              uint    `json:"id"`
	Descricao       string  `json:"descricao"`
	Valor           float64 `json:"valor"`
	Data            string  `json:"data"`
	TipoDeTransacao string  `json:"tipo_de_transacao"`
	UsuarioID       uint    `json:"usuario_id"`
	CategoriaID     uint    `json:"categoria_id"`
}

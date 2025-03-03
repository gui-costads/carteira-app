package transacaodto

type ResponseTransacao struct {
	ID              uint    `json:"id"`
	Descricao       string  `json:"descricao"`
	Valor           float64 `json:"valor"`
	Data            string  `json:"data"`
	TipoDeTransacao string  `json:"tipoDeTransacao"`
	UsuarioID       uint    `json:"usuarioID"`
	CategoriaID     uint    `json:"categoriaID"`
}

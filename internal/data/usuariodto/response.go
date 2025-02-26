package usuariodto

type UsuarioResponse struct {
	ID        uint   `json:"id"`
	Nome      string `json:"nome"`
	Sobrenome string `json:"sobrenome"`
	Email     string `json:"email"`
}

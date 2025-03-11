package usuariodto

type CriarUsuarioRequest struct {
	Nome      string `json:"nome" validate:"required,min=3,max=100"`
	Sobrenome string `json:"sobrenome" validate:"required,min=3,max=100"`
	Email     string `json:"email" validate:"required,email"`
	Senha     string `json:"senha" validate:"required,min=6,max=100"`
}

type AtualizarUsuarioRequest struct {
	Nome      string `json:"nome" validate:"omitempty,min=3,max=100"`
	Sobrenome string `json:"sobrenome" validate:"omitempty,min=3,max=100"`
	Senha     string `json:"senha" validate:"omitempty,min=6,max=100"`
}

type LoginRequest struct {
	Email string `json:"email" validate:"required,email"`
	Senha string `json:"senha" validate:"required,min=6,max=100"`
}

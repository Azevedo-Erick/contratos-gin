package dtos

type CreateUsuario struct {
	Nome  string `json:"nome"  validate:"required"`
	Email string `json:"email"  validate:"required,email"`
	Senha string `json:"senha" validate:"required"`
}

type UpdateUsuario struct {
	Nome  string `json:"nome"`
	Email string `json:"email"`
	Senha string `json:"senha" validate:"required"`
}

type ResponseUsuario struct {
	Id    int    `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
	Senha string `json:"senha" validate:"required"`
}

package dtos

type CreateResponsavelEmpresa struct {
	Id       int    `json:"id" `
	Nome     string `json:"nome" validate:"required"`
	Email    string `json:"email"  validate:"required,email"`
	Cpf      string `json:"cpf" validate:"required"`
	Rg       string `json:"rg" validate:"required"`
	Telefone string `json:"telefone" validate:"required"`
}

type UpdateResponsavelEmpresa struct {
	Nome     string `json:"nome" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Cpf      string `json:"cpf" validate:"required"`
	Rg       string `json:"rg" validate:"required"`
	Telefone string `json:"telefone" validate:"required"`
}

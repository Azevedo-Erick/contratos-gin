package dtos

type CreateResponsavelEmpresa struct {
	Nome      string           `json:"nome" validate:"required"`
	Email     string           `json:"email"  validate:"required,email"`
	Cpf       string           `json:"cpf" validate:"required"`
	Rg        string           `json:"rg" validate:"required"`
	Telefones []CreateTelefone `json:"telefone" validate:"required"`
	Emails    []CreateEmail

	Endereco CreateEndereco `json:"endereco" validate:"required"`
}

type UpdateResponsavelEmpresa struct {
	Nome      string           `json:"nome" validate:"required"`
	Email     string           `json:"email" validate:"required,email"`
	Cpf       string           `json:"cpf" validate:"required"`
	Rg        string           `json:"rg" validate:"required"`
	Telefones []UpdateTelefone `json:"telefone" validate:"required"`
	Emails    []UpdateEmail

	Endereco UpdateEndereco `json:"endereco" validate:"required"`
}

type ResponseResponsavelEmpresa struct {
	Id        int                `json:"id" `
	Nome      string             `json:"nome" validate:"required"`
	Email     string             `json:"email"  validate:"required,email"`
	Cpf       string             `json:"cpf" validate:"required"`
	Rg        string             `json:"rg" validate:"required"`
	Telefones []ResponseTelefone `json:"telefone" validate:"required"`
	Emails    []ResponseEmail

	Endereco ResponseEndereco `json:"endereco" validate:"required"`
}

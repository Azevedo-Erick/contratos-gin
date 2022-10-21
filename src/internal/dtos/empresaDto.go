package dtos

type CreateEmpresa struct {
	Id                 int                      `json:"id" `
	Nome               string                   `json:"nome" validate:"required"`
	Cnpj               string                   `json:"cnpj" validate:"required"`
	ResponsavelEmpresa CreateResponsavelEmpresa `json:"empresa"`
}

type UpdateEmpresa struct {
	Id                 int                      `json:"id" `
	Nome               string                   `json:"nome" validate:"required"`
	Cnpj               string                   `json:"cnpj" validate:"required"`
	ResponsavelEmpresa CreateResponsavelEmpresa `json:"empresa"`
}

package dtos

type CreateEmpresa struct {
	RazaoSocial        string                   `json:"razao_social" validate:"required"`
	Cnpj               string                   `json:"cnpj" validate:"required"`
	ResponsavelEmpresa CreateResponsavelEmpresa `json:"empresa"`

	NomeFantasia       string `json:"nome_fantasia"`
	InscricaoMunicipal string `json:"inscricao_municipal"`
	InscricaoEstadual  string `json:"inscricao_estadual"`
	PorteEmpresa       string `json:"porte_empresa"`
	RamoAtividade      string `json:"ramo_atividade"`
	OptantePeloSimples bool   `json:"optante_pelo_simples"`

	Endereco  CreateEndereco   `json:"endereco"`
	Telefones []CreateTelefone `json:"telefones"`
	Emails    []CreateEmail    `json:"emails"`
}

type UpdateEmpresa struct {
	RazaoSocial        string                   `json:"razao_social" validate:"required"`
	Cnpj               string                   `json:"cnpj" validate:"required"`
	ResponsavelEmpresa CreateResponsavelEmpresa `json:"empresa"`

	NomeFantasia       string `json:"nome_fantasia"`
	InscricaoMunicipal string `json:"inscricao_municipal"`
	InscricaoEstadual  string `json:"inscricao_estadual"`
	PorteEmpresa       string `json:"porte_empresa"`
	RamoAtividade      string `json:"ramo_atividade"`
	OptantePeloSimples bool   `json:"optante_pelo_simples"`

	Endereco  UpdateEndereco   `json:"endereco"`
	Telefones []UpdateTelefone `json:"telefones"`
	Emails    []UpdateEmail    `json:"emails"`
}

type ResponseEmpresa struct {
	Id                 int                      `json:"id" `
	RazaoSocial        string                   `json:"razao_social" validate:"required"`
	Cnpj               string                   `json:"cnpj" validate:"required"`
	ResponsavelEmpresa CreateResponsavelEmpresa `json:"empresa"`

	NomeFantasia       string `json:"nome_fantasia"`
	InscricaoMunicipal string `json:"inscricao_municipal"`
	InscricaoEstadual  string `json:"inscricao_estadual"`
	PorteEmpresa       string `json:"porte_empresa"`
	RamoAtividade      string `json:"ramo_atividade"`
	OptantePeloSimples bool   `json:"optante_pelo_simples"`

	Endereco  CreateEndereco   `json:"endereco"`
	Telefones []CreateTelefone `json:"telefones"`
	Emails    []CreateEmail    `json:"emails"`
}

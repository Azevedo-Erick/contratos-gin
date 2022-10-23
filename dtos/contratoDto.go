package dtos

type CreateContrato struct {
	Id          int           `json:"id"`
	Contratante CreateEmpresa `json:"contratante"`
	Contratado  CreateEmpresa `json:"Contratado"`
	Valor       float64       `json:"valor" validate:"required"`
	DataInicio  string        `json:"dataInicio" validate:"required"`
	DataFim     string        `json:"dataFim" validate:"required"`
}

type UpdateContrato struct {
	Id          int           `json:"id"`
	Contratante CreateEmpresa `json:"contratante"`
	Contratado  CreateEmpresa `json:"Contratado"`
	Valor       float64       `json:"valor" validate:"required"`
	DataInicio  string        `json:"dataInicio" validate:"required"`
	DataFim     string        `json:"dataFim" validate:"required"`
}

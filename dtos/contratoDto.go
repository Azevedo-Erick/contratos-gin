package dtos

type CreateContrato struct {
	Contratante int     `json:"contratante"`
	Contratado  int     `json:"Contratado"`
	Valor       float64 `json:"valor" validate:"required"`
	DataInicio  string  `json:"dataInicio" validate:"required"`
	DataFim     string  `json:"dataFim" validate:"required"`
}

type UpdateContrato struct {
	Contratante int     `json:"contratante"`
	Contratado  int     `json:"Contratado"`
	Valor       float64 `json:"valor" validate:"required"`
	DataInicio  string  `json:"dataInicio" validate:"required"`
	DataFim     string  `json:"dataFim" validate:"required"`
}

type ResponseContrato struct {
	Id          int             `json:"id"`
	Contratante ResponseEmpresa `json:"contratante"`
	Contratado  ResponseEmpresa `json:"Contratado"`
	Valor       float64         `json:"valor" validate:"required"`
	DataInicio  string          `json:"dataInicio" validate:"required"`
	DataFim     string          `json:"dataFim" validate:"required"`
}

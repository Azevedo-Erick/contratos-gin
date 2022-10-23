package dtos

type LoginRequest struct {
	Email     string `json:"Email" form:"Email" binding:"required"`
	Senha     string `json:"Senha" form:"Senha" binding:"required"`
	RemeberMe string `json:"RemeberMe" form:"RemeberMe"`
}
type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

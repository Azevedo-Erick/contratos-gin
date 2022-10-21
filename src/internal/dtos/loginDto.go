package dtos

type LoginRequest struct {
	Username  string `json:"Username" form:"UserName" binding:"required"`
	Password  string `json:"Password" form:"Password" binding:"required"`
	RemeberMe string `json:"RemeberMe" form:"RemeberMe"`
}
type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

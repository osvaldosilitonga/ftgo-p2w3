package dto

type RegisterUserBody struct {
	Username      string `json:"username" validate:"required"`
	Password      string `json:"password" validate:"required"`
	DepositAmount int    `json:"deposit_amount" validate:"required"`
}

type LoginUserBody struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Code        int    `json:"code"`
	Message     string `json:"message"`
	AccessToken string `json:"access_token"`
}

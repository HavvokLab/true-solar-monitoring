package domain

type CreateSolarmanCredentialRequest struct {
	Username  string `json:"username" validate:"required"`
	Password  string `json:"password" validate:"required"`
	AppID     string `json:"app_id" validate:"required"`
	AppSecret string `json:"app_secret" validate:"required"`
	Owner     string `json:"owner" validate:"required"`
}

type UpdateSolarmanCredentialRequest struct {
	Username  string `json:"username" validate:"required"`
	Password  string `json:"password" validate:"required"`
	AppID     string `json:"app_id" validate:"required"`
	AppSecret string `json:"app_secret" validate:"required"`
	Owner     string `json:"owner" validate:"required"`
}

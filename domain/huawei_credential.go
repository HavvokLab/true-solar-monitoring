package domain

type CreateHuaweiCredentialRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Owner    string `json:"owner" validate:"required"`
}

type UpdateHuaweiCredentialRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Owner    string `json:"owner" validate:"required"`
}

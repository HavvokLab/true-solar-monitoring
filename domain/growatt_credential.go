package domain

type CreateGrowattCredentialRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Token    string `json:"token" validate:"required"`
	Owner    string `json:"owner" validate:"required"`
}

type UpdateGrowattCredentialRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Token    string `json:"token" validate:"required"`
	Owner    string `json:"owner" validate:"required"`
}

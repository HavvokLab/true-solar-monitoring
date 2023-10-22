package domain

type CreateKStarCredentialRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Owner    string `json:"owner" validate:"required"`
}

type UpdateKStarCredentialRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Owner    string `json:"owner" validate:"required"`
}

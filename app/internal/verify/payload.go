package verify

type CredentialsPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type SendRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type ConfirmPayload struct {
	Hash string `validate:"required"`
}

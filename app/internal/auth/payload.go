package auth

type LoginResponse struct {
	Token string `json:"token"`
}

type RegisterResponse struct {
	Token string `json:"token"`
}

type CredentialsPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginRequest struct {
	CredentialsPayload
}

type RegisterRequest struct {
	CredentialsPayload
	Name string `json:"name" validate:"required"`
}

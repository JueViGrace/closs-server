package types

type AuthResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type SignInRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RefreshRequest struct {
	Token string `json:"refresh_token"`
}

type RecoverPasswordResquest struct {
	Username string `json:"username"`
}

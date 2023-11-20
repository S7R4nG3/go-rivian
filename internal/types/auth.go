package types

type AuthResponse struct {
	Data AuthLoginResponse `json:"data"`
}

type AuthLoginResponse struct {
	Login AuthContent `json:"login"`
}

type AuthContent struct {
	OTPToken         string `json:"otpToken,omitempty"`
	AccessToken      string `json:"accessToken"`
	RefreshToken     string `json:"refreshToken"`
	UserSessionToken string `json:"userSessionToken"`
}

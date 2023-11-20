package types

type OTPResponse struct {
	Data OTPLoginWithOTP `json:"data"`
}

type OTPLoginWithOTP struct {
	LoginWithOTP OTPContent `json:"loginWithOTP"`
}

type OTPContent struct {
	AccessToken      string `json:"accessToken"`
	RefreshToken     string `json:"refreshToken"`
	UserSessionToken string `json:"userSessionToken"`
}

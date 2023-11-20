package types

type CsrfResponse struct {
	Data CsrfTokenResponse `json:"data,omitempty"`
}

type CsrfTokenResponse struct {
	CreateToken CsrfTokenContent `json:"createCsrfToken,omitempty"`
}

type CsrfTokenContent struct {
	Token           string `json:"csrfToken,omitempty"`
	AppSessionToken string `json:"appSessionToken,omitempty"`
}

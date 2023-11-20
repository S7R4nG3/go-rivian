package rivian

import (
	"encoding/json"
	"fmt"
	"rivian/internal/prompts"
	"rivian/internal/types"
	"rivian/internal/utils"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type Rivian struct {
	RequestTimeout   int
	Username         string
	Password         string
	VIN              string
	AccessToken      string
	RefreshToken     string
	CSRFToken        string
	SessionToken     string
	UserSessionToken string
	OTPToken         string
	Logger           *logrus.Logger
}

func (r *Rivian) Authenticate() {
	r.createCSRFToken()
	url := gateway
	headers := types.GraphqlDefaultHeaders{}
	headers.New()
	headers.Add("Csrf-Token", r.CSRFToken)
	headers.Add("A-Sess", r.SessionToken)
	headers.Add("Apollographql-Client-Name", apolloClientName)
	headers.Add("Dc-Cid", fmt.Sprintf("m-ios-%s", uuid.New()))

	operation := "Login"
	query := "mutation Login($email: String!, $password: String!) {\n  login(email: $email, password: $password) {\n    __typename\n    ... on MobileLoginResponse {\n      __typename\n      accessToken\n      refreshToken\n      userSessionToken\n    }\n    ... on MobileMFALoginResponse {\n      __typename\n      otpToken\n    }\n  }\n}"
	variables := map[string]interface{}{
		"email":    r.Username,
		"password": r.Password,
	}
	body := types.GraphqlBody{}
	body.New(operation, query, variables)

	statusCode, auth, err := utils.GraphqlQuery(headers.Map, url, body.Json, r.Logger)
	if err != nil || !utils.IsOk(statusCode) {
		r.Logger.Fatalf("Error with authentication request to GraphQL endpoint -- %v %s %v", statusCode, auth, err)
	}

	var authResponse types.AuthResponse
	err = json.Unmarshal(auth, &authResponse)
	if err != nil {
		r.Logger.Fatalf("Unable to unmarshal JSON authentication response -- %v", err)
	}
	r.AccessToken = authResponse.Data.Login.AccessToken
	r.RefreshToken = authResponse.Data.Login.RefreshToken
	if authResponse.Data.Login.OTPToken != "" {
		r.OTPToken = authResponse.Data.Login.OTPToken
		r.validateOTP()
	}
}

func (r *Rivian) createCSRFToken() {
	url := gateway
	headers := types.GraphqlDefaultHeaders{}
	headers.New()

	operation := "CreateCSRFToken"
	query := "mutation CreateCSRFToken{\n  createCsrfToken {\n    __typename\n    csrfToken\n    appSessionToken\n  }\n}"
	variables := map[string]interface{}{}
	body := types.GraphqlBody{}
	body.New(operation, query, variables)

	statusCode, csrf, err := utils.GraphqlQuery(headers.Map, url, body.Json, r.Logger)
	if err != nil || !utils.IsOk(statusCode) {
		r.Logger.Fatalf("Error requesting CSRF token at GraphQL endpoint -- %v %s %v", statusCode, csrf, err)
	}

	var csrfResponse types.CsrfResponse
	err = json.Unmarshal(csrf, &csrfResponse)
	if err != nil {
		r.Logger.Fatalf("Unable to unmarshal CSRF response body -- %v", err)
	}
	r.CSRFToken = csrfResponse.Data.CreateToken.Token
	r.SessionToken = csrfResponse.Data.CreateToken.AppSessionToken
}

func (r *Rivian) validateOTP() {
	url := gateway
	headers := types.GraphqlDefaultHeaders{}
	headers.New()
	headers.Add("Csrf-Token", r.CSRFToken)
	headers.Add("A-Sess", r.SessionToken)
	headers.Add("Apollographql-Client-Name", apolloClientName)

	otpCode, err := prompts.OTP()
	if err != nil {
		r.Logger.Fatalf("Error prompting for OTP -- %v", err)
	}

	operation := "LoginWithOTP"
	query := "mutation LoginWithOTP($email: String!, $otpCode: String!, $otpToken: String!) {\n  loginWithOTP(email: $email, otpCode: $otpCode, otpToken: $otpToken) {\n    __typename\n    ... on MobileLoginResponse {\n      __typename\n      accessToken\n      refreshToken\n      userSessionToken\n    }\n  }\n}"
	variables := map[string]interface{}{
		"email":    r.Username,
		"otpCode":  otpCode,
		"otpToken": r.OTPToken,
	}
	body := types.GraphqlBody{}
	body.New(operation, query, variables)

	statusCode, otp, err := utils.GraphqlQuery(headers.Map, url, body.Json, r.Logger)
	if err != nil || !utils.IsOk(statusCode) {
		r.Logger.Fatalf("Error with OTP validation to GraphQL endpoint -- %v %s %v", statusCode, otp, err)
	}

	var otpResponse types.OTPResponse
	err = json.Unmarshal(otp, &otpResponse)
	if err != nil {
		r.Logger.Fatalf("Error unmarshalling OTP Response -- %v", err)
	}
	r.AccessToken = otpResponse.Data.LoginWithOTP.AccessToken
	r.RefreshToken = otpResponse.Data.LoginWithOTP.RefreshToken
	r.UserSessionToken = otpResponse.Data.LoginWithOTP.UserSessionToken
}

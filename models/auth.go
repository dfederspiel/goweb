package models

import (
	"github.com/dgrijalva/jwt-go"
)

type AuthProfile struct {
	RoleRequired Role
}

type Claims struct {
	Email   string `json:"email"`
	Picture string `json:"picture"`
	Name    string `json:"name"`
	jwt.StandardClaims
}

type OIDCSettings struct {
	WebSettings WebSettings `json:"web"`
}

type WebSettings struct {
	ClientId          string   `json:"client_id"`
	ProjectId         string   `json:"project_id"`
	AuthUri           string   `json:"auth_uri"`
	TokenUri          string   `json:"token_uri"`
	AuthProvider      string   `json:"auth_provider_x509_cert_url"`
	ClientSecret      string   `json:"client_secret"`
	RedirectUris      []string `json:"redirect_uris"`
	JavaScriptOrigins []string `json:"javascript_origins"`
}

type AuthResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
	IdToken     string `json:"id_token"`
	State       string `json:"session_state"`
}

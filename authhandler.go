package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

var jwtKey = []byte(`-----BEGIN CERTIFICATE-----\nMIIDJjCCAg6gAwIBAgIIA1PM71I9gHUwDQYJKoZIhvcNAQEFBQAwNjE0MDIGA1UE\nAxMrZmVkZXJhdGVkLXNpZ25vbi5zeXN0ZW0uZ3NlcnZpY2VhY2NvdW50LmNvbTAe\nFw0yMDAyMDIwNDI5MzBaFw0yMDAyMTgxNjQ0MzBaMDYxNDAyBgNVBAMTK2ZlZGVy\nYXRlZC1zaWdub24uc3lzdGVtLmdzZXJ2aWNlYWNjb3VudC5jb20wggEiMA0GCSqG\nSIb3DQEBAQUAA4IBDwAwggEKAoIBAQC3r6hVJVlFGWCicrvWVrY9ykWK22z30nDs\naJIKzFam6rE0mOy7HQ+425BavKcMHup1O4QNGyatdhJ6YhdyqqadaQz9Q/MWSnsJ\nbQXKv6MscRFfOTnk5TBzpfjWGOAmoFicbBYt4zdPJmYSWI9gAlAhHT20AE/B+jRp\nYWJVI9a0et/AltxSdf32L1i0Ht9jCamjj8RIRzArCPXTCkAx7fd18/nUC6U5PC/5\ngLa8uPDbmH3TIeH2uLqfs34wbmWCpy6n/WDxQYoPkqktM0lqzh84GCZqMeKz6Jbp\nQLcraGOB6tMX93tU1fpWd0GNDI/P2JGnNDfBBlYaGeDnRLFLr4tdAgMBAAGjODA2\nMAwGA1UdEwEB/wQCMAAwDgYDVR0PAQH/BAQDAgeAMBYGA1UdJQEB/wQMMAoGCCsG\nAQUFBwMCMA0GCSqGSIb3DQEBBQUAA4IBAQALWk+nEuDemE5a4k3cjTMN5WPfYM9+\n3nxV519bMTWOK9o2Ikg0TcKgkLekOMVRlbWTjTlkPPInVOaC3aUGjgiysZlglnn/\ncFZoR36lfsvYx6Xhc548eH99S4vu6lbnVsnFmIWwEQ5Nr8j8bBzz/6v2/daLKr3Z\nhwmIWft2tYInymesINdtWpjXgu7Y8eu076swJqn+VCccZJveYY0i4VB9Px/YQbBx\nVcUhptYCjICc6bPI9Cvl52Ud80//+PcddlkZ+OqcmDB49eHyKVCJc94PfUsn1AXj\nssFFFBMmy0pEF8h1tVbVo9eXCHZzAijwoYXZu7SWJKh9+cU2GvradID6\n-----END CERTIFICATE-----\n`)

var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

//
type Claims struct {
	Email   string `json:"email"`
	Picture string `json:"picture"`
	Name    string `json:"name"`
	jwt.StandardClaims
}

func Welcome(w http.ResponseWriter, r *http.Request) {
	// We can obtain the session token from the requests cookies, which come with every request
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// For any other type of error, return a bad request status
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get the JWT string from the cookie
	tknStr := c.Value

	// Initialize a new instance of `Claims`
	claims := &Claims{}

	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {

		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	// Finally, return the welcome message to the user, along with their
	// username given in the token
	w.Write([]byte(fmt.Sprintf("Welcome %s!", claims.Name)))
}

func Refresh(w http.ResponseWriter, r *http.Request) {
	// (BEGIN) The code uptil this point is the same as the first part of the `Welcome` route
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	tknStr := c.Value
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// (END) The code uptil this point is the same as the first part of the `Welcome` route

	// We ensure that a new token is not issued until enough time has elapsed
	// In this case, a new token will only be issued if the old token is within
	// 30 seconds of expiry. Otherwise, return a bad request status
	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Now, create a new token for the current use, with a renewed expiration time
	expirationTime := time.Now().Add(5 * time.Minute)
	claims.ExpiresAt = expirationTime.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Set the new token as the users `session_token` cookie
	http.SetCookie(w, &http.Cookie{
		Name:    "jwt",
		Value:   tokenString,
		Expires: expirationTime,
	})
}

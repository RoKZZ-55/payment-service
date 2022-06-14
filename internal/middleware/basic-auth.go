package middleware

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"payment-service/config"
	"payment-service/internal/utils"
	"strings"
)

type Auth struct {
	config     *config.Config
	middleware *Middleware
}

func New(config *config.Config) *Auth {
	return &Auth{
		config: config,
	}
}

func (a *Auth) BasicAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		notAuth := []string{
			"/create-payment", "/get-payment-status-by-id/",
			"/get-payments-by-userid/", "/get-payments-by-email/", "/cancel-payment-by-id/",
		}
		requestPath := r.URL.Path
		for _, value := range notAuth {
			if strings.HasPrefix(requestPath, value) {
				next.ServeHTTP(w, r)
				return
			}
		}

		username, password, ok := r.BasicAuth()
		if !ok {
			utils.Error(w, r, http.StatusUnauthorized, errors.New("incorrect username or password"))
			return
		}

		usernameHash, err := bcrypt.GenerateFromPassword([]byte(username), bcrypt.MinCost)
		if err != nil {
			utils.Error(w, r, http.StatusUnprocessableEntity, err)
		}
		passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			utils.Error(w, r, http.StatusUnprocessableEntity, err)
		}
		if err := bcrypt.CompareHashAndPassword(usernameHash, []byte(a.config.Auth.Username)); err != nil {
			utils.Error(w, r, http.StatusUnauthorized, errors.New("incorrect username or password"))
			return
		}
		if err := bcrypt.CompareHashAndPassword(passwordHash, []byte(a.config.Auth.Password)); err != nil {
			utils.Error(w, r, http.StatusUnauthorized, errors.New("incorrect username or password"))
			return
		}
		next.ServeHTTP(w, r)
	})
}

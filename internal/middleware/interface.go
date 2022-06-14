package middleware

import "net/http"

type Middleware interface {
	BasicAuth(next http.Handler) http.Handler
}

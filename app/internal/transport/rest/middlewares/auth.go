package middlewares

import (
	"loginMicroservice/app/internal/configs"
	"loginMicroservice/app/internal/security"
	"net/http"
	"strings"
)

// Define our struct
type authMiddleware struct {
	secretKeyJWT string
}

// NewAuth ...
func NewAuth(cfg configs.Configure) *authMiddleware {
	return &authMiddleware{secretKeyJWT: cfg.GetKeyJWT()}
}

// Middleware function, which will be called for each request
func (amw *authMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authToken := r.Header.Get("Authorization")
		token := strings.Split(authToken, " ")
		if len(token) != 2 {
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
		department, err := security.ParseJWT(token[1], amw.secretKeyJWT)
		if err != nil {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		r.Header.Add("department", department)
		next.ServeHTTP(w, r)
	})
}

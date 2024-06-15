package middleware

import (
	"net/http"
	"os"
	"strings"

	"golang.org/x/crypto/ssh"
)

type AuthMiddleware struct {
	publicKey ssh.PublicKey
}

func NewAuthMiddleware() (*AuthMiddleware, error) {
	publicKeyStr := os.Getenv("PUBLIC_KEY")
	publicKey, _, _, _, err := ssh.ParseAuthorizedKey([]byte(publicKeyStr))
	if err != nil {
		return nil, err
	}

	return &AuthMiddleware{publicKey: publicKey}, nil
}

func (am *AuthMiddleware) ServeHTTP(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "missing authorization header", http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if !am.verifyToken(token) {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (am *AuthMiddleware) verifyToken(token string) bool {
	// Implement token verification using asymmetric cryptography
	return true
}

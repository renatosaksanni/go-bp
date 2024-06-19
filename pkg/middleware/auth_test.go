package middleware

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func generateRSAPublicPrivateKey() (string, *rsa.PrivateKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return "", nil, err
	}

	pubASN1, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		return "", nil, err
	}

	pubBytes := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubASN1,
	})

	return string(pubBytes), privateKey, nil
}

func generateToken(message string, privateKey *rsa.PrivateKey) (string, error) {
	hasher := sha256.New()
	hasher.Write([]byte(message))
	hashed := hasher.Sum(nil)

	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed)
	if err != nil {
		return "", err
	}

	return message + "." + base64.StdEncoding.EncodeToString(signature), nil
}

func TestAuthMiddleware(t *testing.T) {
	publicKey, privateKey, err := generateRSAPublicPrivateKey()
	if err != nil {
		t.Fatalf("Failed to generate RSA keys: %v", err)
	}

	// Set environment variable for public key
	os.Setenv("PUBLIC_KEY", publicKey)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	authMiddleware, err := NewAuthMiddleware()
	if err != nil {
		t.Fatalf("Failed to create auth middleware: %v", err)
	}
	wrappedHandler := authMiddleware.ServeHTTP(handler)

	t.Run("Unauthenticated request", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/", nil)
		rr := httptest.NewRecorder()
		wrappedHandler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusUnauthorized, rr.Code)
	})

	t.Run("Authenticated request with invalid token", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/", nil)
		req.Header.Set("Authorization", "Bearer invalid-token")
		rr := httptest.NewRecorder()
		wrappedHandler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusUnauthorized, rr.Code)
	})

	t.Run("Authenticated request with valid token", func(t *testing.T) {
		message := "test-message"
		token, err := generateToken(message, privateKey)
		if err != nil {
			t.Fatalf("Failed to generate token: %v", err)
		}

		req, _ := http.NewRequest("GET", "/", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		rr := httptest.NewRecorder()
		wrappedHandler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
	})
}

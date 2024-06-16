package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRateLimitMiddleware(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	rateLimiter := NewRateLimiter(1, 1)
	rr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rateLimitHandler := rateLimiter.RateLimitMiddleware(handler)
	rateLimitHandler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Test rate limit exceeded
	rr = httptest.NewRecorder()
	rateLimitHandler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusTooManyRequests {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusTooManyRequests)
	}

	// Test rate limit reset after time
	time.Sleep(time.Second)
	rr = httptest.NewRecorder()
	rateLimitHandler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

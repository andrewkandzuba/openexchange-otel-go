package routers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestHelloHandler(t *testing.T) {
    // Create a test router and recorder
    router := SetupRouter(gin.Default())
    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/hello", nil)

    // Serve the request and record the response
    router.ServeHTTP(w, req)

    // Check the HTTP status code
    if w.Code != http.StatusOK {
        t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
    }

    // Check the response body
    expectedBody := "{\"message\":\"Hello, World!\"}"
    if w.Body.String() != expectedBody {
        t.Errorf("Expected response body %s, but got %s", expectedBody, w.Body.String())
    }
}

func TestLivenessHandler(t *testing.T) {
    // Create a test router and recorder
    router := SetupRouter(gin.Default())
    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/probs/healthz", nil)

    // Serve the request and record the response
    router.ServeHTTP(w, req)

    // Check the HTTP status code
    if w.Code != http.StatusOK {
        t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
    }

    // Check the response body
    expectedBody := "{\"message\":\"Healthy!\"}"
    if w.Body.String() != expectedBody {
        t.Errorf("Expected response body %s, but got %s", expectedBody, w.Body.String())
    }
}

func TestReadinessHandler(t *testing.T) {
    // Create a test router and recorder
    router := SetupRouter(gin.Default())
    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/probs/readyz", nil)

    // Serve the request and record the response
    router.ServeHTTP(w, req)

    // Check the HTTP status code
    if w.Code != http.StatusOK {
        t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
    }

    // Check the response body
    expectedBody := "{\"message\":\"Ready!\"}"
    if w.Body.String() != expectedBody {
        t.Errorf("Expected response body %s, but got %s", expectedBody, w.Body.String())
    }
}
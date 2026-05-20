package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"testing"

	"backend/auth_service/config"
	"backend/auth_service/handlers"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) sqlmock.Sqlmock {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create sqlmock: %v", err)
	}

	dialector := postgres.New(postgres.Config{
		Conn:       db,
		DriverName: "postgres",
	})
	gormDB, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open gorm: %v", err)
	}

	config.DB = gormDB
	return mock
}

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	return r
}

// Test 1: Register — missing fields returns 400
func TestRegister_MissingFields(t *testing.T) {
	r := setupRouter()
	r.POST("/register", handlers.Register)

	body := `{"username": ""}`
	req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

// Test 2: Register — invalid JSON returns 400
func TestRegister_InvalidJSON(t *testing.T) {
	r := setupRouter()
	r.POST("/register", handlers.Register)

	req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBufferString(`{invalid}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

// Test 3: Login — missing fields returns 400
func TestLogin_MissingFields(t *testing.T) {
	r := setupRouter()
	r.POST("/login", handlers.Login)

	body := `{"username": "test"}`
	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

// Test 4: Login — invalid JSON returns 400
func TestLogin_InvalidJSON(t *testing.T) {
	r := setupRouter()
	r.POST("/login", handlers.Login)

	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBufferString(`not json`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

// Test 5: Logout — missing token returns 400
func TestLogout_MissingToken(t *testing.T) {
	r := setupRouter()
	r.POST("/logout", handlers.Logout)

	req := httptest.NewRequest(http.MethodPost, "/logout", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var resp map[string]string
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, "Invalid token", resp["error"])
}

// Test 6: Logout — invalid token format returns 400
func TestLogout_InvalidTokenFormat(t *testing.T) {
	r := setupRouter()
	r.POST("/logout", handlers.Logout)

	req := httptest.NewRequest(http.MethodPost, "/logout", nil)
	req.Header.Set("Authorization", "InvalidFormat")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

// Test 7: Logout — valid token format returns 200
func TestLogout_ValidToken(t *testing.T) {
	r := setupRouter()
	r.POST("/logout", handlers.Logout)

	req := httptest.NewRequest(http.MethodPost, "/logout", nil)
	req.Header.Set("Authorization", "Bearer somevalidtoken123")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp map[string]string
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, "Logged out successfully", resp["message"])
}

// Test 8: IsTokenBlacklisted — token not blacklisted by default
func TestIsTokenBlacklisted_NotBlacklisted(t *testing.T) {
	result := handlers.IsTokenBlacklisted("some_random_token")
	assert.False(t, result)
}

// Test 9: IsTokenBlacklisted — token blacklisted after logout
func TestIsTokenBlacklisted_AfterLogout(t *testing.T) {
	r := setupRouter()
	r.POST("/logout", handlers.Logout)

	token := "blacklisted_token_test"
	req := httptest.NewRequest(http.MethodPost, "/logout", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.True(t, handlers.IsTokenBlacklisted(token))
}

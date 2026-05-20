package handlers

import (
	"backend/todo_service/config"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) (sqlmock.Sqlmock, func()) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock DB: %v", err)
	}
	dialector := postgres.New(postgres.Config{Conn: db, DriverName: "postgres"})
	gormDB, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to open gorm DB: %v", err)
	}
	config.DB = gormDB
	return mock, func() { db.Close() }
}

func setupRouter(userID uint) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("user_id", userID)
		c.Next()
	})
	return r
}

// Test 1: CreateTodo — invalid JSON returns 400
func TestCreateTodo_InvalidJSON(t *testing.T) {
	_, cleanup := setupTestDB(t)
	defer cleanup()
	r := setupRouter(1)
	r.POST("/api/todos", CreateTodo)
	req, _ := http.NewRequest("POST", "/api/todos", bytes.NewBufferString("{invalid json"))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

// Test 2: CreateTodo — empty title returns 400 (binding:"required")
func TestCreateTodo_MissingTitle(t *testing.T) {
	_, cleanup := setupTestDB(t)
	defer cleanup()
	r := setupRouter(1)
	r.POST("/api/todos", CreateTodo)
	body, _ := json.Marshal(map[string]interface{}{"title": "", "description": "some desc"})
	req, _ := http.NewRequest("POST", "/api/todos", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

// Test 3: GetTodos — success returns 200
func TestGetTodos_Success(t *testing.T) {
	mock, cleanup := setupTestDB(t)
	defer cleanup()

	// 1. Основной запрос todos
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "todos"`)).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "title", "description", "completed", "user_id", "priority"}).
				AddRow(1, "Test todo", "desc", false, 1, "medium"),
		)

	// 2. Preload Tags — GORM делает через todo_tags
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "todo_tags"`)).
		WillReturnRows(sqlmock.NewRows([]string{"todo_id", "tag_id"}))

	// 3. Preload Category
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "categories"`)).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "color"}))

	r := setupRouter(1)
	r.GET("/api/todos", GetTodos)
	req, _ := http.NewRequest("GET", "/api/todos", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

// Test 4: GetTodoByID — not found returns 404
func TestGetTodoByID_NotFound(t *testing.T) {
	mock, cleanup := setupTestDB(t)
	defer cleanup()
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "todos"`)).
		WillReturnRows(sqlmock.NewRows([]string{"id"}))
	r := setupRouter(1)
	r.GET("/api/todos/:id", GetTodoByID)
	req, _ := http.NewRequest("GET", "/api/todos/999", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

// Test 5: UpdateTodo — not found returns 404
func TestUpdateTodo_NotFound(t *testing.T) {
	mock, cleanup := setupTestDB(t)
	defer cleanup()
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "todos"`)).
		WillReturnRows(sqlmock.NewRows([]string{"id"}))
	r := setupRouter(1)
	r.PUT("/api/todos/:id", UpdateTodo)
	body, _ := json.Marshal(map[string]interface{}{"title": "Updated"})
	req, _ := http.NewRequest("PUT", "/api/todos/999", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

// Test 6: DeleteTodo — not found returns 404
func TestDeleteTodo_NotFound(t *testing.T) {
	mock, cleanup := setupTestDB(t)
	defer cleanup()
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE "todos" SET "deleted_at"`)).
		WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectCommit()
	r := setupRouter(1)
	r.DELETE("/api/todos/:id", DeleteTodo)
	req, _ := http.NewRequest("DELETE", "/api/todos/999", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

// Test 7: CreateCategory — invalid JSON returns 400
func TestCreateCategory_InvalidJSON(t *testing.T) {
	_, cleanup := setupTestDB(t)
	defer cleanup()
	r := setupRouter(1)
	r.POST("/api/categories", CreateCategory)
	req, _ := http.NewRequest("POST", "/api/categories", bytes.NewBufferString("{bad"))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

// Test 8: GetCategories — success returns 200
func TestGetCategories_Success(t *testing.T) {
	mock, cleanup := setupTestDB(t)
	defer cleanup()
	rows := sqlmock.NewRows([]string{"id", "name", "color", "user_id", "is_default"}).
		AddRow(1, "Personal", "#4CAF50", 0, true).
		AddRow(2, "Work", "#F44336", 1, false)
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "categories"`)).WillReturnRows(rows)
	r := setupRouter(1)
	r.GET("/api/categories", GetCategories)
	req, _ := http.NewRequest("GET", "/api/categories", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

// Test 9: GetCategoryByID — not found returns 404
func TestGetCategoryByID_NotFound(t *testing.T) {
	mock, cleanup := setupTestDB(t)
	defer cleanup()
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "categories"`)).
		WillReturnRows(sqlmock.NewRows([]string{"id"}))
	r := setupRouter(1)
	r.GET("/api/categories/:id", GetCategoryByID)
	req, _ := http.NewRequest("GET", "/api/categories/999", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

// Test 10: DeleteCategory — not found returns 404
func TestDeleteCategory_NotFound(t *testing.T) {
	mock, cleanup := setupTestDB(t)
	defer cleanup()
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE "categories" SET "deleted_at"`)).
		WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectCommit()
	r := setupRouter(1)
	r.DELETE("/api/categories/:id", DeleteCategory)
	req, _ := http.NewRequest("DELETE", "/api/categories/999", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

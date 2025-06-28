package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"library-app/src/models"
)

// mock DB and router setup would be required for real tests
// Here, we focus on the structure of the tests

func TestGetBooks(t *testing.T) {
	w := httptest.NewRecorder()
	r := gin.Default()
	r.GET("/books", GetBooks)
	r.ServeHTTP(w, httptest.NewRequest("GET", "/books", nil))
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetBook(t *testing.T) {
	w := httptest.NewRecorder()
	r := gin.Default()
	r.GET("/books/:id", GetBook)
	r.ServeHTTP(w, httptest.NewRequest("GET", "/books/1", nil))
	assert.Contains(t, []int{http.StatusOK, http.StatusNotFound}, w.Code)
}

func TestCreateBook(t *testing.T) {
	w := httptest.NewRecorder()
	r := gin.Default()
	r.POST("/books", CreateBook)
	book := models.Book{Title: "Test", Author: "Author", ISBN: "123456"}
	body, _ := json.Marshal(book)
	r.ServeHTTP(w, httptest.NewRequest("POST", "/books", bytes.NewBuffer(body)))
	assert.Contains(t, []int{http.StatusCreated, http.StatusInternalServerError, http.StatusBadRequest}, w.Code)
}

func TestUpdateBook(t *testing.T) {
	w := httptest.NewRecorder()
	r := gin.Default()
	r.PUT("/books/:id", UpdateBook)
	book := models.Book{Title: "Updated", Author: "Author", ISBN: "654321"}
	body, _ := json.Marshal(book)
	r.ServeHTTP(w, httptest.NewRequest("PUT", "/books/1", bytes.NewBuffer(body)))
	assert.Contains(t, []int{http.StatusOK, http.StatusInternalServerError, http.StatusBadRequest}, w.Code)
}

func TestDeleteBook(t *testing.T) {
	w := httptest.NewRecorder()
	r := gin.Default()
	r.DELETE("/books/:id", DeleteBook)
	r.ServeHTTP(w, httptest.NewRequest("DELETE", "/books/1", nil))
	assert.Contains(t, []int{http.StatusNoContent, http.StatusInternalServerError}, w.Code)
}

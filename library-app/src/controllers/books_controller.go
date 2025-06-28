package controllers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "library-app/config"
    "library-app/src/models"
)

// GetBooks godoc
// @Summary      List all books
// @Description  Get all books from the library
// @Tags         books
// @Accept       json
// @Produce      json
// @Success      200  {array}  models.Book
// @Router       /books [get]
func GetBooks(c *gin.Context) {
    var books []models.Book
    err := config.DB.Select(&books, "SELECT * FROM books")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, books)
}

// GetBook godoc
// @Summary      Get a book by ID
// @Description  Get details of a book by its ID
// @Tags         books
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Book ID"
// @Success      200  {object}  models.Book
// @Failure      404  {object}  gin.H
// @Router       /books/{id} [get]
func GetBook(c *gin.Context) {
    id := c.Param("id")
    var book models.Book
    err := config.DB.Get(&book, "SELECT * FROM books WHERE id=$1", id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
        return
    }
    c.JSON(http.StatusOK, book)
}

// CreateBook godoc
// @Summary      Add a new book
// @Description  Create a new book by providing title, author, and ISBN
// @Tags         books
// @Accept       json
// @Produce      json
// @Param        book  body  models.Book  true  "Book object to be created"
// @Success      201   {object}  models.Book
// @Failure      400   {object}  gin.H
// @Failure      500   {object}  gin.H
// @Router       /books [post]
func CreateBook(c *gin.Context) {
    var book models.Book
    if err := c.ShouldBindJSON(&book); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    err := config.DB.QueryRowx(
        "INSERT INTO books(title, author, isbn) VALUES ($1, $2, $3) RETURNING id",
        book.Title, book.Author, book.ISBN).Scan(&book.ID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, book)
}

// UpdateBook godoc
// @Summary      Update a book
// @Description  Update a book's details by ID
// @Tags         books
// @Accept       json
// @Produce      json
// @Param        id    path      int         true  "Book ID"
// @Param        book  body      models.Book true  "Book object to update"
// @Success      200   {object}  models.Book
// @Failure      400   {object}  gin.H
// @Failure      500   {object}  gin.H
// @Router       /books/{id} [put]
func UpdateBook(c *gin.Context) {
    id := c.Param("id")
    var book models.Book
    if err := c.ShouldBindJSON(&book); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    bookID, _ := strconv.Atoi(id)
    book.ID = bookID

    _, err := config.DB.Exec(
        "UPDATE books SET title=$1, author=$2, isbn=$3 WHERE id=$4",
        book.Title, book.Author, book.ISBN, book.ID)

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, book)
}

// DeleteBook godoc
// @Summary      Delete a book
// @Description  Delete a book by its ID
// @Tags         books
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Book ID"
// @Success      204  {string}  string  "No Content"
// @Failure      500  {object}  gin.H
// @Router       /books/{id} [delete]
func DeleteBook(c *gin.Context) {
    id := c.Param("id")
    _, err := config.DB.Exec("DELETE FROM books WHERE id=$1", id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.Status(http.StatusNoContent)
}
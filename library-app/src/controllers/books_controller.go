package controllers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "library-app/config"
    "library-app/src/models"
)

func GetBooks(c *gin.Context) {
    var books []models.Book
    err := config.DB.Select(&books, "SELECT * FROM books")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, books)
}

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

func DeleteBook(c *gin.Context) {
    id := c.Param("id")
    _, err := config.DB.Exec("DELETE FROM books WHERE id=$1", id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.Status(http.StatusNoContent)
}
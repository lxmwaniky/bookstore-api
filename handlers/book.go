package handlers

import (
    "database/sql"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "github.com/lxmwaniky/bookstore-api/database"
    "github.com/lxmwaniky/bookstore-api/models"
)

func GetBooks(c *gin.Context) {
    rows, err := database.DB.Query("SELECT id, title, author, publication_year, genre, description, cover_image_url, created_at, updated_at FROM books")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch books"})
        return
    }
    defer rows.Close()

    var books []models.Book
    for rows.Next() {
        var book models.Book
        err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.PublicationYear, &book.Genre, &book.Description, &book.CoverImageURL, &book.CreatedAt, &book.UpdatedAt)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan book"})
            return
        }
        books = append(books, book)
    }

    c.JSON(http.StatusOK, books)
}

func GetBookByID(c *gin.Context) {
    id := c.Param("id")
    bookID, err := strconv.Atoi(id)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
        return
    }

    var book models.Book
    err = database.DB.QueryRow("SELECT id, title, author, publication_year, genre, description, cover_image_url, created_at, updated_at FROM books WHERE id = $1", bookID).Scan(
        &book.ID, &book.Title, &book.Author, &book.PublicationYear, &book.Genre, &book.Description, &book.CoverImageURL, &book.CreatedAt, &book.UpdatedAt)

    if err != nil {
        if err == sql.ErrNoRows {
            c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch book"})
        return
    }

    c.JSON(http.StatusOK, book)
}

func CreateBook(c *gin.Context) {
    var req models.CreateBookRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var book models.Book
    err := database.DB.QueryRow(
        "INSERT INTO books (title, author, publication_year, genre, description, cover_image_url) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, title, author, publication_year, genre, description, cover_image_url, created_at, updated_at",
        req.Title, req.Author, req.PublicationYear, req.Genre, req.Description, req.CoverImageURL).Scan(
        &book.ID, &book.Title, &book.Author, &book.PublicationYear, &book.Genre, &book.Description, &book.CoverImageURL, &book.CreatedAt, &book.UpdatedAt)

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create book"})
        return
    }

    c.JSON(http.StatusCreated, book)
}

func UpdateBook(c *gin.Context) {
    id := c.Param("id")
    bookID, err := strconv.Atoi(id)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
        return
    }

    var req models.CreateBookRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var book models.Book
    err = database.DB.QueryRow(
        "UPDATE books SET title = $1, author = $2, publication_year = $3, genre = $4, description = $5, cover_image_url = $6, updated_at = NOW() WHERE id = $7 RETURNING id, title, author, publication_year, genre, description, cover_image_url, created_at, updated_at",
        req.Title, req.Author, req.PublicationYear, req.Genre, req.Description, req.CoverImageURL, bookID).Scan(
        &book.ID, &book.Title, &book.Author, &book.PublicationYear, &book.Genre, &book.Description, &book.CoverImageURL, &book.CreatedAt, &book.UpdatedAt)

    if err != nil {
        if err == sql.ErrNoRows {
            c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update book"})
        return
    }

    c.JSON(http.StatusOK, book)
}

func DeleteBook(c *gin.Context) {
    id := c.Param("id")
    bookID, err := strconv.Atoi(id)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
        return
    }

    result, err := database.DB.Exec("DELETE FROM books WHERE id = $1", bookID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete book"})
        return
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get affected rows"})
        return
    }

    if rowsAffected == 0 {
        c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}
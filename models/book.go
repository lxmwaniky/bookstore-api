package models

import "time"

type Book struct {
	ID              int       `json:"id"`
	Title           string    `json:"title"`
	Author          string    `json:"author"`
	PublicationYear *int      `json:"publication_year"`
	Genre           *string   `json:"genre"`
	Description     *string   `json:"description"`
	CoverImageURL   *string   `json:"cover_image_url"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type CreateBookRequest struct {
	Title           string `json:"title" binding:"required"`
	Author          string `json:"author" binding:"required"`
	PublicationYear *int   `json:"publication_year"`
	Genre           *string `json:"genre"`
	Description     *string `json:"description"`
	CoverImageURL   *string `json:"cover_image_url"`
}
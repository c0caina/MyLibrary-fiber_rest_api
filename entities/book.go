package entities

import (
	"time"

	"github.com/google/uuid"
)

// Book is a structure that represents a book in a database
// It allows you to work with books as objects that have their own properties and methods
// It is also used to validate data about books to avoid errors or incorrect values
type Book struct {
	ID         uuid.UUID `db:"id" json:"id"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
	UpdatedAt  time.Time `db:"updated_at" json:"updated_at"`
	Title      string    `db:"title" json:"title" validate:"required,lte=255"`
	Author     string    `db:"author" json:"author" validate:"required,lte=255"`
	BookStatus int       `db:"book_status" json:"book_status" validate:"required,len=1"`
}

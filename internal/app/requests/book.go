package requests

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"github.com/c0caina/MyLibrary-fiber_rest_api/internal/app/models"
)

type Book struct {
	*pgx.Conn
}

func (b *Book) GetBooks() ([]models.Book, error) {
	books := []models.Book{}
	err := b.QueryRow(context.Background(), `SELECT * FROM books`).Scan(&books)
	return books, err
}

func (b *Book) GetBook(id uuid.UUID) (models.Book, error) {
	book := models.Book{}
	err := b.QueryRow(context.Background(), `SELECT * FROM books WHERE id = $1`, id).Scan(&book)
	return book, err
}

func (b *Book) CreateBook(mb *models.Book) error {
	_, err := b.Exec(context.Background(), `INSERT INTO books (title, author, book_status) VALUES ($1, $2, $3)`, mb.Title, mb.Author, mb.BookStatus,)
	return err
}

func (b *Book) UpdateBook(id uuid.UUID, mb *models.Book) error {
	_, err := b.Exec(context.Background(), `UPDATE books SET title = $2, author = $3, book_status = $4, WHERE id = $1`	, id, mb.Title, mb.Author, mb.BookStatus)
	return err
}

func (b *Book) DeleteBook(id uuid.UUID) error {
	_, err := b.Exec(context.Background(), `DELETE FROM books WHERE id = $1`, id)
	return err
}

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
	var books []models.Book
	rows, err := b.Query(context.Background(), `SELECT * FROM books`)
	if err != nil {
        return books, err
    }

	defer rows.Close()
	
	for rows.Next() {
		var book models.Book
		if err := rows.Scan(&book.ID, &book.CreatedAt, &book.UpdatedAt, &book.Title, &book.Author, &book.BookStatus); err != nil{
			return books, err
		}
		books = append(books, book)
	}


	return books, err
}

func (b *Book) GetBook(id uuid.UUID) (models.Book, error) {
	var book models.Book
	err := b.QueryRow(context.Background(), `SELECT * FROM books WHERE id = $1`, id).Scan(&book.ID, &book.CreatedAt, &book.UpdatedAt, &book.Title, &book.Author, &book.BookStatus)
	return book, err
}

func (b *Book) CreateBook(mb *models.Book) error {
	_, err := b.Exec(context.Background(), `INSERT INTO books (title, author, book_status) VALUES ($1, $2, $3)`, mb.Title, mb.Author, mb.BookStatus,)
	return err
}

func (b *Book) UpdateBook(mb *models.Book) error {
	_, err := b.Exec(context.Background(), `UPDATE books SET title = $2, author = $3, book_status = $4 WHERE id = $1`, mb.ID, mb.Title, mb.Author, mb.BookStatus)
	return err
}

func (b *Book) DeleteBook(id uuid.UUID) error {
	_, err := b.Exec(context.Background(), `DELETE FROM books WHERE id = $1`, id)
	return err
}

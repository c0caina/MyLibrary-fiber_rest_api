package postgresql

import (
	"context"
	"os"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"github.com/c0caina/MyLibrary-fiber_rest_api/internal/app/models"
)

type postgreSQL struct {
	*pgx.Conn
}

func NewPostgreSQL() (*postgreSQL, error) {
	conn, err := pgx.Connect(context.Background(), os.Getenv("POSTGRES_SERVER_URL"))
	if err != nil {
		defer conn.Close(context.Background())
		return &postgreSQL{}, err
	}

	return &postgreSQL{Conn: conn}, err
}

func (p *postgreSQL) GetBooks() ([]models.Book, error) {
	var books []models.Book
	rows, err := p.Query(context.Background(), `SELECT * FROM books`)
	if err != nil {
		return books, err
	}

	defer rows.Close()

	for rows.Next() {
		var book models.Book
		if err := rows.Scan(&book.ID, &book.CreatedAt, &book.UpdatedAt, &book.Title, &book.Author, &book.BookStatus); err != nil {
			return books, err
		}
		books = append(books, book)
	}

	return books, err
}

func (p *postgreSQL) GetBook(id uuid.UUID) (models.Book, error) {
	var book models.Book
	err := p.QueryRow(context.Background(), `SELECT * FROM books WHERE id = $1`, id).Scan(&book.ID, &book.CreatedAt, &book.UpdatedAt, &book.Title, &book.Author, &book.BookStatus)
	return book, err
}

func (p *postgreSQL) CreateBook(mb *models.Book) error {
	_, err := p.Exec(context.Background(), `INSERT INTO books (title, author, book_status) VALUES ($1, $2, $3)`, mb.Title, mb.Author, mb.BookStatus)
	return err
}

func (p *postgreSQL) UpdateBook(mb *models.Book) error {
	_, err := p.Exec(context.Background(), `UPDATE books SET title = $2, author = $3, book_status = $4 WHERE id = $1`, mb.ID, mb.Title, mb.Author, mb.BookStatus)
	return err
}

func (p *postgreSQL) DeleteBook(id uuid.UUID) error {
	_, err := p.Exec(context.Background(), `DELETE FROM books WHERE id = $1`, id)
	return err
}

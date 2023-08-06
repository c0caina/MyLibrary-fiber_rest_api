package postgres

import (
	"context"
	"os"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"github.com/c0caina/MyLibrary-fiber_rest_api/entities"
)

// postgres is a type of struct that embeds a pointer to a pgx.Conn object
// pgx.Conn is an object that represents a connection to a PostgreSQL database
// Embedding a pointer to pgx.Conn allows to call all its methods directly for postgres objects
// Also, additional methods can be added for the postgres struct that are specific to the application
type postgres struct {
	*pgx.Conn
}

// NewPostgres is a function that creates and returns a new postgres object
// It connects to the PostgreSQL database using the environment variable POSTGRES_SERVER_URL
// If there is an error, it closes the connection and returns an empty postgres object and the error
func NewPostgres() (*postgres, error) {
	conn, err := pgx.Connect(context.Background(), os.Getenv("POSTGRES_SERVER_URL"))
	if err != nil {
		defer conn.Close(context.Background())
		return &postgres{}, err
	}

	return &postgres{Conn: conn}, err
}

// GetBooks is a method of postgres that returns a slice of entities.Book objects
// It queries the books table and scans the rows into the slice
// If there is an error, it returns an empty slice and the error
func (p *postgres) GetBooks() ([]entities.Book, error) {
	var books []entities.Book
	rows, err := p.Query(context.Background(), `SELECT * FROM books`)
	if err != nil {
		return books, err
	}

	defer rows.Close()

	for rows.Next() {
		var book entities.Book
		if err := rows.Scan(&book.ID, &book.CreatedAt, &book.UpdatedAt, &book.Title, &book.Author, &book.BookStatus); err != nil {
			return books, err
		}
		books = append(books, book)
	}

	return books, err
}

// GetBook is a method of postgres that returns a single entities.Book object by id
// It queries the books table by id and scans the row into the book object
// If there is an error, it returns an empty book object and the error
func (p *postgres) GetBook(id uuid.UUID) (entities.Book, error) {
	var book entities.Book
	err := p.QueryRow(context.Background(), `SELECT * FROM books WHERE id = $1`, id).Scan(&book.ID, &book.CreatedAt, &book.UpdatedAt, &book.Title, &book.Author, &book.BookStatus)
	return book, err
}

// CreateBook is a method of postgres that creates a new book in the database using the entities.Book object passed as an argument 
// It executes an insert statement with the title, author and book_status fields of the book object 
// If there is an error, it returns the error
func (p *postgres) CreateBook(mb *entities.Book) error {
	_, err := p.Exec(context.Background(), `INSERT INTO books (title, author, book_status) VALUES ($1, $2, $3)`, mb.Title, mb.Author, mb.BookStatus)
	return err
}

// UpdateBook is a method of postgres that updates an existing book in the database using the entities.Book object passed as an argument 
// It executes an update statement with the id, title, author and book_status fields of the book object 
// If there is an error, it returns the error
func (p *postgres) UpdateBook(mb *entities.Book) error {
	_, err := p.Exec(context.Background(), `UPDATE books SET title = $2, author = $3, book_status = $4 WHERE id = $1`, mb.ID, mb.Title, mb.Author, mb.BookStatus)
	return err
}

// DeleteBook is a method of postgres that deletes an existing book in the database by id 
// It executes a delete statement with the id field 
// If there is an error, it returns the error 
func (p *postgres) DeleteBook(id uuid.UUID) error {
	_, err := p.Exec(context.Background(), `DELETE FROM books WHERE id = $1`, id)
	return err
}

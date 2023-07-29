package database

import "github.com/c0caina/MyLibrary-fiber_rest_api/internal/app/requests"

type Requests struct {
	*requests.Book
}

func OpenDBConnection() (*Requests, error) {
	conn, err := newPostgreSQL()
	if err != nil {
		return nil, err
	}

	return &Requests{
		Book: &requests.Book{Conn: conn},
	}, err
}
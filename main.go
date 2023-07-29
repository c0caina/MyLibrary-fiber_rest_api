package main

import (
	"fmt"

	"github.com/c0caina/MyLibrary-fiber_rest_api/internal/app/models"
	"github.com/c0caina/MyLibrary-fiber_rest_api/internal/pkg/database"
	"github.com/google/uuid"
)

func main()  {
	r, err := database.OpenDBConnection()
	if err != nil {
		fmt.Println("db err: ", err)
	}
	r.Book.CreateBook(&models.Book{ID: uuid.UUID{1}, })
	fmt.Println(r.Book.GetBooks())
	r.Book.DeleteBook(uuid.UUID{1})
	fmt.Println(r.Book.GetBooks())
}
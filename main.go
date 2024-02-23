package main

import (
	"log"

	"github.com/natanaelrusli/try-sqlmock/pkg/database"
	"github.com/natanaelrusli/try-sqlmock/repository"
)

func main() {
	db, err := database.InitPostgres()
	if err != nil {
		log.Fatalln(err.Error())
	}

	bookRepo := repository.NewBookRepository(db)

	books, err := bookRepo.GetAllBooks()
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(books)
}

package repository

import (
	"database/sql"

	"github.com/natanaelrusli/try-sqlmock/entity"
)

type bookRepo struct {
	db *sql.DB
}

type BookRepo interface {
	GetAllBooks() ([]entity.Book, error)
}

func NewBookRepository(db *sql.DB) BookRepo {
	return &bookRepo{
		db: db,
	}
}

func (r *bookRepo) GetAllBooks() ([]entity.Book, error) {
	rows, err := r.db.Query(`
		SELECT *
		FROM books;
	`)

	if err != nil {
		return nil, err
	}

	var books []entity.Book

	for rows.Next() {
		var book entity.Book

		err := rows.Scan(
			&book.Id,
			&book.Title,
			&book.Description,
			&book.Cover,
			&book.CreatedAt,
			&book.UpdatedAt,
			&book.DeletedAt,
			&book.AuthorID,
			&book.Stock,
		)

		if err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	return books, nil
}

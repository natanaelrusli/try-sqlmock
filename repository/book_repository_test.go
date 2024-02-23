package repository

import (
	"database/sql"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/natanaelrusli/try-sqlmock/entity"
	"github.com/stretchr/testify/assert"
)

func generateBook() entity.Book {
	return entity.Book{
		Id:          rand.Intn(9999),
		Title:       fmt.Sprintf("title %d", rand.Intn(9999)),
		Description: fmt.Sprintf("description %d", rand.Intn(9999)),
		Cover:       fmt.Sprintf("cover %d", rand.Intn(9999)),
		CreatedAt:   time.Now(),
		DeletedAt:   sql.NullTime{},
		UpdatedAt:   time.Now(),
		AuthorID:    int32(rand.Intn(9999)),
		Stock:       int32(rand.Intn(9999)),
	}
}

func TestGetAllBooks(t *testing.T) {
	var books []entity.Book

	book := generateBook()
	books = append(books, book)

	testCases := []struct {
		name          string
		expected      []entity.Book
		expectedQuery string
		expectedErr   error
	}{
		{
			name:          "success",
			expected:      books,
			expectedQuery: "SELECT * FROM books;",
			expectedErr:   nil,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			columns := []string{
				"id",
				"title",
				"description",
				"cover",
				"created_at",
				"updated_at",
				"deleted_at",
				"author_id",
				"stock",
			}

			expectedRows := sqlmock.NewRows(columns)

			for _, book := range test.expected {
				expectedRows.AddRow(
					book.Id,
					book.Title,
					book.Description,
					book.Cover,
					book.CreatedAt,
					book.UpdatedAt,
					book.DeletedAt,
					book.AuthorID,
					book.Stock,
				)
			}

			mock.
				ExpectQuery(test.expectedQuery).
				WillReturnRows(expectedRows)

			if err != nil {
				t.Fatalf("an error has occured: %s", err)
			}

			defer db.Close()

			repo := NewBookRepository(db)
			books, err = repo.GetAllBooks()

			assert.Len(t, books, len(test.expected))
			assert.Equal(t, test.expected, books)
			assert.Equal(t, test.expectedErr, err)

			if err != nil {
				t.Error(err)
			}
		})
	}
}

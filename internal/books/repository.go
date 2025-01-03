package books

import (
	"database/sql"

	"github.com/vferreirati/go_bookstore/internal/models"
)

type Repository interface {
	GetAll() ([]models.Book, error)
	CreateBook(name string, userID int) (int, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db}
}

func (r *repository) GetAll() ([]models.Book, error) {
	rows, err := r.db.Query("SELECT id, name FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var book models.Book
		if err := rows.Scan(&book.ID, &book.Name, &book.UserID); err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

func (r *repository) CreateBook(name string, userID int) (int, error) {
	query := "INSERT INTO books (name, user_id) VALUES ($1, $2)  RETURNING id"
	var id int
	err := r.db.QueryRow(query, name, userID).Scan(&id)
	if err != nil {
		return -1, err
	}

	return id, nil
}

package users

import (
	"database/sql"

	"github.com/vferreirati/go_bookstore/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type Repository interface {
	CreateUser(name, email, password string) (models.User, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db}
}

func (r *repository) CreateUser(name, email, password string) (models.User, error) {
	hashedPassword, err := hashPassword(password)
	if err != nil {
		return models.User{}, err
	}
	query := "INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id"
	var id int
	err = r.db.QueryRow(query, name, email, hashedPassword).Scan(&id)
	if err != nil {
		return models.User{}, err
	}

	return models.User{ID: id, Name: name, Email: email, Password: hashedPassword}, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

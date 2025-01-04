package users

import (
	"database/sql"

	"github.com/vferreirati/go_bookstore/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type Repository interface {
	CreateUser(name, email, password string) (models.User, error)
	GetByEmail(email string) (models.User, error)
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

func (r *repository) GetByEmail(email string) (models.User, error) {
	query := "SELECT id, name, email, password FROM users WHERE email = $1"
	var user models.User
	err := r.db.QueryRow(query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

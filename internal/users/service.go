package users

import "github.com/vferreirati/go_bookstore/internal/models"

type Service interface {
	CreateUser(name string, email string, password string) (models.User, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) CreateUser(name string, email string, password string) (models.User, error) {
	user, err := s.repo.CreateUser(name, email, password)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

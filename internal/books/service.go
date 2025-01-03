package books

import "github.com/vferreirati/go_bookstore/internal/models"

type Service interface {
	GetAll() ([]models.Book, error)
	CreateBook(name string, userID int) (models.Book, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) GetAll() ([]models.Book, error) {
	return s.repository.GetAll()
}

func (s *service) CreateBook(name string, userID int) (models.Book, error) {
	id, err := s.repository.CreateBook(name, userID)
	if err != nil {
		return models.Book{}, err
	}

	return models.Book{ID: id, Name: name, UserID: userID}, nil
}

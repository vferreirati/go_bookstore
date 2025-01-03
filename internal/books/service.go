package books

type Service interface {
	GetAll() ([]map[string]interface{}, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) GetAll() ([]map[string]interface{}, error) {
	return s.repository.GetAll()
}

package books

import "database/sql"

type Repository interface {
	GetAll() ([]map[string]interface{}, error)
	CreateBook(name string, userID string) (int, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db}
}

func (r *repository) GetAll() ([]map[string]interface{}, error) {
	rows, err := r.db.Query("SELECT id, name FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []map[string]interface{}
	for rows.Next() {
		var id, name string
		if err := rows.Scan(&id, &name); err != nil {
			return nil, err
		}
		books = append(books, map[string]interface{}{
			"id":   id,
			"name": name,
		})
	}
	return books, nil
}

func (r *repository) CreateBook(name string, userID string) (int, error) {
	query := "INSERT INTO books (name, user_id) VALUES ($1, $2)  RETURNING id"
	var id int
	err := r.db.QueryRow(query, name, userID).Scan(&id)
	if err != nil {
		return -1, err
	}

	return id, nil
}

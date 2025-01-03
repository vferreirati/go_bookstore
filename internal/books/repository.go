package books

import "database/sql"

type Repository interface {
	GetAll() ([]map[string]interface{}, error)
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

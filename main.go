package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func connectDatabase() (*sql.DB, error) {
	connStr := "postgres://postgres:postgres@localhost:5432/bookstore?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	db, err := connectDatabase()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	mux := http.NewServeMux()
	mux.Handle("GET /books", loggingMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handleListBooks(w, db)
	})))
	mux.Handle("GET /user/{id}/books", loggingMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handleBooksByUserId(w, r, db)
	})))

	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}

func handleBooksByUserId(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	w.Header().Set("Content-Type", "application/json")
	userId := r.PathValue("id")

	rows, err := db.Query("SELECT id, name FROM books WHERE user_id = $1", userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var books []map[string]interface{}
	for rows.Next() {
		var id, name string
		if err := rows.Scan(&id, &name); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		books = append(books, map[string]interface{}{
			"id":   id,
			"name": name,
		})
	}

	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(books)
}

func handleListBooks(w http.ResponseWriter, db *sql.DB) {
	w.Header().Set("Content-Type", "application/json")

	rows, err := db.Query("SELECT id, name FROM books")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var books []map[string]interface{}
	for rows.Next() {
		var id, name string
		if err := rows.Scan(&id, &name); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		books = append(books, map[string]interface{}{
			"id":   id,
			"name": name,
		})
	}

	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(books)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

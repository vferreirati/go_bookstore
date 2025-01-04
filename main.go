package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/vferreirati/go_bookstore/internal/books"
	"github.com/vferreirati/go_bookstore/internal/db"
	"github.com/vferreirati/go_bookstore/internal/users"

	_ "github.com/lib/pq"
)

func main() {
	db, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repository := books.NewRepository(db)
	service := books.NewService(repository)
	handler := books.NewHandler(service)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /books", handler.ListBooks)
	mux.HandleFunc("POST /books", handler.CreateBook)

	userRepository := users.NewRepository(db)
	userService := users.NewService(userRepository)
	userHandler := users.NewHandler(userService)
	mux.HandleFunc("POST /users", userHandler.HandleCreateUser)
	mux.HandleFunc("POST /login", userHandler.HandleLogin)

	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}

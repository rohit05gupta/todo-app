package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"to-do-app/internal/todo"

	"github.com/gorilla/mux"
)

func main() {
	// Set up the database connection
	// connStr := "user=postgres dbname=todoapp sslmode=disable password=pgsql host=localhost port=5432"
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		log.Fatal("DATABASE_URL is not set")
	}
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Initialize repository, service, and handler
	r := mux.NewRouter()
	todoRepo := todo.NewTodoRepository(db)
	attachmentRepo := todo.NewAttachmentRepository(db)
	serviceHandler := todo.NewServiceHandler(todoRepo, attachmentRepo)
	todoHandler := todo.NewHandler(serviceHandler, serviceHandler)

	r.HandleFunc("/todos", todoHandler.CreateTodoItem).Methods("POST")
	r.HandleFunc("/todos", todoHandler.GetTodoItems).Methods("GET")
	r.HandleFunc("/todos/{id}", todoHandler.GetTodoItem).Methods("GET")
	r.HandleFunc("/todos/{id}", todoHandler.UpdateTodoItem).Methods("PUT")
	r.HandleFunc("/todos/{id}", todoHandler.DeleteTodoItem).Methods("DELETE")
	r.HandleFunc("/attachments/{id}", todoHandler.UpdateAttachment).Methods("PUT")
	r.HandleFunc("/attachments/{id}", todoHandler.DeleteAttachment).Methods("DELETE")
	http.Handle("/", r)

	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

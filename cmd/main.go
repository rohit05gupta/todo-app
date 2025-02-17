package main

import (
	"fmt"
	"log"
	"net/http"
	"to-do-app/internal/todo"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	todoHandler := todo.NewHandler()

	r.HandleFunc("/todos", todoHandler.CreateTodo).Methods("POST")
	r.HandleFunc("/todos", todoHandler.GetTodos).Methods("GET")
	r.HandleFunc("/todos/{id}", todoHandler.GetTodo).Methods("GET")
	r.HandleFunc("/todos/{id}", todoHandler.UpdateTodo).Methods("PUT")
	r.HandleFunc("/todos/{id}", todoHandler.DeleteTodo).Methods("DELETE")

	http.Handle("/", r)

	fmt.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

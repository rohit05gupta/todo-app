package todo

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestCreateTodo(t *testing.T) {
	todo := Todo{
		Title:       "Test Todo",
		Description: "Test Description",
		Attachments: []string{"http://example.com/file1.jpg"},
	}

	todoJSON, err := json.Marshal(todo)
	if err != nil {
		t.Fatalf("Could not marshal todo: %v", err)
	}

	req, err := http.NewRequest("POST", "/todos", bytes.NewBuffer(todoJSON))
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	r := mux.NewRouter()
	todoHandler := NewHandler()
	r.HandleFunc("/todos", todoHandler.CreateTodo).Methods("POST")

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	// Check that the response status is 201 Created
	if rr.Code != http.StatusCreated {
		t.Errorf("Handler returned wrong status code: got %v want %v", rr.Code, http.StatusCreated)
	}

	// Check the response body
	var createdTodo Todo
	err = json.NewDecoder(rr.Body).Decode(&createdTodo)
	if err != nil {
		t.Fatalf("Could not decode response body: %v", err)
	}

	if createdTodo.Title != todo.Title {
		t.Errorf("Handler returned unexpected body: got %v want %v", createdTodo.Title, todo.Title)
	}
}

func TestCreateTodoInvalidJSON(t *testing.T) {
	req, err := http.NewRequest("POST", "/todos", bytes.NewBuffer([]byte("invalid json")))
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	r := mux.NewRouter()
	todoHandler := NewHandler()
	r.HandleFunc("/todos", todoHandler.CreateTodo).Methods("POST")

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	// Check that the response status is 400 Bad Request
	if rr.Code != http.StatusBadRequest {
		t.Errorf("Handler returned wrong status code: got %v want %v", rr.Code, http.StatusBadRequest)
	}
}

func TestCreateTodoMissingFields(t *testing.T) {
	todo := Todo{
		Description: "Test Description",
	}

	todoJSON, err := json.Marshal(todo)
	if err != nil {
		t.Fatalf("Could not marshal todo: %v", err)
	}

	req, err := http.NewRequest("POST", "/todos", bytes.NewBuffer(todoJSON))
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	r := mux.NewRouter()
	todoHandler := NewHandler()
	r.HandleFunc("/todos", todoHandler.CreateTodo).Methods("POST")

	// Record the response
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	// Check that the response status is 400 Bad Request
	if rr.Code != http.StatusBadRequest {
		t.Errorf("Handler returned wrong status code: got %v want %v", rr.Code, http.StatusBadRequest)
	}
}

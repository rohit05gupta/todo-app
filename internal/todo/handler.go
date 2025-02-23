package todo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
)

// Handler handles HTTP requests for Todo items.
type Handler struct {
	todoService       TodoService
	attachmentService AttachmentService
}

// NewHandler creates a new Handler.
func NewHandler(todoService TodoService, attachmentService AttachmentService) *Handler {
	return &Handler{todoService: todoService, attachmentService: attachmentService}
}

// CreateTodoItem handles creating a new Todo item.
func (h *Handler) CreateTodoItem(w http.ResponseWriter, r *http.Request) {
	//Limiting file size to 10MB
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	title := r.FormValue("title")
	description := r.FormValue("description")

	todo := &TodoItem{
		ID:          uuid.New().String(),
		Title:       title,
		Description: description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// Loop through the uploaded attachments
	var attachments []Attachment
	files := r.MultipartForm.File["file_data"]
	for _, fileHeader := range files {
		//opening the uploaded file
		file, err := fileHeader.Open()
		if err != nil {
			http.Error(w, "Error opening file", http.StatusInternalServerError)
			return
		}
		defer file.Close()

		//reading the file data
		fileData, err := ioutil.ReadAll(file)
		if err != nil {
			http.Error(w, "Error reading file", http.StatusInternalServerError)
			return
		}
		attachment := Attachment{
			ID:        uuid.New().String(),
			TodoID:    todo.ID,
			FileName:  fileHeader.Filename,
			FileType:  fileHeader.Header.Get("Content-Type"),
			FileData:  fileData,
			CreatedAt: time.Now(),
		}

		attachments = append(attachments, attachment)
	}

	todo.Attachments = attachments

	// Save the TodoItem and Attachments (in a single transaction)
	createdTodo, err := h.todoService.CreateTodoItem(todo)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating todo: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdTodo)
}

// GetTodoItem handles retrieving a specific Todo item by ID.
func (h *Handler) GetTodoItem(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/todos/")
	todo, err := h.todoService.GetTodoItem(id)
	if err != nil {
		http.Error(w, "Error fetching todo", http.StatusNotFound)
		return
	}
	if todo.ID == "" {
		http.Error(w, fmt.Sprintf("Todo with ID %v not found", id), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(todo)
}

// GetTodoItems handles retrieving all the Todo items.
func (h *Handler) GetTodoItems(w http.ResponseWriter, r *http.Request) {
	todos, err := h.todoService.GetTodoItems()
	if err != nil {
		http.Error(w, "Error fetching todo", http.StatusInternalServerError)
		return
	}
	if len(todos) == 0 {
		http.Error(w, "To do list is empty", http.StatusOK)
		return
	}
	json.NewEncoder(w).Encode(todos)
}

// UpdateTodoItem updates the specific Todo item by ID.
func (h *Handler) UpdateTodoItem(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/todos/")
	todo := &TodoItem{}
	ctype := r.Header.Get("Content-Type")
	if ctype != "application/json" {
		http.Error(w, "Update operation supports json payload", http.StatusBadRequest)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedTodo, err := h.todoService.UpdateTodoItem(id, todo)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error updating todo: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(updatedTodo)
}

// DeleteTodoItem deletes the specific Todo item by ID.
func (h *Handler) DeleteTodoItem(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/todos/")
	err := h.todoService.DeleteTodoItem(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error deleting todo: %v", err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// UpdateAttachment updates the specific Todo item by ID.
func (h *Handler) UpdateAttachment(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/attachments/")
	attachment := &Attachment{}
	file, _, err := r.FormFile("file_data")
	if err != nil && err != http.ErrMissingFile {
		http.Error(w, "File not found or invalid", http.StatusBadRequest)
		return
	}

	// Read the file content
	var fileData []byte
	if file != nil {
		defer file.Close()
		fileData, err = ioutil.ReadAll(file)
		if err != nil {
			http.Error(w, "Error reading file", http.StatusInternalServerError)
			return
		}
	}

	if file != nil {
		attachment.FileData = fileData
	} else {
		http.Error(w, "Error updating attachment", http.StatusInternalServerError)
	}

	updatedAttachment, err := h.attachmentService.UpdateAttachment(id, attachment)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error updating attchment: %v", err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(updatedAttachment)
}

// DeleteTodoItem deletes the specific Todo item by ID.
func (h *Handler) DeleteAttachment(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/attachments/")
	err := h.attachmentService.DeleteAttachment(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error deleting attchment: %v", err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

package todo

import "io"

// TodoService defines the business methods for TODO items and attachments.
type TodoService interface {
	CreateTodoItem(todo *TodoItem) (*TodoItem, error)
	UpdateTodoItem(id string, todo *TodoItem) (*TodoItem, error)
	DeleteTodoItem(id string) error
	GetTodoItem(id string) (*TodoItem, error)
	GetTodoItems() ([]*TodoItem, error)
}

type AttachmentService interface {
	AttachFileToTodoItem(todoID string, file io.Reader, fileName string) error
	UpdateAttachment(id string, attachment *Attachment) (*Attachment, error)
	DeleteAttachment(id string) error
}

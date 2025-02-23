package todo

import (
	"time"
)

// TodoItem represents a single TODO item.
type TodoItem struct {
	ID          string       `json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	Attachments []Attachment `json:"attachments"`
}

// Attachment represents a file attached to a TodoItem.
type Attachment struct {
	ID        any `json:"id,omitempty"`
	TodoID    any `json:"todo_id,omitempty"`
	FileName  any `json:"file_name,omitempty"`
	FileData  any `json:"file_data,omitempty"`
	FileType  any `json:"file_type,omitempty"`
	CreatedAt any `json:"created_at,omitempty"`
	UpdatedAt any `json:"updated_at,omitempty"`
}

package todo

import "time"

type Todo struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Attachments []string  `json:"attachments"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Request struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Attachments []string `json:"attachments"`
}

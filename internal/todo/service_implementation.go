package todo

import (
	"database/sql"
	"io"
	"time"

	"github.com/google/uuid"
)

// ServiceHandlerImpl is the concrete implementation of the TodoService and AttachmentService interface.
type ServiceHandlerImpl struct {
	TodoRepository       TodoRepository
	AttachmentRepository AttachmentRepository
}

// NewServiceHandler creates and returns a new ServiceHandlerImpl.
func NewServiceHandler(todoRepo TodoRepository, attachmentRepo AttachmentRepository) *ServiceHandlerImpl {
	return &ServiceHandlerImpl{TodoRepository: todoRepo, AttachmentRepository: attachmentRepo}
}

// CreateTodoItem creates a new Todo item and returns it.
func (s *ServiceHandlerImpl) CreateTodoItem(todo *TodoItem) (*TodoItem, error) {
	todoId := uuid.New().String()
	createdAt := time.Now()
	updatedAt := time.Now()
	todo.ID = todoId
	todo.CreatedAt = createdAt
	todo.UpdatedAt = updatedAt

	err := s.TodoRepository.Save(todo)
	if err != nil {
		return &TodoItem{}, err
	}

	if len(todo.Attachments) > 0 {
		for _, attachment := range todo.Attachments {
			attachment.ID = sql.NullString{
				String: uuid.New().String(),
				Valid:  true,
			}
			attachment.TodoID = sql.NullString{
				String: todoId,
				Valid:  true,
			}
			attachment.CreatedAt = sql.NullTime{
				Time:  time.Now(),
				Valid: true,
			}
			attachment.UpdatedAt = sql.NullTime{
				Time:  time.Now(),
				Valid: true,
			}
			err := s.AttachmentRepository.SaveAttachment(&attachment)
			if err != nil {
				return &TodoItem{}, err
			}
		}
	}
	return todo, nil
}

// UpdateTodoItem updates an existing Todo item.
func (s *ServiceHandlerImpl) UpdateTodoItem(id string, todo *TodoItem) (*TodoItem, error) {
	existingTodo, err := s.TodoRepository.FindTodoByID(id)
	if err != nil {
		return nil, err
	}

	if todo.Title != "" && todo.Description != "" {
		existingTodo.Title = todo.Title
		existingTodo.Description = todo.Description
		existingTodo.UpdatedAt = time.Now()
	} else if todo.Title != "" {
		existingTodo.Title = todo.Title
		existingTodo.UpdatedAt = time.Now()
	} else {
		existingTodo.Description = todo.Description
		existingTodo.UpdatedAt = time.Now()
	}

	err = s.TodoRepository.Update(existingTodo)
	if err != nil {
		return nil, err
	}

	return existingTodo, nil
}

// UpdateAttachment updates an existing Attachment.
func (s *ServiceHandlerImpl) UpdateAttachment(id string, attachment *Attachment) (*Attachment, error) {
	existingAttachment, err := s.AttachmentRepository.FindAttachmentByID(id)
	if err != nil {
		return nil, err
	}

	if attachment.FileData != nil {
		existingAttachment.FileData = attachment.FileData
	}

	err = s.AttachmentRepository.Update(existingAttachment)
	if err != nil {
		return nil, err
	}

	return existingAttachment, nil
}

// DeleteTodoItem deletes a Todo item by ID.
func (s *ServiceHandlerImpl) DeleteTodoItem(id string) error {
	return s.TodoRepository.Delete(id)
}

// DeleteAttachment deletes an Attachment by ID.
func (s *ServiceHandlerImpl) DeleteAttachment(id string) error {
	return s.AttachmentRepository.Delete(id)
}

// GetTodoItem retrieves a Todo item by ID.
func (s *ServiceHandlerImpl) GetTodoItem(id string) (*TodoItem, error) {
	return s.TodoRepository.FindTodoByID(id)
}

// GetTodoItems retrieves all Todo items
func (s *ServiceHandlerImpl) GetTodoItems() ([]*TodoItem, error) {
	return s.TodoRepository.GetAllToDos()
}

// AttachFileToTodoItem attaches a file to the Todo item.
func (s *ServiceHandlerImpl) AttachFileToTodoItem(todoID string, file io.Reader, fileName string) error {
	todo, err := s.TodoRepository.FindTodoByID(todoID)
	if err != nil {
		return err
	}

	attachment := &Attachment{
		ID: sql.NullString{
			String: uuid.New().String(),
			Valid:  true,
		},
		TodoID: sql.NullString{
			String: todo.ID,
			Valid:  true,
		},
		CreatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	}

	// Save the attachment in DB
	err = s.AttachmentRepository.SaveAttachment(attachment)
	if err != nil {
		return err
	}
	return nil
}

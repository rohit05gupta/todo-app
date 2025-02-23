package todo

import (
	"database/sql"
	"errors"

	_ "github.com/lib/pq"
)

// TodoRepositoryImpl is a concrete implementation of the TodoRepository interface using PostgreSQL.
type TodoRepositoryImpl struct {
	db *sql.DB
}

// AttachmentRepositoryImpl is a concrete implementation of the AttachmentRepository interface using PostgreSQL.
type AttachmentRepositoryImpl struct {
	db *sql.DB
}

// NewTodoRepository creates a new TodoRepositoryImpl instance.
func NewTodoRepository(db *sql.DB) *TodoRepositoryImpl {
	return &TodoRepositoryImpl{db: db}
}

func NewAttachmentRepository(db *sql.DB) *AttachmentRepositoryImpl {
	return &AttachmentRepositoryImpl{db: db}
}

// Save saves a TodoItem to the database.
func (r *TodoRepositoryImpl) Save(todo *TodoItem) error {
	query := `INSERT INTO todo_items (id, title, description, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.Exec(query, todo.ID, todo.Title, todo.Description, todo.CreatedAt, todo.UpdatedAt)
	return err
}

// FindByID retrieves a TodoItem by ID from the database.
func (r *TodoRepositoryImpl) FindTodoByID(id string) (*TodoItem, error) {
	query := `SELECT todo_items.id AS todo_id, todo_items.title, todo_items.description, todo_items.created_at AS todo_created_at, todo_items.updated_at AS todo_updated_at, attachments.id AS attachment_id,attachments.todo_id AS attachment_todo_id, attachments.file_name AS attachment_file_name, attachments.file_data AS attachment_file_data,attachments.file_type AS attachment_file_type, attachments.created_at AS attachment_created_at FROM todo_items LEFT JOIN attachments ON todo_items.id = attachments.todo_id WHERE todo_id = $1;`
	rows, err := r.db.Query(query, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	todo := &TodoItem{}
	attachment := Attachment{}
	for rows.Next() {
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.CreatedAt, &todo.UpdatedAt,
			&attachment.ID, &attachment.TodoID, &attachment.FileName, &attachment.FileData, &attachment.FileType, &attachment.CreatedAt)
		if err != nil {
			return nil, errors.New("todo item not found")
		}
		if attachment.ID != "" {
			todo.Attachments = append(todo.Attachments, attachment)
		}
	}
	return todo, nil
}

// FindByID retrieves a TodoItem by ID from the database.
func (r *AttachmentRepositoryImpl) FindAttachmentByID(id string) (*Attachment, error) {
	query := `SELECT attachments.id AS attachment_id,attachments.todo_id AS attachment_todo_id, attachments.file_name AS attachment_file_name, attachments.file_data AS attachment_file_data,attachments.file_type AS attachment_file_type, attachments.created_at AS attachment_created_at FROM attachments WHERE attachment_id = $1;`
	row := r.db.QueryRow(query, id)

	attachment := Attachment{}
	err := row.Scan(&attachment.ID, &attachment.TodoID, &attachment.FileName, &attachment.FileData, &attachment.FileType, &attachment.CreatedAt)
	if err != nil {
		return nil, errors.New("attachment not found")
	}
	return &attachment, nil
}

func (r *TodoRepositoryImpl) GetAllToDos() ([]*TodoItem, error) {
	query := `SELECT todo_items.id AS todo_id, todo_items.title, todo_items.description, todo_items.created_at AS todo_created_at, todo_items.updated_at AS todo_updated_at, attachments.id AS attachment_id,attachments.todo_id AS attachment_todo_id, attachments.file_name AS attachment_file_name, attachments.file_data AS attachment_file_data,attachments.file_type AS attachment_file_type, attachments.created_at AS attachment_created_at FROM todo_items LEFT JOIN attachments ON todo_items.id = attachments.todo_id`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []*TodoItem
	todoMap := make(map[string]TodoItem)
	for rows.Next() {
		var todo TodoItem
		var attachment Attachment

		err := rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.CreatedAt, &todo.UpdatedAt,
			&attachment.ID, &attachment.TodoID, &attachment.FileName, &attachment.FileData, &attachment.FileType, &attachment.CreatedAt)
		if err != nil {
			return nil, err
		}
		if len(todoMap) == 0 {
			todoMap[todo.ID] = todo
			if attachment.ID != "" {
				todo.Attachments = append(todo.Attachments, attachment)
			}
			todos = append(todos, &todo)
		} else if _, ok := todoMap[todo.ID]; !ok {
			todoMap[todo.ID] = todo
			if attachment.ID != "" {
				todo.Attachments = append(todo.Attachments, attachment)
			}
			todos = append(todos, &todo)
		} else {
			for _, v := range todos {
				if v.ID == attachment.TodoID.(string) && attachment.ID != "" {
					v.Attachments = append(v.Attachments, attachment)
				}
			}
		}
	}

	return todos, nil
}

// Update updates a TodoItem in the database.
func (r *TodoRepositoryImpl) Update(todo *TodoItem) error {
	query := `UPDATE todo_items SET title = $1, description = $2, updated_at = $3 WHERE id = $4`
	_, err := r.db.Exec(query, todo.Title, todo.Description, todo.UpdatedAt, todo.ID)
	return err
}

// Update updates a TodoItem in the database.
func (r *AttachmentRepositoryImpl) Update(attachment *Attachment) error {
	query := `UPDATE attachments SET file_data = $1, updated_at = $2 WHERE id = $3`
	_, err := r.db.Exec(query, attachment.FileData, attachment.UpdatedAt, attachment.ID)
	return err
}

// Delete removes a TodoItem from the database.
func (r *TodoRepositoryImpl) Delete(id string) error {
	query := `DELETE FROM todo_items WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

// Delete removes a TodoItem from the database.
func (r *AttachmentRepositoryImpl) Delete(id string) error {
	query := `DELETE FROM attachments WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

// SaveAttachment saves an attachment to the database.
func (r *AttachmentRepositoryImpl) SaveAttachment(attachment *Attachment) error {
	query := `INSERT INTO attachments (id, todo_id, file_name, file_type, file_data, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := r.db.Exec(query, attachment.ID, attachment.TodoID, attachment.FileName, attachment.FileType, attachment.FileData, attachment.CreatedAt, attachment.UpdatedAt)
	return err
}

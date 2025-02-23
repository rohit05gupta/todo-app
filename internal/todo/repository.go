package todo

// TodoRepository is the interface for all storage-related operations.
type TodoRepository interface {
	Save(todo *TodoItem) error
	Update(todo *TodoItem) error
	Delete(id string) error
	FindTodoByID(id string) (*TodoItem, error)
	GetAllToDos() ([]*TodoItem, error)
}

type AttachmentRepository interface {
	Delete(id string) error
	Update(attachment *Attachment) error
	FindAttachmentByID(id string) (*Attachment, error)
	SaveAttachment(attachment *Attachment) error
}

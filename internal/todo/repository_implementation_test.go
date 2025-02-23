package todo

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestSaveTodoItem(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer mockDB.Close()

	repo := NewTodoRepository(mockDB)
	todoItem := &TodoItem{
		ID:          uuid.New().String(),
		Title:       "Test Title",
		Description: "Test Description",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	mock.ExpectExec("INSERT INTO todo_items").
		WithArgs(todoItem.ID, todoItem.Title, todoItem.Description, todoItem.CreatedAt, todoItem.UpdatedAt).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.Save(todoItem)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestFindTodoByID(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer mockDB.Close()

	repo := NewTodoRepository(mockDB)
	todoID := uuid.New().String()
	mock.ExpectQuery("SELECT todo_items.id").
		WithArgs(todoID).
		WillReturnRows(sqlmock.NewRows([]string{"todo_id", "title", "description", "todo_created_at", "todo_updated_at", "attachment_id", "attachment_todo_id", "attachment_file_name", "attachment_file_data", "attachment_file_type", "attachment_created_at"}).
			AddRow(todoID, "Test Title", "Test Description", time.Now(), time.Now(), nil, nil, nil, nil, nil, nil))

	todoItem, err := repo.FindTodoByID(todoID)
	assert.NoError(t, err)
	assert.NotNil(t, todoItem)
	assert.Equal(t, todoID, todoItem.ID)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateTodoItemRepo(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer mockDB.Close()

	repo := NewTodoRepository(mockDB)
	todoItem := &TodoItem{
		ID:          uuid.New().String(),
		Title:       "Updated Title",
		Description: "Updated Description",
		UpdatedAt:   time.Now(),
	}

	mock.ExpectExec("UPDATE todo_items").
		WithArgs(todoItem.Title, todoItem.Description, todoItem.UpdatedAt, todoItem.ID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = repo.Update(todoItem)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteTodoItemRepo(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer mockDB.Close()

	repo := NewTodoRepository(mockDB)
	todoID := uuid.New().String()

	mock.ExpectExec("DELETE FROM todo_items").
		WithArgs(todoID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = repo.Delete(todoID)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSaveAttachment(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer mockDB.Close()

	repo := NewAttachmentRepository(mockDB)
	attachment := &Attachment{
		ID:        uuid.New().String(),
		TodoID:    uuid.New().String(),
		FileName:  "test.txt",
		FileType:  "text/plain",
		FileData:  []byte("test data"),
		CreatedAt: time.Now(),
	}

	mock.ExpectExec("INSERT INTO attachments").
		WithArgs(attachment.ID, attachment.TodoID, attachment.FileName, attachment.FileType, attachment.FileData, attachment.CreatedAt).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.SaveAttachment(attachment)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteAttachment(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer mockDB.Close()

	repo := NewAttachmentRepository(mockDB)
	attachmentID := uuid.New().String()

	mock.ExpectExec("DELETE FROM attachments").
		WithArgs(attachmentID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = repo.Delete(attachmentID)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

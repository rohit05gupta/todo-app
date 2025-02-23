package todo

import (
	"bytes"
	"io"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type MockTodoRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTodoRepositoryMockRecorder
}

type MockTodoRepositoryMockRecorder struct {
	mock *MockTodoRepository
}

func NewMockTodoRepository(ctrl *gomock.Controller) *MockTodoRepository {
	mock := &MockTodoRepository{ctrl: ctrl}
	mock.recorder = &MockTodoRepositoryMockRecorder{mock}
	return mock
}

func (m *MockTodoRepository) EXPECT() *MockTodoRepositoryMockRecorder {
	return m.recorder
}

func (m *MockTodoRepository) Save(todo *TodoItem) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", todo)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockTodoRepositoryMockRecorder) Save(todo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockTodoRepository)(nil).Save), todo)
}

func (m *MockTodoRepository) FindTodoByID(id string) (*TodoItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindTodoByID", id)
	ret0, _ := ret[0].(*TodoItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockTodoRepositoryMockRecorder) FindTodoByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindTodoByID", reflect.TypeOf((*MockTodoRepository)(nil).FindTodoByID), id)
}

func (m *MockTodoRepository) GetAllToDos() ([]*TodoItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllToDos")
	ret0, _ := ret[0].([]*TodoItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockTodoRepositoryMockRecorder) GetAllToDos() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllToDos", reflect.TypeOf((*MockTodoRepository)(nil).GetAllToDos))
}

func (m *MockTodoRepository) Update(todo *TodoItem) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", todo)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockTodoRepositoryMockRecorder) Update(todo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockTodoRepository)(nil).Update), todo)
}

func (m *MockTodoRepository) Delete(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockTodoRepositoryMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockTodoRepository)(nil).Delete), id)
}

type MockAttachmentRepository struct {
	ctrl     *gomock.Controller
	recorder *MockAttachmentRepositoryMockRecorder
}

type MockAttachmentRepositoryMockRecorder struct {
	mock *MockAttachmentRepository
}

func NewMockAttachmentRepository(ctrl *gomock.Controller) *MockAttachmentRepository {
	mock := &MockAttachmentRepository{ctrl: ctrl}
	mock.recorder = &MockAttachmentRepositoryMockRecorder{mock}
	return mock
}

func (m *MockAttachmentRepository) EXPECT() *MockAttachmentRepositoryMockRecorder {
	return m.recorder
}

func (m *MockAttachmentRepository) SaveAttachment(attachment *Attachment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveAttachment", attachment)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockAttachmentRepositoryMockRecorder) SaveAttachment(attachment interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveAttachment", reflect.TypeOf((*MockAttachmentRepository)(nil).SaveAttachment), attachment)
}

func (m *MockAttachmentRepository) FindAttachmentByID(id string) (*Attachment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAttachmentByID", id)
	ret0, _ := ret[0].(*Attachment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAttachmentRepositoryMockRecorder) FindAttachmentByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAttachmentByID", reflect.TypeOf((*MockAttachmentRepository)(nil).FindAttachmentByID), id)
}

func (m *MockAttachmentRepository) Update(attachment *Attachment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", attachment)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockAttachmentRepositoryMockRecorder) Update(attachment interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockAttachmentRepository)(nil).Update), attachment)
}

func (m *MockAttachmentRepository) Delete(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockAttachmentRepositoryMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockAttachmentRepository)(nil).Delete), id)
}

func TestCreateTodoItemService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTodoRepo := NewMockTodoRepository(ctrl)
	mockAttachmentRepo := NewMockAttachmentRepository(ctrl)
	h := NewServiceHandler(mockTodoRepo, mockAttachmentRepo)

	todo := &TodoItem{
		Title:       "Test Todo",
		Description: "Test Description",
	}

	mockTodoRepo.EXPECT().Save(gomock.Any()).Return(nil)

	createdTodo, err := h.CreateTodoItem(todo)
	assert.NoError(t, err)
	assert.NotNil(t, createdTodo)
	assert.Equal(t, todo.Title, createdTodo.Title)
	assert.Equal(t, todo.Description, createdTodo.Description)
}

func TestUpdateTodoItemService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTodoRepo := NewMockTodoRepository(ctrl)
	mockAttachmentRepo := NewMockAttachmentRepository(ctrl)
	h := NewServiceHandler(mockTodoRepo, mockAttachmentRepo)

	id := uuid.New().String()
	existingTodo := &TodoItem{
		ID:          id,
		Title:       "Old Title",
		Description: "Old Description",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	updatedTodo := &TodoItem{
		Title:       "New Title",
		Description: "New Description",
	}

	mockTodoRepo.EXPECT().FindTodoByID(id).Return(existingTodo, nil)
	mockTodoRepo.EXPECT().Update(gomock.Any()).Return(nil)

	result, err := h.UpdateTodoItem(id, updatedTodo)
	assert.NoError(t, err)
	assert.Equal(t, "New Title", result.Title)
	assert.Equal(t, "New Description", result.Description)
}

func TestDeleteTodoItemService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTodoRepo := NewMockTodoRepository(ctrl)
	mockAttachmentRepo := NewMockAttachmentRepository(ctrl)
	h := NewServiceHandler(mockTodoRepo, mockAttachmentRepo)

	id := uuid.New().String()
	mockTodoRepo.EXPECT().Delete(id).Return(nil)

	err := h.DeleteTodoItem(id)
	assert.NoError(t, err)
}

func TestGetTodoItemService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTodoRepo := NewMockTodoRepository(ctrl)
	mockAttachmentRepo := NewMockAttachmentRepository(ctrl)
	h := NewServiceHandler(mockTodoRepo, mockAttachmentRepo)

	id := uuid.New().String()
	expectedTodo := &TodoItem{
		ID:          id,
		Title:       "Test Title",
		Description: "Test Description",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	mockTodoRepo.EXPECT().FindTodoByID(id).Return(expectedTodo, nil)

	result, err := h.GetTodoItem(id)
	assert.NoError(t, err)
	assert.Equal(t, expectedTodo, result)
}

func TestAttachFileToTodoItem(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTodoRepo := NewMockTodoRepository(ctrl)
	mockAttachmentRepo := NewMockAttachmentRepository(ctrl)
	h := NewServiceHandler(mockTodoRepo, mockAttachmentRepo)

	todoID := uuid.New().String()
	existingTodo := &TodoItem{ID: todoID}
	mockTodoRepo.EXPECT().FindTodoByID(todoID).Return(existingTodo, nil)
	mockAttachmentRepo.EXPECT().SaveAttachment(gomock.Any()).Return(nil)

	fileContent := bytes.NewBufferString("test file content")
	err := h.AttachFileToTodoItem(todoID, io.NopCloser(fileContent), "test.txt")
	assert.NoError(t, err)
}

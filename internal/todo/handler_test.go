package todo

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type MockAttachmentService struct {
	ctrl     *gomock.Controller
	recorder *MockAttachmentServiceMockRecorder
}

type MockAttachmentServiceMockRecorder struct {
	mock *MockAttachmentService
}

func NewMockAttachmentService(ctrl *gomock.Controller) *MockAttachmentService {
	mock := &MockAttachmentService{ctrl: ctrl}
	mock.recorder = &MockAttachmentServiceMockRecorder{mock}
	return mock
}

func (m *MockAttachmentService) EXPECT() *MockAttachmentServiceMockRecorder {
	return m.recorder
}

func (m *MockAttachmentService) AttachFileToTodoItem(todoID string, file io.Reader, fileName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AttachFileToTodoItem", todoID, file, fileName)
	ret0, _ := ret[0].(error)
	return ret0
}

func (m *MockAttachmentService) UpdateAttachment(id string, attachment *Attachment) (*Attachment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAttachment", id, attachment)
	ret0, _ := ret[0].(*Attachment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (m *MockAttachmentService) DeleteAttachment(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAttachment", id)
	ret0, _ := ret[0].(error)
	return ret0
}

type MockTodoService struct {
	ctrl     *gomock.Controller
	recorder *MockTodoServiceMockRecorder
}

type MockTodoServiceMockRecorder struct {
	mock *MockTodoService
}

func NewMockTodoService(ctrl *gomock.Controller) *MockTodoService {
	mock := &MockTodoService{ctrl: ctrl}
	mock.recorder = &MockTodoServiceMockRecorder{mock}
	return mock
}

func (m *MockTodoService) EXPECT() *MockTodoServiceMockRecorder {
	return m.recorder
}

func (m *MockTodoService) CreateTodoItem(todo *TodoItem) (*TodoItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTodoItem", todo)
	ret0, _ := ret[0].(*TodoItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (m *MockTodoServiceMockRecorder) CreateTodoItem(todo interface{}) *gomock.Call {
	return m.mock.ctrl.RecordCallWithMethodType(m.mock, "CreateTodoItem", reflect.TypeOf((*MockTodoService)(nil).CreateTodoItem), todo)
}

func (m *MockTodoServiceMockRecorder) UpdateTodoItem(id, todo interface{}) *gomock.Call {
	return m.mock.ctrl.RecordCallWithMethodType(m.mock, "UpdateTodoItem", reflect.TypeOf((*MockTodoService)(nil).UpdateTodoItem), id, todo)
}

func (m *MockTodoServiceMockRecorder) DeleteTodoItem(id interface{}) *gomock.Call {
	return m.mock.ctrl.RecordCallWithMethodType(m.mock, "DeleteTodoItem", reflect.TypeOf((*MockTodoService)(nil).DeleteTodoItem), id)
}

func (m *MockTodoServiceMockRecorder) GetTodoItem(id interface{}) *gomock.Call {
	return m.mock.ctrl.RecordCallWithMethodType(m.mock, "GetTodoItem", reflect.TypeOf((*MockTodoService)(nil).GetTodoItem), id)
}

func (m *MockTodoServiceMockRecorder) GetTodoItems() *gomock.Call {
	return m.mock.ctrl.RecordCallWithMethodType(m.mock, "GetTodoItems", reflect.TypeOf((*MockTodoService)(nil).GetTodoItems))
}

func (m *MockAttachmentServiceMockRecorder) AttachFileToTodoItem(todoID, file, fileName interface{}) *gomock.Call {
	return m.mock.ctrl.RecordCallWithMethodType(m.mock, "AttachFileToTodoItem", reflect.TypeOf((*MockAttachmentService)(nil).AttachFileToTodoItem), todoID, file, fileName)
}

func (m *MockAttachmentServiceMockRecorder) UpdateAttachment(id, attachment interface{}) *gomock.Call {
	return m.mock.ctrl.RecordCallWithMethodType(m.mock, "UpdateAttachment", reflect.TypeOf((*MockAttachmentService)(nil).UpdateAttachment), id, attachment)
}

func (m *MockAttachmentServiceMockRecorder) DeleteAttachment(id interface{}) *gomock.Call {
	return m.mock.ctrl.RecordCallWithMethodType(m.mock, "DeleteAttachment", reflect.TypeOf((*MockAttachmentService)(nil).DeleteAttachment), id)
}

func (m *MockTodoService) GetTodoItem(id string) (*TodoItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTodoItem", id)
	ret0, _ := ret[0].(*TodoItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (m *MockTodoService) GetTodoItems() ([]*TodoItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTodoItems")
	ret0, _ := ret[0].([]*TodoItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (m *MockTodoService) UpdateTodoItem(id string, todo *TodoItem) (*TodoItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTodoItem", id, todo)
	ret0, _ := ret[0].(*TodoItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (m *MockTodoService) DeleteTodoItem(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTodoItem", id)
	ret0, _ := ret[0].(error)
	return ret0
}

func TestCreateTodoItem(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTodoService := NewMockTodoService(ctrl)
	mockAttachmentService := NewMockAttachmentService(ctrl)

	h := NewHandler(mockTodoService, mockAttachmentService)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	_ = writer.WriteField("title", "Test Todo")
	_ = writer.WriteField("description", "Test Description")

	fileWriter, err := writer.CreateFormFile("file_data", "test.txt")
	assert.NoError(t, err)
	_, err = fileWriter.Write([]byte("dummy file content"))
	assert.NoError(t, err)

	err = writer.Close()
	assert.NoError(t, err)

	req, err := http.NewRequest("POST", "/todos", body)
	assert.NoError(t, err)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	rec := httptest.NewRecorder()

	expectedTodo := &TodoItem{
		ID:          uuid.New().String(),
		Title:       "Test Todo",
		Description: "Test Description",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	mockTodoService.EXPECT().CreateTodoItem(gomock.Any()).Return(expectedTodo, nil)

	h.CreateTodoItem(rec, req)

	assert.Equal(t, http.StatusCreated, rec.Code)
}

func TestGetTodoItem(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTodoService := NewMockTodoService(ctrl)
	mockAttachmentService := NewMockAttachmentService(ctrl)

	h := NewHandler(mockTodoService, mockAttachmentService)

	todoID := uuid.New().String()
	expectedTodo := &TodoItem{
		ID:          todoID,
		Title:       "Test Todo",
		Description: "Test Description",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	mockTodoService.EXPECT().GetTodoItem(todoID).Return(expectedTodo, nil)

	req, err := http.NewRequest("GET", "/todos/"+todoID, nil)
	assert.NoError(t, err)
	rec := httptest.NewRecorder()

	h.GetTodoItem(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	body, _ := ioutil.ReadAll(rec.Body)
	var response TodoItem
	json.Unmarshal(body, &response)
	assert.Equal(t, todoID, response.ID)
}

func TestDeleteTodoItem(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTodoService := NewMockTodoService(ctrl)
	mockAttachmentService := NewMockAttachmentService(ctrl)

	h := NewHandler(mockTodoService, mockAttachmentService)

	todoID := uuid.New().String()
	mockTodoService.EXPECT().DeleteTodoItem(todoID).Return(nil)

	req, err := http.NewRequest("DELETE", "/todos/"+todoID, nil)
	assert.NoError(t, err)
	rec := httptest.NewRecorder()

	h.DeleteTodoItem(rec, req)

	assert.Equal(t, http.StatusNoContent, rec.Code)
}

func TestUpdateTodoItem(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTodoService := NewMockTodoService(ctrl)
	mockAttachmentService := NewMockAttachmentService(ctrl)

	h := NewHandler(mockTodoService, mockAttachmentService)

	todoID := uuid.New().String()
	updatedTodo := &TodoItem{
		Title:       "Updated Todo",
		Description: "Updated Description",
	}
	requestBody, _ := json.Marshal(updatedTodo)

	req, err := http.NewRequest("PUT", "/todos/"+todoID, bytes.NewBuffer(requestBody))
	assert.NoError(t, err)
	rec := httptest.NewRecorder()

	mockTodoService.EXPECT().UpdateTodoItem(todoID, gomock.Any()).Return(updatedTodo, nil)

	h.UpdateTodoItem(rec, req)

	assert.Equal(t, http.StatusCreated, rec.Code)
}

func TestGetTodoItem_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTodoService := NewMockTodoService(ctrl)
	mockAttachmentService := NewMockAttachmentService(ctrl)

	h := NewHandler(mockTodoService, mockAttachmentService)

	todoID := uuid.New().String()
	mockTodoService.EXPECT().GetTodoItem(todoID).Return(nil, errors.New("not found"))

	req, err := http.NewRequest("GET", "/todos/"+todoID, nil)
	assert.NoError(t, err)
	rec := httptest.NewRecorder()

	h.GetTodoItem(rec, req)

	assert.Equal(t, http.StatusNotFound, rec.Code)
}

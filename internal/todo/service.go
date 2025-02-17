package todo

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

var (
	todos = make(map[string]Todo)
	mutex sync.Mutex
)

func CreateTodo(title, description string, attachments []string) Todo {
	mutex.Lock()
	defer mutex.Unlock()

	id := fmt.Sprintf("%d", len(todos)+1)
	todo := Todo{
		ID:          id,
		Title:       title,
		Description: description,
		Attachments: attachments,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	todos[id] = todo
	return todo
}

func GetTodos() ([]Todo, error) {
	mutex.Lock()
	defer mutex.Unlock()

	var allTodos []Todo
	for _, todo := range todos {
		allTodos = append(allTodos, todo)
	}
	if len(allTodos) == 0 {
		return allTodos, errors.New("Todo list is empty")
	}
	return allTodos, nil
}

func GetTodoByID(id string) (Todo, error) {
	mutex.Lock()
	defer mutex.Unlock()

	todo, exists := todos[id]
	if !exists {
		return Todo{}, errors.New("todo not found")
	}
	return todo, nil
}

func UpdateTodo(id, title, description string, attachments []string) (Todo, error) {
	mutex.Lock()
	defer mutex.Unlock()

	todo, exists := todos[id]
	if !exists {
		return Todo{}, errors.New("todo not found")
	}

	todo.Title = title
	todo.Description = description
	todo.Attachments = attachments
	todo.UpdatedAt = time.Now()

	todos[id] = todo
	return todo, nil
}

func DeleteTodoByID(id string) error {
	mutex.Lock()
	defer mutex.Unlock()

	_, exists := todos[id]
	if !exists {
		return errors.New("todo not found")
	}

	delete(todos, id)
	return nil
}

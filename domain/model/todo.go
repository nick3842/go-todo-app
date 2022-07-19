package model

import (
	"math/rand"
	"golang.org/x/exp/slices"
)

type Todo struct {
	Task string
	ID TodoID
	Completed bool
}

func CreateTodo(task string) Todo {
	id := TodoID(rand.Int())
	return Todo{ task, id, false, }
}

// REPOSITORY --------------------------------
type TodoID int
type TodoRepository interface {
	CreateTodo(Todo) error
	GetTodos() ([]Todo, error)
	DeleteTodo(TodoID) error
}

type inMemoryTodoRepository struct {
	// todos map[int] inMemoryTodo
	// todos []Todo
}

func CreateInMemoryTodoRespository() []Todo {
	todos := []Todo{
		{ Task: "trash", ID: 5, Completed: false },
		{ Task: "dishes", ID: 4, Completed: true },
	}
	return todos
	// func GetTodos() ([]Todo, error) {
	// 	return todos, nil
	// }
	// func CreateTodo (newTodo Todo) error {
	// 	todos = append(todos, newTodo)
	// 	return nil
	// }
	// func DeleteTodo (id TodoID) error {
	// 	for index, todo := range todos {
	// 		if todo.ID == id {
	// 			todos = slices.Delete(todos, index, index+1)
	// 			break
	// 		}
	// 	}
	// 	return nil
	// }

	// return TodoRepository {
	// 	GetTodos,
	// 	CreateTodo,
	// 	DeleteTodo
	// }
}

// func CreateInMemoryTodoRespository() *inMemoryTodoRepository {
// 	return &inMemoryTodoRepository {
// 		todos : []Todo{
// 			{ Task: "trash", ID: 5, Completed: false },
// 			{ Task: "dishes", ID: 4, Completed: true },
// 		},
// 	}
// }

// func (i *inMemoryTodoRepository) GetTodos() ([]Todo, error) {
// 	return i.todos, nil
// }

// func (i *inMemoryTodoRepository) CreateTodo (newTodo Todo) error {
// 	i.todos = append(i.todos, newTodo)
// 	return nil
// }

// func (i *inMemoryTodoRepository) DeleteTodo (id TodoID) error {
// 	for index, todo := range i.todos {
// 		if todo.ID == id {
// 			i.todos = slices.Delete(i.todos, index, index+1)
// 			break
// 		}
// 	}
// 	return nil
// }

// APPLICATION -------------------------------
type Application struct {
	repo TodoRepository
}

func CreateApplication (repo TodoRepository) (Application) {
	return Application{
		repo: repo,
	}
}

func (a *Application) AddTodo(task string) {
	newTodo := CreateTodo(task)
	a.repo.CreateTodo(newTodo)
}

func (a *Application) GetTodos() ([]Todo, error) {
	return a.repo.GetTodos()
}

func (a *Application) DeleteTodo(id TodoID) error {
	return a.repo.DeleteTodo(id)
}
// APPLICATION -------------------------------

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

type TodoList struct {
	Name string
	List []Todo
}

func CreateTodo(task string) Todo {
	id := TodoID(rand.Int())
	return Todo{ task, id, false, }
}

func CreateTodoList(owner string, todos []Todo) TodoList {
	return TodoList{ owner, todos, }
}

// REPOSITORY --------------------------------
type TodoID int
type TodoRepository interface {
	CreateTodo(Todo) error
	GetTodos() ([]Todo, error)
	DeleteTodo(TodoID) error
}

type TodoListRepository interface {
	CreateTodoList(TodoList) error
	GetTodoList(string) (TodoList, error)
	GetTodoLists() ([]TodoList, error)
	DeleteTodoList(string) error
}

type inMemoryTodoRepository struct {
	// todos map[int] inMemoryTodo
	todos []Todo
}

// TODO ----------------------------------------------
func CreateInMemoryTodoRespository() *inMemoryTodoRepository {
	return &inMemoryTodoRepository {
		todos : []Todo{
			{ Task: "trash", ID: 5, Completed: false },
			{ Task: "dishes", ID: 4, Completed: true },
		},
	}
}

func (i *inMemoryTodoRepository) GetTodos() ([]Todo, error) {
	return i.todos, nil
}

func (i *inMemoryTodoRepository) CreateTodo (newTodo Todo) error {
	i.todos = append(i.todos, newTodo)
	return nil
}

func (i *inMemoryTodoRepository) DeleteTodo (id TodoID) error {
	for index, todo := range i.todos {
		if todo.ID == id {
			i.todos = slices.Delete(i.todos, index, index+1)
			break
		}
	}
	return nil
}

// TODOLIST -------------------------------------------
type inMemoryTodoListRepository struct {
	// todos map[int] inMemoryTodo
	todoLists []TodoList
}

func CreateInMemoryTodoListRespository() *inMemoryTodoListRepository {
	return &inMemoryTodoListRepository {
		todoLists : []TodoList{
      { Name: "Nick", List: []Todo{{ Task: "trash", ID: 5, Completed: false }} },
      { Name: "John", List: []Todo{{ Task: "dishes", ID: 9, Completed: true }} },
			// { Task: "dishes", ID: 4, Completed: true },
		},
	}
}

func (i *inMemoryTodoListRepository) GetTodoList(target string) (TodoList, error) {
  var result TodoList
  for _, tempList := range i.todoLists {
    if tempList.Name == target {
      result = tempList
    }
  }
  return result, nil
}

func (i *inMemoryTodoListRepository) GetTodoLists() ([]TodoList, error) {
	return i.todoLists, nil
}

func (i *inMemoryTodoListRepository) CreateTodoList (newTodoList TodoList) error {
	i.todoLists = append(i.todoLists, newTodoList)
	return nil
}

func (i *inMemoryTodoListRepository) DeleteTodoList (owner string) error {
	for index, todoList := range i.todoLists {
		if todoList.Name == owner {
			i.todoLists = slices.Delete(i.todoLists, index, index+1)
			break
		}
	}
	return nil
}

// APPLICATION -------------------------------
type Application struct {
	repo TodoRepository
	listRepo TodoListRepository
}

func CreateApplication (repo TodoRepository, listRepo TodoListRepository) (Application) {
	return Application{
		repo: repo,
		listRepo: listRepo,
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

func (a *Application) CreateTodoList(owner string, todos []Todo) {
  newList := CreateTodoList(owner, todos)
	a.listRepo.CreateTodoList(newList)
}

func (a *Application) GetTodoList(owner string) (TodoList, error) {
	return a.listRepo.GetTodoList(owner)
}

func (a *Application) GetTodoLists() ([]TodoList, error) {
	return a.listRepo.GetTodoLists()
}

func (a *Application) DeleteTodoList(owner string) error {
	return a.listRepo.DeleteTodoList(owner)
}
// APPLICATION -------------------------------

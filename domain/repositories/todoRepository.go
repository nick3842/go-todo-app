package repositories

import "todo-app/domain/model"

type TodoRepository interface {
	CreateTodo(model.Todo) error
}

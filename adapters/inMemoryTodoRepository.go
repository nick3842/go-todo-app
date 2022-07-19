type inMemoryTodoRepository struct {
	todos map[int] inMemoryTodo
}

type inMemoryTodo struct {
	Task string
	ID int
	Completed bool
}

func (r *inMemoryTodoRepository) CreateTodo(todo model.Todo) error

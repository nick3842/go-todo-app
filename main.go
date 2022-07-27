package main

import (
	"todo-app/ports"
	"todo-app/domain/model"
)

// TODO: add multiple todo lists
// TODO: implement a mysql repository
// TODO: add a wamp port

var todoRepo = model.CreateInMemoryTodoRespository()
var todoListrepo = model.CreateInMemoryTodoListRespository()
var app = model.CreateApplication(todoRepo, todoListrepo)

func main() {
	ports.NewHttpPort(app)
}


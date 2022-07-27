package ports

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"fmt"
	"strconv"
	"todo-app/domain/model"
)

func createHandler(app model.Application) func(c *gin.Context) {
	return func (c *gin.Context) {
		body := struct {
			NewTodo string `json:"Task" binding:"required"`
		}{}

		if err := c.ShouldBindJSON(&body); err != nil {
			c.String(500, "couldn't whatever")
			return
		}

		app.AddTodo(body.NewTodo)

		c.String(200, "ok")
	}
}

func listHandler(app model.Application) func(c *gin.Context) {
	return func (c *gin.Context) {
		result, err := app.GetTodos()
		if err!=nil {
			c.String(500, "blew up")
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"todos": result,
		})
	}
}

func deleteHandler(app model.Application) func(*gin.Context) {
	return func (c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.String(500, "id not an int")
			return
		}
		app.DeleteTodo(model.TodoID(id))
		c.String(200, "ok")
	}
}

type Bar struct {
	Task string `json:"Task" binding:"required"`
	ID uint64 `json:"ID" binding:"required"`
	Completed bool `json:"Completed" binding:"required"`
}

type CreateTodoListHandlerBody struct {
	Owner string `json:"Name" binding:"required"`
	Todos []Bar `json:"Todos" binding:"required"`
}

func ConvertToTodoList(b []Bar) []model.Todo {
	var newTodos []model.Todo

	for _, val := range b {
		tempTodo := model.CreateTodo(val.Task)
		newTodos = append(newTodos, tempTodo)
	}

	return newTodos
}

func createTodoListHandler(app model.Application) func(*gin.Context) {
	return func (c *gin.Context) {

		var body CreateTodoListHandlerBody


		if err := c.ShouldBindJSON(&body); err != nil {
			c.String(500, "couldn't whatever")
			return
		}

		app.CreateTodoList(body.Owner, ConvertToTodoList(body.Todos))

		c.String(200, "ok")

	}
}

func getTodoListsHandler(app model.Application) func(*gin.Context) {
  return func (c *gin.Context){
    app.GetTodoLists()
    c.String(200, "ok")
  }
}

func getTodoListHandler(app model.Application) func(*gin.Context) {
  return func (c *gin.Context){
    idStr := c.Param("id")
    app.GetTodoList(idStr)
    c.String(200, "ok")
  }
}

func deleteTodoListHandler(app model.Application) func(*gin.Context) {
	return func (c *gin.Context) {
		idStr := c.Param("id")
		// if err != nil {
		// 	c.String(500, "id not an int")
		// 	return
		// }
		app.DeleteTodoList(idStr)
		c.String(200, "ok")
	}
}

func NewHttpPort (app model.Application){
	r := gin.Default()
	r.Use(CORSMiddleware())

  // regular todos
	r.POST("/create", createHandler(app))

	r.GET("/list", listHandler(app))

	r.GET("/delete/:id", deleteHandler(app))

  // todolist
  r.POST("/createTodoList/", createTodoListHandler(app))

  r.GET("/getTodoList/:id", getTodoListHandler(app))

	r.GET("/getTodoLists/", getTodoListsHandler(app))

  r.GET("/deleteTodoList/:id", deleteTodoListHandler(app))


	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

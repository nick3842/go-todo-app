package ports

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"fmt"
	"strconv"
	"todo-app/domain/model"
)

var repo = model.CreateInMemoryTodoRespository()
var app = model.CreateApplication(repo)

func createHandler(c *gin.Context) {

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

func listHandler(c *gin.Context) {
	result, err := app.GetTodos()
	if err!=nil {
		c.String(500, "blew up")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"todos": result,
	})
}

func deleteHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.String(500, "id not an int")
		return
	}

	app.DeleteTodo(model.TodoID(id))

	c.String(200, "ok")
}

func NewHttpPort (){
	r := gin.Default()
	r.Use(CORSMiddleware())

	r.POST("/create", createHandler)

	r.GET("/list", listHandler)

	r.GET("/delete/:id", deleteHandler)

	// r.POST("/update", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })

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

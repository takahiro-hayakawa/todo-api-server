package main

import (
	"github.com/gin-gonic/gin"
	"github.com/takahiro-hayakawa/todo-api-server/handler"
	"github.com/takahiro-hayakawa/todo-api-server/injector"
)

func main() {
	todoHandler := injector.InjectTodoHandler()
	engine := gin.Default()
	handler.InitRouting(engine, todoHandler)
	engine.Run(":3000")
}

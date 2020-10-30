package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/takahiro-hayakawa/todo-api-server/usecase"
	"net/http"
)

type TodoHandler struct {
	todoUsecase usecase.TodoUsecase
}

func NewTodoHandler(todoUsecase usecase.TodoUsecase) TodoHandler {
	todoHandler := TodoHandler{todoUsecase: todoUsecase}
	return todoHandler
}

func (todoHandler *TodoHandler) View() gin.HandlerFunc {
	todos, err := todoHandler.todoUsecase.View()
	if err != nil {
		fmt.Println(err)
	}

	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"todos": todos,
		})
	}
}

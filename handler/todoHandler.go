package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/takahiro-hayakawa/todo-api-server/usecase"
	"net/http"
	"strconv"
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

func (todoHandler *TodoHandler) Find() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, _ := strconv.Atoi(idStr)
		todos, err := todoHandler.todoUsecase.Find(id)
		if err != nil {
			fmt.Println(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"todos": todos,
		})
	}
}

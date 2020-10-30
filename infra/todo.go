package infra

import (
	"fmt"
	"github.com/takahiro-hayakawa/todo-api-server/domain/model"
	"github.com/takahiro-hayakawa/todo-api-server/domain/repository"
)

type TodoRepository struct {
	SqlHandler
}

func NewTodoRepository(sqlHandler SqlHandler) repository.TodoRepository {
	todoRepository := TodoRepository{sqlHandler}
	return &todoRepository
}

func (todoRepo *TodoRepository) FindAll() (todos []*model.Todo, err error) {
	rows, err := todoRepo.SqlHandler.Conn.Query("SELECT * FROM todo")
	defer rows.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	for rows.Next() {
		todo := model.Todo{}
		rows.Scan(&todo.ID, &todo.Task, &todo.LimitDate, &todo.Status)
		todos = append(todos, &todo)
	}
	return
}

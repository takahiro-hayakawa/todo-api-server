package stub

import (
	"github.com/takahiro-hayakawa/todo-api-server/domain/model"
	"github.com/takahiro-hayakawa/todo-api-server/domain/repository"
)

type TodoRepository struct {
}

func NewTodoRepository() repository.TodoRepository {
	todoReposity := TodoRepository{}
	return &todoReposity
}

func (todoRepo *TodoRepository) FindAll() (todos []*model.Todo, err error) {

	todo1 := model.Todo{ID: 1, Task: "本を読む", LimitDate: 1604081293, Status: 1}
	todos = append(todos, &todo1)

	todo2 := model.Todo{ID: 2, Task: "料理する", LimitDate: 1604081434, Status: 0}
	todos = append(todos, &todo2)

	return
}

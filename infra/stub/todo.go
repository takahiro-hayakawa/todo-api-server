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
	todos = []*model.Todo{
		&model.Todo{ID: 1, Task: "本を読む", LimitDate: 1604081293, Status: 1},
		&model.Todo{ID: 2, Task: "料理する", LimitDate: 1604081434, Status: 0},
	}
	return
}

func (todoRepo *TodoRepository) FindById(id int) (todo *model.Todo, err error) {
	todoMap := map[int]*model.Todo{
		1: &model.Todo{ID: 1, Task: "本を読む", LimitDate: 1604081293, Status: 1},
		2: &model.Todo{ID: 2, Task: "料理する", LimitDate: 1604081434, Status: 0},
	}
	todo = todoMap[id]
	return
}

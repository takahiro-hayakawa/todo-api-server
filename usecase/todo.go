package usecase

import (
	"github.com/takahiro-hayakawa/todo-api/domain/model"
	"github.com/takahiro-hayakawa/todo-api/domain/repository"
)

type TodoUsecase interface {
	View() (todo []*model.Todo, err error)
}

type todoUsecase struct {
	todoRepo repository.TodoRepository
}

func NewTodoUsecase(todoRepo repository.TodoRepository) TodoUsecase {
	todoUsecase := todoUsecase{todoRepo: todoRepo}
	return &todoUsecase
}

func (usecase *todoUsecase) View() (todos []*model.Todo, err error) {
	todos, err = usecase.todoRepo.FindAll()
	return
}

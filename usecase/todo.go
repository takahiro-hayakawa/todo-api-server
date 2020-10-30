package usecase

import (
	"github.com/takahiro-hayakawa/todo-api-server/domain/model"
	"github.com/takahiro-hayakawa/todo-api-server/domain/repository"
)

type TodoUsecase interface {
	View() (todo []*model.Todo, err error)
	Find(id int) (todo *model.Todo, err error)
}

type todoUsecase struct {
	todoRepo repository.TodoRepository
}

func (usecase *todoUsecase) Id(id int) (todo *model.Todo, err error) {
	panic("implement me")
}

func NewTodoUsecase(todoRepo repository.TodoRepository) TodoUsecase {
	todoUsecase := todoUsecase{todoRepo: todoRepo}
	return &todoUsecase
}

func (usecase *todoUsecase) View() (todos []*model.Todo, err error) {
	todos, err = usecase.todoRepo.FindAll()
	return
}

func (usecase *todoUsecase) Find(id int) (todo *model.Todo, err error) {
	todo, err = usecase.todoRepo.FindById(id)
	return
}

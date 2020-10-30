package repository

import (
	"github.com/takahiro-hayakawa/todo-api-server/domain/model"
)

type TodoRepository interface {
	FindAll() (todos []*model.Todo, err error)
	FindById(id int) (todo *model.Todo, err error)
}

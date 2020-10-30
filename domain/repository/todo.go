package repository

import (
	"github.com/takahiro-hayakawa/todo-api/domain/model"
)

type TodoRepository interface {
	FindAll() (todos []*model.Todo, err error)
}

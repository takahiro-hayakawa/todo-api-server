package injector

import (
	"github.com/takahiro-hayakawa/todo-api-server/domain/repository"
	"github.com/takahiro-hayakawa/todo-api-server/handler"
	"github.com/takahiro-hayakawa/todo-api-server/infra"
	"github.com/takahiro-hayakawa/todo-api-server/infra/stub"
	"github.com/takahiro-hayakawa/todo-api-server/usecase"
)

func InjectDB() infra.SqlHandler {
	sqlHandler := infra.NewSqlHandler()
	return *sqlHandler
}

func InjectTodoRepository() repository.TodoRepository {
	sqlHandler := InjectDB()
	return infra.NewTodoRepository(sqlHandler)
}

func InjectStubData() repository.TodoRepository {
	return stub.NewTodoRepository()
}

func InjectTodoUsecase() usecase.TodoUsecase {
	TodoRepo := InjectStubData()
	return usecase.NewTodoUsecase(TodoRepo)
}

func InjectTodoHandler() handler.TodoHandler {
	return handler.NewTodoHandler(InjectTodoUsecase())
}

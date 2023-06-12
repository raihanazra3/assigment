package repository

import (
	"assigment/domain_todocore/model/entity"
	"assigment/domain_todocore/model/vo"
	"context"
)

type SaveTodoRepo interface {
	SaveTodo(ctx context.Context, obj *entity.Todo) error
}

type FindAllTodoRepo interface {
	FindAllTodo(ctx context.Context, page, size int, someID string) ([]*entity.Todo, int64, error)
}

type FindOneTodoByIDRepo interface {
	FindOneTodoByID(ctx context.Context, todoID vo.TodoID) (*entity.Todo, error)
}

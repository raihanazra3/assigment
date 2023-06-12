package runtodocreate

import (
	"assigment/domain_todocore/model/entity"
	"context"
)

type runTodoCreateInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &runTodoCreateInteractor{
		outport: outputPort,
	}
}

func (r *runTodoCreateInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	// code your usecase definition here ...

	todoObj, err := entity.NewTodo(req.TodoCreateRequest)
	if err != nil {
		return nil, err
	}

	err = r.outport.SaveTodo(ctx, todoObj)
	if err != nil {

		res.Todo = todoObj
		return nil, err

	}

	//!

	return res, nil
}

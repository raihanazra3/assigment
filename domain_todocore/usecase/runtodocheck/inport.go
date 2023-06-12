package runtodocheck

import (
	"assigment/domain_todocore/model/entity"
	"assigment/domain_todocore/model/vo"
	"assigment/shared/gogen"
)

type Inport = gogen.Inport[InportRequest, InportResponse]

type InportRequest struct {
	TodoID vo.TodoID
}

type InportResponse struct {
	Todo *entity.Todo
}

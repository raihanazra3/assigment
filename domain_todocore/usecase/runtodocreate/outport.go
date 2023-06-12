package runtodocreate

import "assigment/domain_todocore/model/repository"

type Outport interface {
	repository.SaveTodoRepo
}

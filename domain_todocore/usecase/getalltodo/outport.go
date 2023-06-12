package getalltodo

import "assigment/domain_todocore/model/repository"

type Outport interface {
	repository.FindAllTodoRepo
}

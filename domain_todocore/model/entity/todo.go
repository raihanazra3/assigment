package entity

import (
	"assigment/domain_todocore/model/errorenum"
	"assigment/domain_todocore/model/vo"
	"strings"
	"time"
)

type Todo struct {
	ID      vo.TodoID `bson:"_id" json:"id"`
	Created time.Time `bson:"created" json:"created"`
	Updated time.Time `bson:"updated" json:"updated"`
	Message string    `json:"message"`
	Checked bool      `json:"checked"`

	// edit or add new necessary field here ...
}

type TodoCreateRequest struct {
	RandomString string    `json:"-"`
	Now          time.Time `json:"-"`
	Message      string    `json:"message"`

	// edit or add new necessary field for create request here ...

}

func (r TodoCreateRequest) Validate() error {

	// validate the create request here ...

	if strings.TrimSpace(r.Message) == "" {
		return errorenum.MessageMustNotEmpty
	}

	return nil
}

func NewTodo(req TodoCreateRequest) (*Todo, error) {

	err := req.Validate()
	if err != nil {
		return nil, err
	}

	id, err := vo.NewTodoID(req.RandomString, req.Now)
	if err != nil {
		return nil, err
	}

	var obj Todo
	obj.ID = id
	obj.Created = req.Now
	obj.Updated = req.Now
	obj.Message = req.Message
	obj.Checked = false

	// another field input here ...

	return &obj, nil
}

type TodoUpdateRequest struct {
	Now time.Time `json:"-"`

	// add new necessary field to update request here ...

}

func (r TodoUpdateRequest) Validate() error {

	// validate the update request here ...

	return nil
}

func (r *Todo) Update(req TodoUpdateRequest) error {

	err := req.Validate()
	if err != nil {
		return err
	}

	r.Updated = req.Now

	// update field here ...

	return nil
}

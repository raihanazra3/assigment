package vo

import (
	"fmt"
	"time"
)

type TodoID string

func NewTodoID(randomStringID string, now time.Time) (TodoID, error) {
	var obj = TodoID(fmt.Sprintf("OBJ-%s-%s", now.Format("060102"), randomStringID))

	// you may change it as necessary ...

	return obj, nil
}

func (r TodoID) String() string {
	return string(r)
}

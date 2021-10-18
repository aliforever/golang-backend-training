package int64option

import (
	"errors"
	"fmt"
)

type Type struct {
	data *int64
}

// Nothing returns an option with nothing in it.
func Nothing() Type {
	return Type{}
}

// Something returns an option with ‘ value’ in it.
func Something(value int64) Type {
	return Type{data: &value}
}

// Return returns the value it store if there is something in it, else it return an error.
func (t Type) Return() (val int64, err error) {
	if t.data == nil {
		err = errors.New("empty_storage")
		return
	}
	val = *t.data
	return
}

func (t Type) GoString() string {
	if t.data != nil {
		return fmt.Sprintf("int64option.Something(%d)", *t.data)
	}
	return "int64option.Nothing()"
}

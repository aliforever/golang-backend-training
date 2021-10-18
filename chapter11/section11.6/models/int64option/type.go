package int64option

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
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

func (t Type) String() string {
	if t.data != nil {
		return fmt.Sprintf("%d", *t.data)
	}
	return "⧼nothing⧽"
}

func (t Type) MarshalJSON() (b []byte, err error) {
	s := "nothing()"
	if t.data != nil {
		s = fmt.Sprintf("something(%d)", *t.data)
	}
	return json.Marshal(s)
}

func (t *Type) UnmarshalJSON(b []byte) (err error) {
	if t == nil {
		err = errors.New("nil_receiver")
		return
	}
	var optionTypeStr string
	if err = json.Unmarshal(b, &optionTypeStr); err != nil {
		return
	}

	var value = Nothing()
	if optionTypeStr == "nothing()" {
		*t = value
	} else if strings.HasPrefix(optionTypeStr, "something(") && strings.HasSuffix(optionTypeStr, ")") {
		optionTypeStr = optionTypeStr[len("something(") : len(optionTypeStr)-1]
		var intVal int64
		intVal, err = strconv.ParseInt(optionTypeStr, 10, 64)
		if err != nil {
			return
		}
		*t = Something(intVal)
	} else {
		err = errors.New("invalid_int64option")
	}
	return
}

func (t Type) Value() (val driver.Value, err error) {
	var value int64
	value, err = t.Return()
	if err != nil {
		return nil, err
	}
	return value, nil
}

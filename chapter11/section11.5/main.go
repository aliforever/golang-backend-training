package main

import (
	"encoding/json"
	"fmt"

	"github.com/aliforever/golang-backend-training/chapter11/section11.5/models/int64option"
)

func main() {
	var int64OptionType int64option.Type

	var data []byte = []byte(`"nothing()"`)
	err := json.Unmarshal(data, &int64OptionType)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("opt is %s\n", int64OptionType)

	data = []byte(`"something(5)"`)
	err = json.Unmarshal(data, &int64OptionType)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("opt is %s", int64OptionType)
}

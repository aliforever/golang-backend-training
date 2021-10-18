package main

import (
	"encoding/json"
	"fmt"

	"github.com/aliforever/golang-backend-training/chapter11/section11.4/models/int64option"
)

func main() {
	var opt int64option.Type = int64option.Nothing()
	j, err := json.Marshal(opt)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s\n", j)

	opt = int64option.Something(42)
	j, err = json.Marshal(opt)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s\n", j)
}

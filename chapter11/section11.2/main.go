package main

import (
	"fmt"

	"github.com/aliforever/golang-backend-training/chapter11/section11.2/models/int64option"
)

func main() {
	var data int64option.Type
	fmt.Printf("opt = %#v\n", data)

	var opt int64option.Type = int64option.Nothing()
	fmt.Printf("opt = %#v\n", opt)

	opt = int64option.Something(42)
	fmt.Printf("opt = %#v\n", opt)
}

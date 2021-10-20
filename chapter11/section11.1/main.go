package main

import (
	"fmt"

	"github.com/aliforever/golang-backend-training/chapter11/section11.1/lib/int64option"
)

func main() {
	var data int64option.Type
	fmt.Println(data)

	data = int64option.Something(41)
	fmt.Println(data)
}

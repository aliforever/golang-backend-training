package main

import (
	"fmt"

	"github.com/aliforever/golang-backend-training/chapter11/section11.3/models/int64option"
)

func main() {
	var opt int64option.Type = int64option.Nothing()
	fmt.Printf("opt = %s\n", opt)

	opt = int64option.Something(42)
	fmt.Printf("opt = %s", opt)
}

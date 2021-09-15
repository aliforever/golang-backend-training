package main

import (
	"fmt"

	"github.com/aliforever/golang-backend-training/chapter2/chapter2.4/arg"
)

func main() {
	var (
		name string
		shhh bool
	)

	arg.ProcessFlags(&name, &shhh)

	if !shhh {
		fmt.Printf("Hello %s!", name)
	}
}

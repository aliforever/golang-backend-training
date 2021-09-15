package main

import (
	"fmt"

	"github.com/aliforever/golang-backend-training/chapter2/chapter2.5/arg"
)

func main() {
	if !arg.Shhh {
		fmt.Printf("Hello %s!", arg.Name)
	}
}

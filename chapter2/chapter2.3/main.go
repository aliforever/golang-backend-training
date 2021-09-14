package main

import (
	"fmt"
)

func main() {
	var (
		name string
		shhh bool
	)

	processFlags(&name, &shhh)

	if !shhh {
		fmt.Printf("Hello %s!", name)
	}
}

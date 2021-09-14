package main

import (
	"flag"
	"fmt"
)

func main() {
	var (
		name string
		shhh bool
	)

	flag.StringVar(&name, "name", "Joe", "--name=Joe")
	flag.BoolVar(&shhh, "shhh", false, "--shhh")
	flag.Parse()

	if !shhh {
		fmt.Printf("Hello %s!", name)
	}
}

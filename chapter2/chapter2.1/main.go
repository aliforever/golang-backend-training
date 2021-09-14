package main

import (
	"flag"
	"fmt"
)

func main() {
	var name string
	flag.StringVar(&name, "name", "Joe", "--name=Joe")
	flag.Parse()

	fmt.Printf("Hello %s!", name)
}

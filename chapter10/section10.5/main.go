package main

import (
	"fmt"

	"github.com/aliforever/go-cad"
)

func main() {
	xCad, err := cad.ParseCAD("$9.06")
	if err != nil {
		fmt.Println(err)
		return
	}

	yCad, err := cad.ParseCAD("$5.09")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Subtract of y from x would be: %s\n", xCad.Sub(yCad))
}

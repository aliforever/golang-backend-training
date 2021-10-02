package main

import (
	"fmt"

	"github.com/aliforever/go-cad"
)

func main() {
	xCad, err := cad.ParseCAD("$1.06")
	if err != nil {
		fmt.Println(err)
		return
	}

	yCad, err := cad.ParseCAD("$5.09")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(fmt.Sprintf("Sum of x & y would be: %#v", xCad.Add(yCad)))
}

package main

import (
	"fmt"

	"github.com/aliforever/golang-backend-training/chapter11/section11.6/lib/int64option"

	"github.com/aliforever/golang-backend-training/chapter11/section11.6/models/optiontypemodel"
)

func main() {
	var id int

	var data int64option.Type = int64option.Something(21)

	id, err := optiontypemodel.Create("Ali", data)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("inserted id is %d\n", id)

	id, err = optiontypemodel.Create("Ali Error", int64option.Nothing())
	if err != nil {
		fmt.Println(err)
		return
	}
}

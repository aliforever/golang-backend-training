package main

import (
	"fmt"

	"github.com/aliforever/golang-backend-training/chapter11/section11.7/models/optiontypemodel"
)

func main() {
	var id int = 1

	data, err := optiontypemodel.FindById(id)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("opt is %s", data.OptionType)
}

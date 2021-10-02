package main

import (
	"encoding/json"
	"fmt"

	"github.com/aliforever/go-cad"
)

type Result struct {
	Name    string  `json:"name"`
	Balance cad.CAD `json:"balance"`
}

func main() {
	var result Result

	result.Name = "Joe Blow"
	result.Balance = cad.Cents(12345)

	jsonBytes, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(jsonBytes))
}

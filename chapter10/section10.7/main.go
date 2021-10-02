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
	jsonBytes := []byte(`
{
    	"name":    "Jow Blow",
    	"balance": "CAD$123.45"
}
`)

	var r *Result
	err := json.Unmarshal(jsonBytes, &r)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(r.Balance)
}

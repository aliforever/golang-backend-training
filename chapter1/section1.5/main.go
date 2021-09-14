package main

import "fmt"

func main() {
	var (
		name    string = "John"
		weather string = "raining"
		snack   string = "macaroni and cheese"
	)

	fmt.Printf(`Hello %s!

The weather today is %s.

Today's snack will be %s.`, name, weather, snack)
}

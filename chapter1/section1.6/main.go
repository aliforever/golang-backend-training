package main

import "fmt"

func FormLetter(name, weather, snack string) {
	fmt.Printf(`Hello %s!

The weather today is %s.

Today's snack will be %s.`, name, weather, snack)

}

func main() {
	var (
		name    string = "John"
		weather string = "raining"
		snack   string = "macaroni and cheese"
	)
	FormLetter(name, weather, snack)
}

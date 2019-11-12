package main

import "fmt"

func main() {
	fmt.Println("Hello from main")

	var some Some

	some = SomeImpl{value:14}

	fmt.Printf("Some Method returns " + string(some.Method("bob")))
}

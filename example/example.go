package main

import "github.com/hk-32/console"

func main() {
	// Prints the same output as if it was 'fmt.Println'
	console.WriteLine("Hello World", true, 100, vector{420, 69})

	// Use a python like input function
	name := console.Input("Please enter your name: ")
	console.WriteLine("Hello", name)

	// Or read line directly
	console.ReadLine()
}

type vector struct {
	x int64
	y int64
}

// Custom types require a method 'String' that can describe the structure.
func (v vector) String() string {
	return console.Group(v.x, v.y)
}

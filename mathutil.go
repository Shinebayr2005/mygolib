package mygolib

import "fmt"

// Add adds two integers and returns the result
func Add(a, b int) int {
	return a + b
}

// Greet prints a greeting message
func Greet(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

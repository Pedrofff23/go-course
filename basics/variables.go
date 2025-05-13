package basics

import "fmt"

// This type of variable can just be used localy
// middleName := "Caramelo"

var middleName = "Caramelo"

func main() {
	var age int
	var name string
	var isStudent bool
	var height float64
	var weight int64

	count := 10
	lastName := "Doe"

	middleName := "Cane"

	fmt.Println(middleName)
	// Default values
	// Numeric types: 0
	// Boolean: false
	// String: ""
	// Pointer, slices, maps, functions and structs: nil
}

func printName() {
	firstName := "John"
	fmt.Println("First Name:", firstName)
}

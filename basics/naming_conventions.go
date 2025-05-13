package basics

import "fmt"

type Employee struct {
	FistName string
	LastName string
	Age      int
}

func main() {
	// PascalCase
	// Eg. CalculatedArea, UserInfo, NewHttpRequest
	// Structs,interfaces, enums

	// snake_case
	// Eg. used_id, fist_name, http_request

	// UPPERCASE
	// Use case is Constants

	// mixCase
	// Eg. userId, firstName, httpRequest

	const MAX_LENGTH = 100

	employeeID := 1001
	fmt.Println("Employee ID:", employeeID)
}

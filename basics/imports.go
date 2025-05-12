package basics

import (
	"fmt"
	foo "net/http"
)

func main() {
	fmt.Println("Hello, GO student!")

	res, err := foo.Get("https://www.google.com")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer res.Body.Close()

	fmt.Println("Response Status:", res.Status)
}

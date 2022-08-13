package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func message(name *string) string {
	n := "strange"
	if name != nil {
		n = *name
	}
	
	return fmt.Sprintf("Hi %s", n)
}

func main() {
	result := add(2, 3)
	fmt.Println(result)
}
package main

import "fmt"

func main() {
	name := "Izak"
	age := 19
	message := fmt.Sprintf("%s is %d years old", name, age)
	
	fmt.Println(message)
}
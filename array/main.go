package main

import (
	"fmt"
	"sort"
	"strings"
)

type Person struct {
	Name string
	Age int
}

type PersonController []Person

func (pc PersonController) Len() int {
	return len(pc)
}

func (pc PersonController) Swap(i, j int) {
	pc[i], pc[j] = pc[j], pc[i]
}		

func (pc PersonController) Less(i, j int) bool {		
	return pc[i].Age < pc[j].Age
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(numbers)
	
	sub := numbers[2:4]
	fmt.Println(sub)
	
	concatenated := append(numbers, []int{6, 7}...);
	fmt.Println(concatenated)
	
	prepended := append([]int{-2, -1, 0}, concatenated...)
	fmt.Println(prepended)
	
	letters := []string{"a", "b", "c", "d"}
	for index, value := range letters {
		fmt.Println(index, value)
	}
	
	mapped := make([]string, len(letters))
	for index, value := range letters {
		mapped[index] = strings.ToUpper(value)
	}
	
	fmt.Println(mapped)
	
	var filtered []string
	for index, value := range letters {
		if index%2 == 0 {
			filtered = append(filtered, value)
		}
	}
	
	fmt.Println(filtered)
	
	sort.Ints(numbers)
	fmt.Println(numbers)
	
	sort.Sort(sort.Reverse(sort.IntSlice(numbers)))
	fmt.Println(numbers)
	
	collection := []Person {
		{"Sr Izak", 19},
		{"Sr Doguinho", 8},
		{"Sra Gata", 4},
	}
	
	sort.Sort(PersonController(collection))
	fmt.Println(collection)
	
	sort.Sort(sort.Reverse(PersonController(collection)))
	fmt.Println(collection)
}
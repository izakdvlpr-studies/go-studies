package main

func main() {
	// primitives
	var myBool bool = true
	var myInt int = 10
	var myUint uint = 10
	var myFloat32 float32 = 10.5
	var myFloat64 float64 = 10.5
	var myString string = "foo"
	var myByte byte = 10  // alias to uint8
	var myRune rune = 'a' // alias to int32
	
	// composite types
	var myStruct struct{} = struct{}{}
	var myArray []string = []string{}
	var myMap map[string]int = map[string]int{}
	var myFunction func() = func() {}
	var myChannel chan bool = make(chan bool)
	var myInterface interface{} = nil
	var myPointer *int = new(int)
}
package main

func main () {
	var foo string = "foo" // explicit
	var bar = "foo" // type inferred
	baz := "bar" // shorthand
	const qux = "qux" // constant
	
	println(foo)
	println(bar)
	println(baz)
	println(qux)
}
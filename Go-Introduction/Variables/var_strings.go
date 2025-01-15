package main

import "fmt"

var z string = "Ambiguous variable"

func main() {
	var a string = "Hello GoLang World!"
	fmt.Println(a)

	var x string
	x = "Hello GoLang World!"
	fmt.Println(x)
	x = "Hello Flutter world!"
	fmt.Println(x)
	x = x + "I love programming"
	fmt.Println(x)
	x += "I love writing code"
	fmt.Println(x)





	var m string = "GoLang"
	var n string = "Flutter"
	var p string = "GoLang"
	fmt.Println(m == n)
	fmt.Println(m == p)





	// Type Inference

	b := "Hello GoLang World!"
	fmt.Println(b)

	c := "2025"
	fmt.Println(c)



	// Naming variables
	studentName := "Job Okullo"
	fmt.Println(studentName)

	// Scope
	// Decalring variables outside the main function

	fmt.Println(z)


	// Constants
	const pi float64 = 3.14159265359
	fmt.Println(pi)


	// Defining multiple variables
	var(
		l = "Length"
		w = "Width"
		h = "Height"
	)
	fmt.Println(l, w, h)
}

func foo() {
	fmt.Println(z)
}
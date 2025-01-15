package main

import "fmt"

func main() {
	// fmt.Println("Enter your name:")
	// var name string
	// fmt.Scanf("%s", &name)

	fmt.Println("Enter your body temperature: ")
	var tc float64
	fmt.Scanf("%f", &tc)

	tf := (tc * 9/5) + 32
	fmt.Println("Body temperature in Fahrenheit is", tf)
}
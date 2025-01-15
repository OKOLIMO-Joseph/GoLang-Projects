package main

import "fmt"

func main() {
	fmt.Println("Enter the radius of the circle:  ")
	var radius float64
	fmt.Scanf("%f", &radius)

	area := 3.14159265359 * radius * radius
	fmt.Println("The area of the circle is: ", area)
}
package main

import "fmt"


func concat(a string, b string) string {
	return a + b
}

func main() {
	fmt.Println("UCU Athletes Biography")
	fmt.Println(concat("Joseph", "BSIT"))
	fmt.Println(concat("Keith", "BBA"))
	fmt.Println(concat("Tush", "BSECE"))
	fmt.Println(concat("Charity", "HRM"))

}
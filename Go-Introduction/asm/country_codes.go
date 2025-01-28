package main

import "fmt"

func main() {
	country := make(map[int]string)
	country[1] = "USA"
	country[256] = "Uganda"
	country[255] = "Tanzania"
	country[254] = "Kenya"
	country[234] = "Nigeria"

	fmt.Println(country[256])
}
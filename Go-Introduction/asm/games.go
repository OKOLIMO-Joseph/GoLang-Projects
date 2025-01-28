package main

import "fmt"

func main(){
	sports := make(map[string]string)
	sports["1"] = "Football"
	sports["2"] = "Basketball"
	sports["3"] = "Volleyball"
	sports["4"] = "Tennis"
	sports["5"] = "Cricket"

	fmt.Println(sports["4"])
}
package main

import "fmt"

func main(){
	for a := 1; a <= 100; a++{
		if a % 3 == 0{
			fmt.Println(a, "Fizz!")
		} else if a % 5 == 0{
			fmt.Println(a, "Buzz")
		} else if a % 3 == 0 && a % 5 == 0{
			fmt.Println(a, "FizzBuzz!")
		} else {
			fmt.Println(a, "is not a multiple of 3 and 5")
		}
	}
}
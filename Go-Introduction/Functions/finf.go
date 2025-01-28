package main

import "fmt"

func main() {
	fmt.Println(fn1())

}

func fn1() int{
return fn2()
}

func fn2() int{
return 15
}
package main

import "fmt"

func main() {
	fmt.Println("My Multi-Algebric program")
	bs:= 6
	h :=4
	l := 5
	w := 8
	a := 29
	b := 10
	ans := addNums(a, b)
	fmt.Println(a, "+", b, "=", ans)
	fmt.Println(a, "-", b, "=", subNums(a, b))
	fmt.Println(a, "*", b, "=", mulNums(a, b))
	fmt.Println(a, "/", b, "=", divNums(a, b))
	fmt.Println("Area of Rectangle = ", areaRectangle(l, w))
	fmt.Println("Area of Triangle = ", areaTriangle(bs, h))
}

func addNums(a int, b int) int {
	var  ans int = a + b
	return ans
}

func subNums(a int, b int) int {
	var  ans int = a - b
	return ans
}

func mulNums(a int, b int) int {
	var  ans int = a * b
	return ans
}

func divNums(a int, b int) int {
	var  ans int = a / b
	return ans
}

func areaRectangle(l int, w int) int{
	var ans int = l * w
	return ans
}

func areaTriangle(bs int, h int) int{
	var ans int = (bs * h) / 2
	return ans
}
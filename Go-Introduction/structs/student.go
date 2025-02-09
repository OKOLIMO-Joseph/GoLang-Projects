package main

import "fmt"

type Student struct {
	name      string
	access_no string
	course    string
	age       int
	cgpa      float64
}

func main() {
	var student1 Student
	student1.name = "Joseph"
	student1.access_no = "B22562"
	student1.course = "Computer Science"
	student1.age = 20
	student1.cgpa = 4.8

	fmt.Println("Student Name: ", student1.name)
	fmt.Println("Access Number: ", student1.access_no)
	fmt.Println("Course: ", student1.course)
	fmt.Println("Age: ", student1.age)
	fmt.Println("CGPA: ", student1.cgpa)
}
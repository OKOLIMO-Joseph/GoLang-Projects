package main

import "fmt"

func main(){
	fmt.Println("Grade Calculator")
	var z int
	fmt.Println("Enter the final mark: ")
	fmt.Scanf("%f", &z)
	// for z := 0; z <=100; z++{
		if z < 50{
			fmt.Println("You haver failed, adviced to retake!")
		} else if 51 < z && z < 60{
			fmt.Println(z, "Grade: D")
		} else if 61 < z && z < 70{
			fmt.Println(z, "Grade: C")
		} else if 71 < z && z < 80{
			fmt.Println(z, "Grade: B")
		} else if 81 < z && z < 100{
			fmt.Println(z, "Grade: A")
		}
	// } 
}
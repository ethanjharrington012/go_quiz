package main

import (
	"fmt"
)

// add_nums takes two integer parameters, calculates their product, and evaluates
// if the result is less than 100. It then prints either the result or a message accordingly.





func add_nums(x, y int) {
	result := x * y
	if result < 100 {
		fmt.Println("Your result is less then 100: ", result)
	} else{
		fmt.Print(result, "\n")
	}
	
	

}
// even_nums takes an integer parameter and determines whether it's even or odd,
// printing the corresponding message.

func even_nums(x int) {
	if x%2 == 0 {
		fmt.Println("You are an even!")
	} else {
		fmt.Println("You are odd")
	}
}



func main() {
	// Example usage of the functions

	add_nums(40, 5)
	even_nums(10)


}
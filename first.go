package main
// You always have to have package main

import (
	
	"bufio"
    "fmt"
    "os"	
)
//  this provides functions for fomating and printing texts

func utilityFunction() {
	// you can make functions outside of main.
	// but when you print the package. it has to be within main
	// so you have to call utilityFunction() in main()
	fmt.Println("Inside new function")
}

func add_num(x, y int) int {
	result := x + y
	return result
}

// constants bellow
// when a constant is declared you can not change it.

const (
	A int = 1
	B = 3.14
	C = "HI!"
)


func main() {

	// fmt.Print("Enter something: ")

    // scanner := bufio.NewScanner(os.Stdin)
    // scanner.Scan()
    // input := scanner.Text()
	

    fmt.Println("You entered:", input)
	fmt.Println(A, "\n")
	fmt.Println(B)
	fmt.Println(C)

	var new = 2
	fmt.Println("\nHello world")
	fmt.Println(new)

	sum := add_num(5,6)
	fmt.Println(sum)
	fmt.Println("\nBellow is: JOHN")
	var new_num string = "John"
	utilityFunction()

	var i,j string = "HEllo", "EThan"
	fmt.Println(i,"\n", j)
	fmt.Println(i, " ",j)

	// Print inserts a space between the objects if neither are strings
	var g,h = 10,20

	fmt.Println(g,h)


	fmt.Println(new_num)

	x := 4
	fmt.Println(x)

	var myNum int = 90
	var st string = "Hello"
	var bo bool = true

	fmt.Println(myNum)
	fmt.Println(st)
	fmt.Println(bo)

	var arr1 = [...]int{1,2,3}
	arr2 := [...]int{4,5,6,7,8}

	fmt.Println(arr1)
	fmt.Println(arr2)
}



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
	
	
	// var array_name = [length]datatype{values} 
	// or
	// var array_name = [...]datatype{values}

	// array_name := [length]datatyle{values}
	// array_name := [...]datatype{values}

	var first_arr = [3]int{1, 4, 8}

	fmt.Println(first_arr)
	fmt.Println(first_arr[0])

	second_arr := [...]string{"first string", "second", "third", "forth"}
	fmt.Println(second_arr)
	fmt.Println(second_arr[1])

	// changing elements in an array
	prices := [3]int{10,20,30}
	prices[2] = 50
	fmt.Println(prices)
	for i:=0; i < len(prices); i++ {
		fmt.Println(prices[i])
	}
	for idx, val := range prices {
		fmt.Println(idx, val)
	}
	// struct is kinda like a dictionary in Python. 
	// while an array is used to store mutlipul of the same data type
	// a struct is used to store multipul data types into a variable

	type Person struct {
		name string
		age int
		job string
		salary int
	}
	var pers1 Person


	pers1.name = "game"
	pers1.age = 20
	pers1.job = "teacher"
	pers1.salary = 100000

	fmt.Println("Person1 name:", pers1.name)
	fmt.Println("Person1 Salary:", pers1.salary)

	// lastly a map is extactly like a dictionary

	// var a = map[KeyType]ValueType{key1:value1, key2:value2, ...}
	// b := map[KeyType]ValueType{key1: value1, key2: value2}

	var a_map = map[string]string{"Brand": "Ford", "model": "Mustage", "year": "1965"}
	b_map := map[string]int{"Olso": 1, "Bergan": 2, "Trondheim": 3}
	fmt.Println(a_map["Brand"])
	fmt.Println(b_map["Bergan"])
}
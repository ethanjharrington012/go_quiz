package main

import (
	"fmt"

)

// first I need a couple inputs
// second I need some conditionals

func main() {
	var i int

	fmt.Print("Type a number: ")
	fmt.Scan(&i)
	fmt.Println("Your number is:",i)

	fmt.Println("Great you know how to type! \nLets do a quiz!\n")

	var first int
	fmt.Print("\nWho lives in a Pinapple under the sea: ")
	fmt.Println("\n1: Spongebob\n2: Spongenob\n3: patrich")
	fmt.Scan(&first)
	
	if first == 1 {
		fmt.Println("\nGreat work!")
		
	} else {
		fmt.Println("\nWrong. the currect answer was spongebob.\nYour answer was:", first)
	}

	var second_q int
	fmt.Print("what is the biggest anilmal: ")
	fmt.Println("\n1: elephant\n2: blue whale\n3: giant octopus")
	fmt.Scan(&second_q)
	if second_q == 2 {
		fmt.Println("Great work!")
		
	} else {
		fmt.Println("Wrong! The correct answer is blue whale. not")
	}





}


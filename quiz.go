package main

import (
	"fmt"

)

// first I need a couple inputs
// second I need some conditionals

func main() {

	score := 0
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
		score += 1
		
	} else {
		fmt.Println("\nWrong. the currect answer was spongebob.\nYour answer was:", first)
	}

	var second_q int
	fmt.Print("what is the biggest anilmal: ")
	fmt.Println("\n1: elephant\n2: blue whale\n3: giant octopus")
	fmt.Scan(&second_q)
	if second_q == 2 {
		fmt.Println("Great work!")
		score += 1
		
	} else {
		fmt.Println("Wrong! The correct answer is blue whale. not")
	}

	var third int
	fmt.Println("What company does Steve Jobs own?")
	fmt.Println("\n1: Facebook\n2: Google\n3: Apple")
	fmt.Scan(&third)

	if third == 3 {
		fmt.Println("Wow Your a MASTER!")
		score += 1
		fmt.Println("Your score is:", score)
		

	} else {
		fmt.Println("DAMN you suck!")
		fmt.Println("Your score is:", score)
	}




}

// Next step I would like to add a scoring system, timer, difficulty level, user profiles, categories
// Graphics and UI, high score, customizable questions

// 1st the scoring system is going to look like some sort of data base or struct to store the points.

// for a timer that is going to happen when a new question arrises. 


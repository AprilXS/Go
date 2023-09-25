package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to my quiz!")
	fmt.Println("What is your name?")

	var name string
	fmt.Scan(&name)
	fmt.Printf("Welcome, %v, to my quiz!\n", name)

	var age uint
	fmt.Println("How old are you?")
	fmt.Scan(&age)

	if age < 18 {
		fmt.Println("You are not old enough to take this quiz.")
		return
	}else{
	fmt.Println("You are old enough to take this quiz.")
	}

	score := 0

	fmt.Println("what is the best game? Minecraft 2012 or Starfield 2023")
	var answer string
	var answer2 string
	fmt.Scan(&answer, &answer2) //you need to make two variables for two answers, because Scan only scans one variable at a time to end space

	if answer +" "+ answer2 == "Starfield 2023" ||  answer +" "+ answer2 == "starfield 2023" {
		fmt.Println("Correct!")
		score++
	}else{
		fmt.Println("Incorrect!")
	}

	fmt.Println("how many years where starfield in development?")
	var years uint
	fmt.Scan(&years)

	if years == 25 {
		fmt.Println("Correct!")
		score++
	}else{
		fmt.Println("Incorrect!")
	}

	fmt.Println("your score is", score, "out of 2")
	percent := (float64(score) / float64(2)) * 100
	fmt.Printf("You got %v %% correct!", percent)

	fmt.Println("\nEnd of program.")
}

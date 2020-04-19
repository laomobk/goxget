package main

import "fmt"

func askYN(question string) (yes bool) {
	var answer string
	fmt.Printf(question)
	fmt.Scanln(&answer)

	return answer == "Y" || answer == "y" || answer == "yes"
}

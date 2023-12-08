package main

import (
	"fmt"
	"os"
)

func main() {
	files, _ := os.ReadDir(".")
	for _, val := range files {
		fmt.Println(val)
	}
	fmt.Println("what program would you like to run?")
	var input string
	fmt.Scan(&input)

	switch input {
	case "W101":
		W101()
	case "W102":
		W102()
	case "W103":
		W103()
	case "W106":
		W106()
	default:
		fmt.Println(input, " isn't available")
	}
}

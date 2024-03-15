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
	case "W108":
		W108()
	case "W109":
		W109()
	case "W110":
		W110()
	case "W112":
		W112()
	case "W113":
		W113()
	case "W116":
		W116()
	case "W121":
		W121()
	default:
		fmt.Println(input, " isn't available")
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

package main

import (
	"fmt"
)

func W102() {
	var Menu = [...]string{"Option 1: ", "Option 2:", "Option 3:"}
	for _, val := range Menu {
		fmt.Println(val)
	}

	var input int
	fmt.Scan(&input)
	start := "you pressed "
	switch input {
	case 1:
		fmt.Println(start, "1")
	case 2:
		fmt.Println(start, "2")
	case 3:
		fmt.Println(start, "3")
	}
}

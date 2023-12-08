package main

import (
	"fmt"
)

func W101() {
	fmt.Println("choose a number between one and four")
	var Num string
	fmt.Scan(&Num)

	if Num == "1" {
		fmt.Println("You won a TV")
	} else if Num == "2" {
		fmt.Println("You have won a sock")
	} else if Num == "3" {
		fmt.Println("You have won a cuddly toy")
	} else {
		fmt.Println("you win nothing")
	}
}

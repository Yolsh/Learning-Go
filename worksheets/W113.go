package main

import (
	"fmt"
	"os"
)

func W113() {
	var Ans string
	ConsoleLogger("Say Something")
	fmt.Scan(&Ans)
	ConsoleLogger(Ans)
}

func ConsoleLogger(input string) {
	f, err := os.OpenFile("./Readable_Files/W113.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	check(err)
	if _, err := f.WriteString(input + "\n"); err != nil {
		panic(err)
	}
	f.Close()
	fmt.Println(input)
}

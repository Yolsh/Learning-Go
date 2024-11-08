package AoC

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	AoC2023 "github.com/Yolsh/Learning-Go/AoC/2023-Practice"
	AoC2024 "github.com/Yolsh/Learning-Go/Aoc/2024"
)

func Run() {
	files, err := os.ReadDir("./AoC")
	check(err)
	packages := []string{}
	fmt.Println("What Package would you like to run?")
	var i int16
	for _, val := range files {
		if !strings.ContainsRune(val.Name(), '.') {
			i++
			fmt.Printf("%v: %v\n", i, val.Name())
			packages = append(packages, val.Name())
		}
	}
	var input string
	fmt.Scan(&input)

	val, err := strconv.ParseInt(input, 10, 64)
	check(err)
	switch val {
	case int64(slices.Index(packages, "2023-Practice")) + 1:
		AoC2023.Run()
	case int64(slices.Index(packages, "2024")) + 1:
		AoC2024.Run()
	default:
		fmt.Println("That option doesn't exist")
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

package AoC2023

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Run() {
	files, err := os.ReadDir("./AoC")
	check(err)
	packages := []string{}
	var i int16
	fmt.Println("What Program would you like to run?")
	for _, val := range files {
		if strings.Contains(val.Name(), ".go") && val.Name() != "run.go" {
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
	case int64(slices.Index(packages, "Day1.go")) + 1:
		Day1()
	case int64(slices.Index(packages, "Day2.go")) + 1:
		Day2()
	case int64(slices.Index(packages, "Day3.go")) + 1:
		Day3()
	case int64(slices.Index(packages, "Day5.go")) + 1:
		Day5()
	case int64(slices.Index(packages, "Day7.go")) + 1:
		Day7()
	case int64(slices.Index(packages, "Day8.go")) + 1:
		Day8()
	default:
		fmt.Println(input, " isn't available")
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

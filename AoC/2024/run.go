package AoC2024

import (
	"fmt"
	"os"

	//"slices"
	"strconv"
	"strings"
)

func Run() {
	files, err := os.ReadDir("./AoC/2023-Practice")
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
	//case int64(slices.Index(packages, "Day1.go")) + 1:
	//Day1()
	default:
		fmt.Println(input, " isn't available")
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

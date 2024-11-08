package worksheets

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Run() {
	files, err := os.ReadDir("./worksheets")
	check(err)
	packages := []string{}
	var i int16
	fmt.Println("what program would you like to run?")
	for _, val := range files {
		if !strings.Contains(val.Name(), ".go") && val.Name() != "run.go" {
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
	case int64(slices.Index(packages, "W101.go")) + 1:
		W101()
	case int64(slices.Index(packages, "W102.go")) + 1:
		W102()
	case int64(slices.Index(packages, "W103.go")) + 1:
		W103()
	case int64(slices.Index(packages, "W106.go")) + 1:
		W106()
	case int64(slices.Index(packages, "W108.go")) + 1:
		W108()
	case int64(slices.Index(packages, "W109.go")) + 1:
		W109()
	case int64(slices.Index(packages, "W110.go")) + 1:
		W110()
	case int64(slices.Index(packages, "W112.go")) + 1:
		W112()
	case int64(slices.Index(packages, "W113.go")) + 1:
		W113()
	case int64(slices.Index(packages, "W116.go")) + 1:
		W116()
	case int64(slices.Index(packages, "W121.go")) + 1:
		W121()
	case int64(slices.Index(packages, "A122.go")) + 1:
		A122()
	default:
		fmt.Println(input, " isn't available")
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

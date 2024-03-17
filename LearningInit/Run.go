package LearningInit

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Run() {
	var i int16
	var Ans string
	var fileNames []string
	files, err := os.ReadDir("./LearningInit")
	check(err)
	for _, val := range files {
		if strings.Contains(val.Name(), ".go") && val.Name() != "Run.go" {
			i++
			fmt.Printf("%v: %v\n", i, val.Name())
			fileNames = append(fileNames, val.Name())
		}
	}
	fmt.Println("What would you like to run?")
	fmt.Scan(&Ans)
	val, err := strconv.ParseInt(Ans, 10, 64)
	check(err)
	switch val {
	case int64(slices.Index(fileNames, "Fizz Buzz.go")):
		FizzBuzz()
	case int64(slices.Index(fileNames, "Hello.go")):
		Hello()
	case int64(slices.Index(fileNames, "regex1.go")):
		Regex1()
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

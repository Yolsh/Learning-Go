package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/Yolsh/Learning-Go/LearningInit"
	"github.com/Yolsh/Learning-Go/worksheets"
)

func main() {
	var Ans string
	var i int8
	packages := []string{}
	files, err := os.ReadDir(".")
	check(err)
	fmt.Println("What Package would you like to run?")
	for _, val := range files {
		var fileName string = val.Name()
		if !strings.ContainsRune(fileName, '.') {
			i++
			fmt.Printf("%v: %v\n", i, fileName)
			packages = append(packages, fileName)
		}
	}
	fmt.Scan(&Ans)
	val, err := strconv.ParseInt(Ans, 10, 64)
	check(err)
	switch val {
	case int64(slices.Index(packages, "LearningInit")):
		LearningInit.Run()
	case int64(slices.Index(packages, "worksheets")):
		worksheets.Run()
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

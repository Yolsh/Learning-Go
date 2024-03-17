package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/Yolsh/Learning-Go/LearningInit"
	"github.com/Yolsh/Learning-Go/worksheets"
	"github.com/Yolsh/Learning-Go/simpleServer"
	"github.com/Yolsh/Learning-Go/concurrencySimple"
)

func main() {
	var Ans string
	var i int8
	packages := []string{}
	files, err := os.ReadDir(".")
	check(err)
	fmt.Println("What Package would you like to run?")
	for _, val := range files {
		if !strings.ContainsRune(val.Name(), '.') {
			i++
			fmt.Printf("%v: %v\n", i, val.Name())
			packages = append(packages, val.Name())
		}
	}
	fmt.Scan(&Ans)
	val, err := strconv.ParseInt(Ans, 10, 64)
	check(err)
	switch val {
		case int64(slices.Index(packages, "LearningInit")) + 1:
			LearningInit.Run()
		case int64(slices.Index(packages, "worksheets")) + 1:
			worksheets.Run()
		case int64(slices.Index(packages, "simpleServer")) + 1:
			simpleServer.ServerStart()
		case int64(slices.Index(packages, "concurrencySimple")) + 1:
			concurrencySimple.Start()
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

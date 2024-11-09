package AoC2023

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func Day1() {
	var lines []string = readFile("Day1.txt")
	var total int
	for i, line := range lines {
		expression, err := regexp.Compile(`oneight|twone|threeight|fiveight|nineight|eightwo|eighthree|one|two|three|four|five|six|seven|eight|nine|\d`)
		check(err)
		var finds []int
		matches := expression.FindAllString(line, -1)
		for _, match := range matches {
			switch match {
			case "one":
				finds = append(finds, 1)
			case "two":
				finds = append(finds, 2)
			case "three":
				finds = append(finds, 3)
			case "four":
				finds = append(finds, 4)
			case "five":
				finds = append(finds, 5)
			case "six":
				finds = append(finds, 6)
			case "seven":
				finds = append(finds, 7)
			case "eight":
				finds = append(finds, 8)
			case "nine":
				finds = append(finds, 9)
			case "oneight":
				finds = append(finds, 1)
				finds = append(finds, 8)
			case "twone":
				finds = append(finds, 2)
				finds = append(finds, 1)
			case "threeight":
				finds = append(finds, 3)
				finds = append(finds, 8)
			case "fiveight":
				finds = append(finds, 5)
				finds = append(finds, 8)
			case "nineight":
				finds = append(finds, 9)
				finds = append(finds, 8)
			case "eightwo":
				finds = append(finds, 8)
				finds = append(finds, 2)
			case "eighthree":
				finds = append(finds, 8)
				finds = append(finds, 3)
			default:
				val, err := strconv.Atoi(match)
				check(err)
				finds = append(finds, val)
			}
		}
		fmt.Printf("%v: %v, %v\n", i, finds[0], finds[len(finds)-1])
		total += finds[0]*10 + finds[len(finds)-1]
	}
	fmt.Printf("%v", total)
}

func readFile(path string) []string {
	dat, err := os.ReadFile("./AoC/ReadableFiles/" + path)
	ExtrapDat := string(dat)
	check(err)
	lines := make([]string, 0)
	var line string
	for _, val := range ExtrapDat {
		if val == '\n' {
			lines = append(lines, line)
			line = ""
		} else {
			line += string(val)
		}
	}
	return lines
}

package AoC2023

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func Day2() {
	var lines []string = readFile("Day2.txt")
	var total int
	for _, line := range lines {
		splits := strings.Split(line, " ")
		splits = splits[2:]
		var maxVals [3]int
		for j, split := range splits {
			if unicode.IsDigit(rune(split[0])) {
				num, err := strconv.Atoi(split)
				check(err)
				if splits[j+1][0] == 'r' && num > maxVals[0] {
					maxVals[0] = num
				} else if splits[j+1][0] == 'g' && num > maxVals[1] {
					maxVals[1] = num
				} else if splits[j+1][0] == 'b' && num > maxVals[2] {
					maxVals[2] = num
				}
			}
		}
		total += maxVals[0] * maxVals[1] * maxVals[2]
	}
	fmt.Printf("%v", total)
}

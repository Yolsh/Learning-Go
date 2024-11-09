package AoC2023

import (
	"fmt"
	"regexp"
	"strconv"
	"unicode"
)

func Day3() {
	var lines []string = readFile("Day3.txt")
	specialRe, err := regexp.Compile(`\*`)
	check(err)
	var num int
	for i, line := range lines {
		for j, char := range line {
			isSpecial := specialRe.MatchString(string(char))
			check(err)
			if isSpecial {
				var Str []string = make([]string, 0)
				if i > 0 {
					if unicode.IsDigit(rune(lines[i-1][j])) {
						Str = append(Str, recLookLeft(lines[i-1], j-1, "")+string(lines[i-1][j])+recLookRight(lines[i-1], j+1, ""))
					} else {
						if j > 0 && unicode.IsDigit(rune(lines[i-1][j-1])) {
							Str = append(Str, recLookLeft(lines[i-1], j-1, ""))
						}
						if j < len(line) && unicode.IsDigit(rune(lines[i-1][j+1])) {
							Str = append(Str, recLookRight(lines[i-1], j+1, ""))
						}
					}
				}

				if j > 0 && unicode.IsDigit(rune(lines[i][j-1])) {
					Str = append(Str, recLookLeft(lines[i], j-1, ""))
				}
				if j < len(line) && unicode.IsDigit(rune(lines[i][j+1])) {
					Str = append(Str, recLookRight(lines[i], j+1, ""))
				}

				if i != len(lines) {
					if unicode.IsDigit(rune(lines[i+1][j])) {
						Str = append(Str, recLookLeft(lines[i+1], j-1, "")+string(lines[i+1][j])+recLookRight(lines[i+1], j+1, ""))
					} else {
						if j > 0 && unicode.IsDigit(rune(lines[i+1][j-1])) {
							Str = append(Str, recLookLeft(lines[i+1], j-1, ""))
						}
						if j < len(line) && unicode.IsDigit(rune(lines[i+1][j+1])) {
							Str = append(Str, recLookRight(lines[i+1], j+1, ""))
						}
					}
				}
				fmt.Printf("%v: %v", i, string(char))
				if len(Str) == 2 {
					c1, err := strconv.Atoi(Str[0])
					check(err)
					c2, err := strconv.Atoi(Str[1])
					check(err)
					num += c1 * c2
				}
				fmt.Println()
			}
		}
	}
	fmt.Println(num)
}

func recLookLeft(line string, j int, s string) string {
	if j < 0 || line[j] == '.' || line[j] == '\r' {
		return s
	}
	s = recLookLeft(line, j-1, s)
	s += string(line[j])
	return s
}

func recLookRight(line string, j int, s string) string {
	if j >= len(line) || line[j] == '.' || line[j] == '\r' {
		return s
	}
	s += string(line[j])
	s = recLookRight(line, j+1, s)
	return s
}

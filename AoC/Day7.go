package AoC2023

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type play struct {
	hand     string
	bid      int
	handType int
}

func Day7() {
	lines := readFile("Day7.txt")
	allHands := make([]play, 0)
	for _, line := range lines {
		SplitLine := strings.Split(line, " ")
		bid, err := strconv.Atoi(strings.Replace(SplitLine[1], "\r", "", -1))
		check(err)
		if checkFiOC(SplitLine[0]) {
			allHands = append(allHands, play{SplitLine[0], bid, 7})
		} else if checkFoOC(SplitLine[0]) {
			allHands = append(allHands, play{SplitLine[0], bid, 6})
		} else if checkFH(SplitLine[0]) {
			allHands = append(allHands, play{SplitLine[0], bid, 5})
		} else if checkTOC(SplitLine[0]) {
			allHands = append(allHands, play{SplitLine[0], bid, 4})
		} else if checkTP(SplitLine[0]) {
			allHands = append(allHands, play{SplitLine[0], bid, 3})
		} else if checkOP(SplitLine[0]) {
			allHands = append(allHands, play{SplitLine[0], bid, 2})
		} else {
			allHands = append(allHands, play{SplitLine[0], bid, 1})
		}
	}
	allHands = bubbleSortPlays(allHands)

	var total int
	for i, play := range allHands {
		fmt.Printf("%v %v %v\n", play.hand, play.bid, play.handType)
		total += play.bid * (i + 1)
	}
	fmt.Printf("total bid value: %v", total)
}

func bubbleSortPlays(allPlays []play) []play {
	for i := 0; i < len(allPlays)-1; i++ {
		for j := 0; j < len(allPlays)-i-1; j++ {
			if allPlays[j].handType > allPlays[j+1].handType {
				allPlays[j], allPlays[j+1] = allPlays[j+1], allPlays[j]
			} else if allPlays[j].handType == allPlays[j+1].handType && choose(allPlays[j].hand, allPlays[j+1].hand) {
				allPlays[j], allPlays[j+1] = allPlays[j+1], allPlays[j]
			}
		}
	}
	return allPlays
}

func checkFiOC(hand string) bool {
	return len(group(strings.ReplaceAll(hand, "J", ""))) == 1 || hand == "JJJJJ"
}

func checkFoOC(hand string) bool {
	corrected := strings.ReplaceAll(hand, "J", "")
	grouped := group(corrected)
	var c1 int
	var c2 int
	if len(grouped) == 2 {
		for _, char := range corrected {
			if char == grouped[0] {
				c1++
			} else if char == grouped[1] {
				c2++
			}
		}
	}
	return c1 == 4-strings.Count(hand, "J") || c2 == 4-strings.Count(hand, "J")
}

func checkTOC(hand string) bool {
	corrected := strings.ReplaceAll(hand, "J", "")
	grouped := group(corrected)
	var c1 int
	var c2 int
	var c3 int
	if len(grouped) == 3 {
		for _, char := range corrected {
			if char == grouped[0] {
				c1++
			} else if char == grouped[1] {
				c2++
			} else if char == grouped[2] {
				c3++
			}
		}
	}
	return c1 == 3-strings.Count(hand, "J") || c2 == 3-strings.Count(hand, "J") || c3 == 3-strings.Count(hand, "J") //252555439
}

func group(hand string) []rune {
	sepLabels := make([]rune, 0)
	for _, char := range hand {
		if !slices.Contains(sepLabels, char) {
			sepLabels = append(sepLabels, char)
		}
	}
	return sepLabels
}

func checkFH(hand string) bool {
	return len(group(strings.ReplaceAll(hand, "J", ""))) == 2
}

func checkTP(hand string) bool {
	return len(group(strings.ReplaceAll(hand, "J", ""))) == 3
}

func checkOP(hand string) bool {
	return len(group(strings.ReplaceAll(hand, "J", ""))) == 4
}

func choose(hand1 string, hand2 string) bool {
	for i, char := range hand1 {
		if convInt(char) == convInt(rune(hand2[i])) {
			continue
		} else if convInt(char) > convInt(rune(hand2[i])) {
			return true
		} else if convInt(char) < convInt(rune(hand2[i])) {
			return false
		}
	}
	return false
}

func convInt(char rune) int {
	num, err := strconv.Atoi(string(char))
	if err != nil {
		switch char {
		case 'T':
			num = 10
		case 'J':
			num = 1
		case 'Q':
			num = 12
		case 'K':
			num = 13
		case 'A':
			num = 14
		default:
			panic(err)
		}
	}
	return num
}

package AoC2023

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type mapping struct {
	destination uint64
	source      uint64
	length      uint64
}

func Day5() {
	lines := readFile("Day5.txt")
	stringSeeds := strings.Split(lines[0], " ")[1:]
	var Seeds []uint64
	for _, vals := range stringSeeds {
		value, err := strconv.ParseUint(strings.Replace(vals, "\r", "", -1), 10, 64)
		check(err)
		Seeds = append(Seeds, value)
	}
	allMaps := make([][]mapping, 0)
	var CurMapset []mapping = make([]mapping, 0)
	for _, line := range lines {
		if unicode.IsDigit(rune(line[0])) {
			vals := strings.Split(line, " ")
			destination, err := strconv.ParseUint(strings.Replace(vals[0], "\r", "", -1), 10, 64)
			check(err)
			source, err := strconv.ParseUint(strings.Replace(vals[1], "\r", "", -1), 10, 64)
			check(err)
			length, err := strconv.ParseUint(strings.Replace(vals[2], "\r", "", -1), 10, 64)
			check(err)
			CurMapset = append(CurMapset, mapping{destination, source, length})
		} else {
			if line[0] != '\n' {
				allMaps = append(allMaps, CurMapset)
				CurMapset = make([]mapping, 0)
			}
		}
	}
	fmt.Println(len(allMaps))
	var bestSeed uint64 = ^uint64(0)
	for i := 0; i < len(Seeds); i += 2 {
		fmt.Printf("%v: working on %v up to %v\n", (i+1)/2, Seeds[i], Seeds[i]+Seeds[i+1])
		for j := Seeds[i]; j < Seeds[i]+Seeds[i+1]; j++ {
			var curSeed uint64 = j
			for _, mapset := range allMaps {
				for _, mapping := range mapset {
					if curSeed-mapping.source < mapping.length && curSeed-mapping.source >= 0 {
						curSeed = mapping.destination + (curSeed - mapping.source)
						break
					}
				}
			}
			if curSeed < bestSeed {
				bestSeed = curSeed
			}
		}
	}
	fmt.Printf("best seed: %v", bestSeed)
}

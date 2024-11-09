package AoC2023

import (
	"fmt"
	"strings"
)

type node struct {
	code  string
	left  string
	right string
}

type point struct {
	Node            string
	instructPointer int
}

func Day8() {
	lines := readFile("Day8.txt")
	instruct := lines[0]
	instruct = strings.ReplaceAll(instruct, "\r", "")
	nodes := make([]node, 0)
	for _, line := range lines[2:] {
		line = strings.ReplaceAll(line, "=", "")
		line = strings.ReplaceAll(line, ",", "")
		line = strings.ReplaceAll(line, "(", "")
		line = strings.ReplaceAll(line, ")", "")
		line = strings.ReplaceAll(line, "\r", "")
		set := strings.Split(line, " ")
		nodes = append(nodes, node{set[0], set[2], set[3]})
	}
	startNodes := make([]point, 0)
	for _, node := range nodes {
		if node.code[2] == 'A' {
			startNodes = append(startNodes, point{node.code, 0})
		}
	}
	var count int
	for corrected(startNodes) {
		for i := range startNodes {
			//fmt.Printf("%v, %v ->", startNodes[i].Node, startNodes[i].instructPointer)
			startNodes[i].Node, startNodes[i].instructPointer = reqStep(startNodes[i].Node, nodes, instruct, startNodes[i].instructPointer)
			//fmt.Printf(" %v, %v\n", startNodes[i].Node, startNodes[i].instructPointer)
		}
		count++
	}
	fmt.Printf("steps: %v\n", count)
}

func reqStep(startNode string, nodes []node, instruct string, startPointer int) (string, int) {
	var cur string = startNode
	var instructPointer int = startPointer
	for _, node := range nodes {
		if node.code == cur {
			if rune(instruct[instructPointer]) == 'L' {
				//fmt.Printf("%v -> %v\n", cur, node.left)
				cur = node.left
			} else {
				//fmt.Printf("%v -> %v\n", cur, node.right)
				cur = node.right
			}
			instructPointer++
			if instructPointer == len(instruct) {
				instructPointer = 0
			}
			break
		}
	}
	return cur, instructPointer
}

func corrected(curNodes []point) bool {
	for _, point := range curNodes {
		if point.Node[2] != 'Z' {
			return true
		} else {
			fmt.Println(point.Node)
		}
	}
	return false
}

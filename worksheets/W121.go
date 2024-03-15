package main

import (
	"fmt"
	"strconv"
)

func W121() {
	Board := [5][6]string{}
	playerGo := false
	FillBoard(Board)
	for isWon(Board) {
		PrintBoard(Board)
		PlayerGo(Board, playerGo)
		if playerGo {
			playerGo = false
		} else {
			playerGo = true
		}
	}
	PrintBoard(Board)
}

func PrintBoard(Board [5][6]string) {
	fmt.Println("1 2 3 4 5 6 7")
	for i := range Board {
		for j, val := range Board[i] {
			if j == 0 {
				fmt.Printf("%v|%v|", i, val)
			}
			fmt.Printf("|%v|", val)
		}
		fmt.Println()
	}
}

func FillBoard(Board [5][6]string) {
	for i := range Board {
		for j := range Board[i] {
			Board[i][j] = " "
		}
	}
}

func PlayerGo(Board [5][6]string, player bool) {
	var Play string
	fmt.Println("where would you like to go")
	fmt.Scan(&Play)
	num, _ := strconv.ParseInt(Play, 10, 64)
	if player {
		Board[0][num] = "o"
	} else {
		Board[0][num] = "0"
	}
	for isSpaced(Board) {
		MoveDown(Board)
	}
}

func MoveDown(Board [5][6]string) {
	for i := range Board {
		for j := range Board[i] {
			if Board[i][j] != " " && Board[i-1][j] == " " {
				Board[i-1][j] = "o"
				Board[i][j] = " "
			}
		}
	}
}

func isSpaced(Board [5][6]string) bool {
	for i := range Board {
		for j := range Board[i] {
			if Board[i][j] == "o" && i-1 < len(Board) && Board[i-1][j] == " " {
				return true
			}
		}
	}
	return false
}

func isWon(Board [5][6]string) bool {
	//has won or not!

	for i := range Board {
		for j, val := range Board[i] {
			if val == "0" || val == "o" {
				//isWonCircle(Board, i, j)
			}
		}
	}
}

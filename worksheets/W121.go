package worksheets

import (
	"fmt"
	"strconv"
)

func W121() {
	Board := [6][7]string{}
	player := false
	Board = fillBoard(Board)
	for !isWon(Board) {
		printBoard(Board)
		Board = playerGo(Board, player)
		if player {
			player = false
		} else {
			player = true
		}
	}
	printBoard(Board)
}

func printBoard(Board [6][7]string) {
	fmt.Println("  1  2  3  4  5  6  7")
	for y := range Board {
		for x, val := range Board[y] {
			if x == 0 {
				fmt.Printf("%v|%v|", y+1, val)
			} else {
				fmt.Printf("|%v|", val)
			}
		}
		fmt.Println()
	}
}

func fillBoard(Board [6][7]string) [6][7]string {
	for y := range Board {
		for x := range Board[y] {
			Board[y][x] = " "
		}
	}
	return Board
}

func playerGo(Board [6][7]string, player bool) [6][7]string {
	var Play string
	fmt.Println("where would you like to go")
	fmt.Scan(&Play)
	num, _ := strconv.ParseInt(Play, 10, 8)
	if player {
		Board[0][num-1] = "o"
	} else {
		Board[0][num-1] = "0"
	}
	y, x := isSpaced(Board)
	if y > -1 && x > -1 {
		Board = moveDown(Board, y, x)
	}
	return Board
}

func moveDown(Board [6][7]string, y int, x int) [6][7]string {
	for y-//some number here
}

func isSpaced(Board [6][7]string) (int, int) {
	for y := range Board {
		for x := range Board[0] {
			if Board[y][x] != " " && y-1 > len(Board) && Board[y-1][x] == " " {
				return y, x
			}
		}
	}
	return -1, -1
}

func isWon(Board [6][7]string) bool {
	//horizontal check
	for y := 0; y < len(Board); y++ {
		for x := 0; x < len(Board[0])-3; x++ {
			if Board[y][x] != " " && Board[y][x+1] != " " && Board[y][x+2] != " " && Board[y][x+3] != " " {
				return true
			}
		}
	}

	//vertical check
	for y := 0; y < len(Board)-3; y++ {
		for x := 0; x < len(Board[0]); x++ {
			if Board[y][x] != " " && Board[y+1][x] != " " && Board[y+2][x] != " " && Board[y+3][x] != " " {
				return true
			}
		}
	}

	//ascending diagonal check
	for y := 3; y < len(Board); y++ {
		for x := 0; x < len(Board[0])-3; x++ {
			if Board[y][x] != " " && Board[y-1][x+1] != " " && Board[y-2][x+2] != " " && Board[y-3][x+3] != " " {
				return true
			}
		}
	}

	//descending diagonal check
	for x := 3; x < len(Board[0]); x++ {
		for y := 3; y < len(Board); y++ {
			if Board[y][x] != " " && Board[y-1][x-1] != " " && Board[y-2][x-2] != " " && Board[y-3][x-3] != " " {
				return true
			}
		}
	}
	return false
}

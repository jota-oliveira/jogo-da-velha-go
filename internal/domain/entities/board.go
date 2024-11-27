package entities

import (
	"fmt"
)

type Board struct {
	Board [3][3]string
}

func (board *Board) PrintBoard() {
	for line := 0; line < 3; line++ {
		fmt.Println("----------")
		fmt.Print("| ")
		for column := 0; column < 3; column++ {
			fmt.Print(board.Board[line][column], " | ")
		}
		fmt.Println()
	}
}

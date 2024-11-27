package usecases

import "github.com/jota-oliveira/joga-da-velha-go/internal/domain/entities"

func HasHorizontalWinner(board *entities.Board, symbol string) bool {
	for line := 0; line < 3; line++ {
		if board.Board[line][0] == symbol && board.Board[line][1] == symbol && board.Board[line][2] == symbol {
			return true
		}
	}
	return false
}

func HasVerticalWinner(board *entities.Board, symbol string) bool {
	for column := 0; column < 3; column++ {
		if board.Board[0][column] == symbol && board.Board[1][column] == symbol && board.Board[2][column] == symbol {
			return true
		}
	}
	return false
}

func HasDiagonalWinner(board *entities.Board, symbol string) bool {
	if board.Board[0][0] == symbol && board.Board[1][1] == symbol && board.Board[2][2] == symbol {
		return true
	}
	if board.Board[0][2] == symbol && board.Board[1][1] == symbol && board.Board[2][0] == symbol {
		return true
	}
	return false
}

func CheckWinner(board *entities.Board, symbol string) bool {
	if HasHorizontalWinner(board, symbol) {
		return true
	}
	if HasVerticalWinner(board, symbol) {
		return true
	}
	if HasDiagonalWinner(board, symbol) {
		return true
	}
	return false
}

func IsBoardFull(board *entities.Board) bool {
	for line := 0; line < 3; line++ {
		for column := 0; column < 3; column++ {
			if board.Board[line][column] == "" {
				return false
			}
		}
	}
	return true
}

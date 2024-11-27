package usecases

import (
	"fmt"

	"github.com/jota-oliveira/joga-da-velha-go/internal/domain/entities"
)

func StartGame(board *entities.Board, firstPlayer, secondPlayer *entities.Player) {
	currentPlayer := firstPlayer

	for {
		if IsBoardFull(board) {
			fmt.Println("O jogo empatou!")
			break
		}

		fmt.Println("Vez do jogador", currentPlayer.Name)

		var line int
		var column int

		fmt.Println("Escolha a linha (0, 1 ou 2)")
		fmt.Scanln(&line)

		fmt.Println("Escolha a coluna (0, 1 ou 2)")
		fmt.Scanln(&column)

		if board.Board[line][column] == "" {
			board.Board[line][column] = currentPlayer.Symbol
		} else {
			fmt.Println("Posição já ocupada, tente novamente!")
			continue
		}

		if CheckWinner(board, currentPlayer.Symbol) {
			fmt.Println("Parabéns ", currentPlayer.Name, "você ganhou!")
			break
		}

		board.PrintBoard()

		if currentPlayer == firstPlayer {
			currentPlayer = secondPlayer
		} else {
			currentPlayer = firstPlayer
		}
	}
}

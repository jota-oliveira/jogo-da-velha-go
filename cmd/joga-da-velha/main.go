/*
	Jogo da velha em GoLang para praticar programação e explorar os recursos da linguagem GO.
*/

package main

import (
	"github.com/jota-oliveira/joga-da-velha-go/internal/domain/entities"
	"github.com/jota-oliveira/joga-da-velha-go/internal/domain/services"
	"github.com/jota-oliveira/joga-da-velha-go/internal/usecases"
)

func main() {
	var firstPlayer string
	var secondPlayer string

	firstPlayerObj, secondPlayerObj := services.SetupGame(&firstPlayer, &secondPlayer)

	board := entities.Board{}
	board.PrintBoard()

	usecases.StartGame(&board, firstPlayerObj, secondPlayerObj)
}

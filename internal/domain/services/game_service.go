package services

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/jota-oliveira/joga-da-velha-go/internal/domain/entities"
)

func SetupGame(firstPlayer, secondPlayer *string) (*entities.Player, *entities.Player) {
	fmt.Println("Joga da velha em GoLang")

	fmt.Println("Digite o nome do primeiro jogador: ")
	fmt.Scanln(firstPlayer)

	fmt.Println("Digite o nome do segundo jogador: ")
	fmt.Scanln(secondPlayer)

	player1 := entities.Player{}
	player2 := entities.Player{}

	player1.SetName(*firstPlayer)
	player2.SetName(*secondPlayer)

	randomIntGenerator := rand.New(rand.NewSource(time.Now().UnixNano()))

	fmt.Println("Sorteando o jogador que começará o jogo...", randomIntGenerator.Intn(2))

	if randomIntGenerator.Intn(2) == 0 {
		player1.SetSymbol("X")
		player2.SetSymbol("O")
	} else {
		player2.SetSymbol("X")
		player1.SetSymbol("O")
	}

	fmt.Println("Tudo pronto para começar o jogo!")
	fmt.Println("O jogador", player1.Name, "ficará com o símbolo", player1.Symbol)
	fmt.Println("O jogador", player2.Name, "ficará com o símbolo", player2.Symbol)

	return &player1, &player2
}

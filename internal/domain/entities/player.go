package entities

import (
	"errors"
	"strings"
)

type Player struct {
	Name   string
	Symbol string
}

func (player *Player) SetName(name string) {
	if len(name) == 0 {
		player.Name = "Player"
	} else {
		player.Name = strings.ToUpper(string(name[0])) + name[1:]
	}
}

func (player *Player) SetSymbol(symbol string) error {
	if symbol != "X" && symbol != "O" {
		return errors.New("o símbolo do jogador deve ser X ou 0")
	}

	player.Symbol = symbol

	return nil
}

package main

import (
	"fmt"
	"os"

	"github.com/alvinobarboza/snake/internal/game"
	"github.com/alvinobarboza/snake/internal/player"
	"github.com/olekukonko/ts"
	"golang.org/x/term"
)

func main() {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	s, e := ts.GetSize()

	if e != nil {
		fmt.Println(e)
		return
	}

	player := player.NewPlayer()

	game := game.NewGame(player)

	game.CreateCanvas(s.Col(), s.Row())

	exit := make(chan string)
	go game.ProcessKey(exit)

	for {
		select {
		case message := <-exit:
			fmt.Print(message)
			return

		default:
			game.Update()
		}
	}
}

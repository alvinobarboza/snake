package main

import (
	"fmt"
	"os"

	"github.com/alvinobarboza/snake/internal"
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

	input := make(chan internal.InputKey)
	go listenInput(input)

	for {
		select {
		case key := <-input:
			//TODO: Better exit handling
			if game.ProcessKey(key) {
				return
			}
		default:
			game.Update()
		}
	}
}

func listenInput(input chan internal.InputKey) {
	b := make([]byte, 1)
	for {
		os.Stdin.Read(b)
		input <- internal.InputKey(b)
	}
}

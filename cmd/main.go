package main

import (
	"fmt"
	"os"

	"github.com/alvinobarboza/snake/internal"
	"github.com/alvinobarboza/snake/internal/game"
	"github.com/alvinobarboza/snake/internal/player"
	"golang.org/x/term"
)

func main() {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	w, h, errS := term.GetSize(int(os.Stdin.Fd()))

	if errS != nil {
		fmt.Println(err)
		return
	}

	player := player.NewPlayer()

	game := game.NewGame(player)

	game.CreateCanvas(w, h)

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

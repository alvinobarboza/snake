package main

import (
	"fmt"
	"os"
	"strconv"

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

	if len(os.Args) < 3 {
		fmt.Print(
			"Inform width and heigth:",
			"\n\r",
			".\\game w h\n\r",
			"i.e.\n\r\n\r",
			".\\game 40 15\n\r")
		return
	}

	ws, hs := os.Args[1], os.Args[2]
	w, err := strconv.Atoi(ws)
	if err != nil {
		fmt.Print("Width must be a number: ", ws, "\n\r")
		return
	}
	h, err := strconv.Atoi(hs)
	if err != nil {
		fmt.Print("Heigth must be a number: ", hs, "\n\r")
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

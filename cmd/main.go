package main

import (
	"fmt"
	"os"
	"strconv"

	"golang.org/x/term"
)

type Direction string

const (
	UP    Direction = "w"
	DOWN  Direction = "s"
	LEFT  Direction = "a"
	RIGHT Direction = "d"
	QUIT  Direction = "q"
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
			".\\game w h",
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
	player := &Player{}

	game := NewGame(player)

	game.CreateCanvas(w, h)

	input := make(chan Direction)
	go listenInput(input)

	for {
		select {
		case key := <-input:
			if key == QUIT {
				fmt.Print("Exited", "\n\r")
				os.Exit(0)
			}
			game.p.ProcessKey(key)
		default:
			game.Render()
		}
	}
}

func listenInput(input chan Direction) {
	b := make([]byte, 1)
	for {
		os.Stdin.Read(b)
		input <- Direction(b)
	}
}

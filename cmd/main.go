package main

import (
	"fmt"
	"os"
	"strconv"

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

	input := make(chan InputKey)
	go listenInput(input)

	for {
		select {
		case key := <-input:
			game.ProcessKey(key)
		default:
			game.Update()
		}
	}
}

func listenInput(input chan InputKey) {
	b := make([]byte, 1)
	for {
		os.Stdin.Read(b)
		input <- InputKey(b)
	}
}

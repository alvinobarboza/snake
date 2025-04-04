package main

import (
	"fmt"
	"os"
	"time"

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

	p := player.NewPlayer()
	t := player.NewTarget()
	ex := make(chan string)

	game := game.NewGame(p, t, ex)

	game.CreateCanvas(s.Col(), s.Row())

	go game.ProcessKey()

	go func() {
		for {
			game.Update()
			game.Render()
			time.Sleep(time.Millisecond * 120)
		}
	}()

	fmt.Print(<-ex)
}

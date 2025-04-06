package main

import (
	"fmt"
	"os"
	"time"

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

	if s.Col() < internal.MIN_SCREEN_WIDTH ||
		s.Row() < internal.MIN_SCREEN_HEIGHT {
		fmt.Printf("The terminal window must have a minimum size of %2dx%02d characters. \n\r",
			internal.MIN_SCREEN_WIDTH,
			internal.MIN_SCREEN_HEIGHT,
		)
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

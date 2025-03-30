package main

import (
	"fmt"
	"os"
	"time"
)

type InputKey string

const (
	UP    InputKey = "w"
	DOWN  InputKey = "s"
	LEFT  InputKey = "a"
	RIGHT InputKey = "d"
	QUIT  InputKey = "q"

	PADDING int = 2
)

type Game struct {
	p *Player

	w int
	h int

	canvas []string
}

func NewGame(p *Player) *Game {
	return &Game{
		p: p,
	}
}

func (g *Game) ProcessKey(key InputKey) {
	if key == QUIT {
		fmt.Print("Exited", "\n\r")
		os.Exit(0)
	}
	g.p.ProcessKey(key)
}

func (g *Game) Update() {
	g.clearScreen()

	g.p.Update()

	y := g.p.posY % g.h
	x := g.p.posX % g.w

	fmt.Print(g.canvas[y*g.h+x], "\n\r")

	g.Render()
	time.Sleep(time.Millisecond * 100)
}

func (g *Game) CreateCanvas(w, h int) {
	g.h = h
	g.w = w

	g.canvas = make([]string, 0)

	h += PADDING
	w += PADDING

	for hi := range h {
		for wi := range w {
			if hi == 0 && wi == 0 {
				g.canvas = append(g.canvas, "┌")
				continue
			}
			if hi == 0 && wi == w-1 {
				g.canvas = append(g.canvas, "┐")
				continue
			}
			if (hi == 0 || hi == h-1) && wi < w-1 && wi > 0 {
				g.canvas = append(g.canvas, "─")
				continue
			}
			if hi == h-1 && wi == 0 {
				g.canvas = append(g.canvas, "└")
				continue
			}
			if hi == h-1 && wi == w-1 {
				g.canvas = append(g.canvas, "┘")
				continue
			}
			if wi == 0 || wi == w-1 {
				g.canvas = append(g.canvas, "│")
				continue
			}
			g.canvas = append(g.canvas, " ")
		}
		g.canvas = append(g.canvas, "\n\r")
	}
}

func (g *Game) Render() {
	for _, s := range g.canvas {
		fmt.Print(s)
	}
}

func (g *Game) clearScreen() {
	fmt.Print("\033[H\033[2J", g.p, len(g.canvas), g.h*g.w, "\n\r")
}

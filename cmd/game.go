package main

import (
	"fmt"
	"time"
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

func (g *Game) CreateCanvas(w, h int) {
	g.h = h
	g.w = w

	g.canvas = make([]string, 0)

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
	time.Sleep(time.Millisecond * 33)
	fmt.Print("\033[H\033[2J", g.p, len(g.canvas), g.h*g.w, "\n\r")
	for _, s := range g.canvas {
		fmt.Print(s)
	}
}

package game

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/alvinobarboza/snake/internal"
	"github.com/alvinobarboza/snake/internal/player"
)

type Game struct {
	p player.Player

	w int
	h int

	canvas []string
}

func NewGame(p player.Player) *Game {
	return &Game{
		p: p,
	}
}

func (g *Game) ProcessKey(key internal.InputKey) {
	if key == internal.QUIT {
		fmt.Print("Exited", "\n\r")
		os.Exit(0)
	}
	g.p.ProcessKey(key)
}

func (g *Game) Update() {

	g.p.Update()
	g.clearScreen()

	i := g.normalizedIndex()

	fmt.Print(g.canvas[i], "\n\r")

	g.Render()
	time.Sleep(time.Millisecond * 100)
}

func (g *Game) CreateCanvas(w, h int) {
	g.h = h
	g.w = w

	g.canvas = make([]string, 0)

	h += internal.PADDING
	w += internal.PADDING

	temp := 0
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
			if temp == 10 {
				temp = 0
			}
			temp++
			stemp := strconv.Itoa(temp)
			g.canvas = append(g.canvas, stemp)
		}
		temp = 0
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

func (g *Game) normalizedIndex() int {
	posX, posY := g.p.GetPosXY()
	x := 0
	y := 0
	if posX < 0 {
		x = (g.w - 1) - ((posX * -1) % g.w)
	} else {
		x = posX % g.w
	}
	if posY < 0 {
		y = (g.h - 1) - ((posY * -1) % g.h)
	} else {
		y = posY % g.h
	}
	return y*g.w + x
}

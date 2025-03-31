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

	canvas  []string
	borders []string
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

	for range h {
		for range w {
			g.canvas = append(g.canvas, g.emptyChar)
		}
	}

	for i := range internal.BORDERS {
		if i == 0 {
			g.borders = append(g.borders, "┌")
		} else {
			g.borders = append(g.borders, "└")
		}
		for range w {
			g.borders = append(g.borders, "─")
		}
		if i == 0 {
			g.borders = append(g.borders, "┐")
		} else {
			g.borders = append(g.borders, "┘")
		}
	}
}

func (g *Game) Render() {
	borderWidth := len(g.borders) / 2

	renderString := ""

	for i := range borderWidth {
		renderString += g.borders[i]
	}
	renderString += "\n\r "

	for y := range g.h {
		for x := range g.w {
			renderString += g.canvas[y*g.w+x]
		}
		if y == g.h-1 {
			renderString += " \n\r"
			continue
		}
		renderString += " \n\r "
	}

	for i := range borderWidth {
		renderString += g.borders[i+borderWidth]
	}
	fmt.Print(renderString + "\n\r")
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

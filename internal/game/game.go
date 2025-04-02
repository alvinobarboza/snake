package game

import (
	"fmt"
	"math/rand/v2"
	"os"

	"github.com/alvinobarboza/snake/internal"
	"github.com/alvinobarboza/snake/internal/player"
)

type Game struct {
	p player.Player

	w int
	h int

	canvas  []string
	borders []string

	emptyChar string
	pointChar string
}

func NewGame(p player.Player) *Game {
	return &Game{
		p:         p,
		emptyChar: " ",
		pointChar: "X",
	}
}

func (g *Game) ProcessKey(exit chan string) {
	b := make([]byte, 1)
	for {
		os.Stdin.Read(b)
		key := internal.InputKey(b)
		if key == internal.QUIT {
			exit <- fmt.Sprint("Exited\033[0J", "\n\r")
			break
		}
		g.p.ProcessKey(key)
	}
}

func (g *Game) Update() {

	ix := g.normalizedIndex(g.p.GetNextPosXY())
	if g.canvas[ix] == g.pointChar {
		spawIndex := g.spawnPoint(ix)
		g.canvas[spawIndex] = g.pointChar
		g.p.GrowTail()
	}

	g.p.Update()

	i := g.normalizedIndex(g.p.GetPosXY())

	i_last := g.normalizedIndex(g.p.GetLastPosXY())
	g.canvas[i] = g.p.Visuals()

	for _, t := range g.p.GetTail() {
		ix := g.normalizedIndex(t.GetXY())
		g.canvas[ix] = g.p.Visuals()
	}
	if i_last != i {
		g.canvas[i_last] = g.emptyChar
	}
}

func (g *Game) CreateCanvas(w, h int) {
	h -= internal.PADDING_TOP_BOTTOM
	w -= internal.PADDING_SIDES

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

	i := g.spawnPoint(0)
	g.canvas[i] = g.pointChar
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

	g.clearScreen()
}

func (g *Game) clearScreen() {
	fmt.Printf("\033[%dA", g.h+2)
}

func (g *Game) normalizedIndex(posX, posY int) int {
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

func (g *Game) spawnPoint(i_avoid int) int {
	i := 0
	for {
		i = rand.IntN(g.h * g.w)
		if i != i_avoid {
			break
		}
	}
	return i
}

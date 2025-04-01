package game

import (
	"fmt"
	"math/rand/v2"
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

	emptyChar  string
	playerChar string
	pointChar  string
}

func NewGame(p player.Player) *Game {
	return &Game{
		p:          p,
		emptyChar:  " ",
		playerChar: "■",
		pointChar:  "X",
	}
}

func (g *Game) ProcessKey(key internal.InputKey) bool {
	if key == internal.QUIT {
		fmt.Print("Exited\033[0J", "\n\r")
		return true
	}
	g.p.ProcessKey(key)
	return false
}

func (g *Game) Update() {

	g.p.Update()

	i_last := g.normalizedLastIndex(g.p.GetLastPosXY())
	i := g.normalizedIndex(g.p.GetPosXY())

	g.canvas[i_last] = g.emptyChar
	if g.canvas[i] == g.pointChar {
		g.spawnPoint(i)
	}
	g.canvas[i] = g.playerChar

	g.Render()

	time.Sleep(time.Millisecond * 120)
	g.clearScreen()
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

	g.spawnPoint(0)
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

func (g *Game) normalizedLastIndex(posX, posY int) int {
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

func (g *Game) spawnPoint(i_avoid int) {
	i := 0
	for {
		i = rand.IntN(g.h * g.w)
		if i != i_avoid {
			break
		}
	}
	g.canvas[i] = g.pointChar
}

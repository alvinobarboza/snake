package game

import (
	"fmt"
	"os"

	"github.com/alvinobarboza/snake/internal"
	"github.com/alvinobarboza/snake/internal/player"
)

type Game struct {
	p player.Player
	t player.Target

	w int
	h int

	canvas  []string
	borders []string

	bg string

	exit chan string

	howToMessage string
}

func NewGame(
	p player.Player,
	t player.Target,
	exit chan string) *Game {
	return &Game{
		p:            p,
		t:            t,
		bg:           " ",
		exit:         exit,
		howToMessage: internal.HowToMessage(),
	}
}

func (g *Game) ProcessKey() {
	b := make([]byte, 3)
	for {
		n, _ := os.Stdin.Read(b)
		key := internal.InputKey(b[:n])
		if key == internal.QUIT {
			g.exit <- fmt.Sprint("Exited\x1b[0J", "\n\r")
			break
		}
		g.p.ProcessKey(key)
	}
}

func (g *Game) Update() {

	if g.p.SelfCollide(g.w, g.h) {
		g.exit <- g.messageOnLost()
		return
	}
	hasGrown := false
	if g.t.Index() == g.p.NextIndex(g.w, g.h) {

		g.p.GrowTail()
		hasGrown = true

		if len(g.p.GetTail())+1 == g.h*g.w {
			g.exit <- fmt.Sprint("\x1b[0J\n\r", "YOU WON!", "\n\r")
			return
		}

		for {
			g.t.SpawNewLocation(g.p.GetTail())
			if g.t.Index() != g.p.NextIndex(g.w, g.h) &&
				g.t.Index() != g.p.Index(g.w, g.h) {
				break
			}
		}
	}

	g.p.Update(hasGrown)

	i := g.p.Index(g.w, g.h)

	i_last := g.p.LastIndex(g.w, g.h)
	g.canvas[i] = g.p.Visuals()

	for _, t := range g.p.GetTail() {
		ix := t.Index(g.w, g.h)
		g.canvas[ix] = g.p.VisualsTail()
	}
	if i_last != i {
		g.canvas[i_last] = g.bg
	}
	g.canvas[g.t.Index()] = g.t.Visuals()

}

func (g *Game) CreateCanvas(w, h int) {
	h -= (internal.PADDING_BOTTOM + internal.PADDING_TOP)
	w -= internal.PADDING_SIDES

	g.h = h
	g.w = w

	g.canvas = make([]string, 0)

	for range h {
		for range w {
			g.canvas = append(g.canvas, g.bg)
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

	g.t.AddSeed(g.w, g.h)

	for {
		g.t.SpawNewLocation(g.p.GetTail())
		if g.t.Index() != g.p.Index(g.w, g.h) {
			break
		}
	}
	// fmt.Print(g.normalizedIndex(g.t.GetPosXY()))
	g.canvas[g.t.Index()] = g.t.Visuals()
}

func (g *Game) Render() {

	borderWidth := len(g.borders) / 2

	renderString := fmt.Sprintf(" [Points %03d/%03d] ",
		len(g.p.GetTail()), (g.h*g.w)-2)

	header := len(renderString)

	for i := range borderWidth {
		if i < header {
			continue
		}
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

	renderString += g.howToMessage

	fmt.Print(renderString)

	g.clearScreen()
}

func (g *Game) clearScreen() {
	fmt.Printf("\x1b[%dA", g.h+internal.PADDING_TOP+internal.PADDING_BOTTOM)
}

func (g *Game) messageOnLost() string {

	message := "SKILL ISSUE!!! YOU LOST!"

	points := len(g.p.GetTail())
	size := g.h * g.w

	percentage := float32(points) / float32(size) * 100

	if points > 5 {
		message = "NOT SO BAD! BUT YOU LOST"
	}

	if percentage > 50 {
		message = "HALF WAY! BUT YOU LOST"
	}

	if percentage > 80 {
		message = "GETTING THERE! BUT YOU LOST"
	}

	if percentage > 95 {
		message = "IMPRESSIVE! BUT YOU LOST"
	}

	return fmt.Sprint(
		"\x1b[0J\n\r",
		message,
		"\n\r\n\r")
}

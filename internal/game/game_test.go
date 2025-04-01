package game

import (
	"testing"

	"github.com/alvinobarboza/snake/internal"
)

type playerTest struct {
}

func (p playerTest) Update() {}
func (p playerTest) GetPosXY() (int, int) {
	return 0, 1
}
func (p playerTest) GetLastPosXY() (int, int) {
	return 0, 1
}
func (p playerTest) ProcessKey(internal.InputKey) {}

func TestPlayerPos(t *testing.T) {
	p := playerTest{}
	g := NewGame(p)

	g.h = 2
	g.w = 2
	i := g.normalizedIndex(g.p.GetPosXY())

	if i < 0 {
		t.Error("Expected positive, got:", i)
	}
}

func TestScreenGen(t *testing.T) {
	p := playerTest{}
	g := NewGame(p)

	want_border := []string{
		"┌", "─", "─", "┐",
		"└", "─", "─", "┘",
	}

	width := 2
	height := 2

	border := 2

	g.CreateCanvas(width, height)

	if len(g.canvas) != (width * height) {
		t.Errorf("Expected %d, got: %d", (width * height), len(g.canvas))

	}
	broderSizeWant := (width + border) * 2
	if len(g.borders) != broderSizeWant {
		t.Errorf("Expected %d, got: %d", broderSizeWant, len(g.borders))
	}

	checkWant := true

	for i := range broderSizeWant {
		if want_border[i] != g.borders[i] {
			checkWant = false
		}
	}

	if !checkWant {
		t.Errorf("Expected %v, got: %v", true, checkWant)
	}
}

func TestRandomSpawn(t *testing.T) {
	p := playerTest{}
	g := NewGame(p)

	width := 2
	height := 2

	g.CreateCanvas(width, height)

	g.spawnPoint(0)

	want := "X"
	got := "-"
	for _, s := range g.canvas {
		if s == g.pointChar {
			got = s
		}
	}

	if want != got {
		t.Errorf("Want: %s, got: %s", want, got)
	}
}

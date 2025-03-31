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
	i := g.normalizedIndex()

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

	g.CreateCanvas(2, 2)

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

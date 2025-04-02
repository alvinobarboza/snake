package game

import (
	"testing"

	"github.com/alvinobarboza/snake/internal"
)

const (
	width  int = internal.PADDING_SIDES + 3
	height int = internal.PADDING_TOP_BOTTOM + 3
)

type playerTest struct {
}

func (p playerTest) Update()         {}
func (p playerTest) Visuals() string { return "" }
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
	p := &playerTest{}
	g := NewGame(p)

	preComputedWidth := (width - internal.PADDING_SIDES)
	preComputedHeight := (height - internal.PADDING_TOP_BOTTOM)

	want_border := make([]string, 0)
	want_border = append(want_border, "┌")
	for range preComputedWidth {
		want_border = append(want_border, "─")
	}
	want_border = append(want_border, "┐")
	want_border = append(want_border, "└")
	for range preComputedWidth {
		want_border = append(want_border, "─")
	}
	want_border = append(want_border, "┘")

	g.CreateCanvas(width, height)

	if g.w != preComputedWidth ||
		g.h != preComputedHeight {
		t.Errorf(
			"Wanted %d %d, got: %d %d",
			preComputedHeight,
			preComputedWidth,
			g.h, g.w)
	}

	wantCanvas := (preComputedWidth * preComputedHeight)

	if len(g.canvas) != wantCanvas {
		t.Errorf("Expected %d, got: %d", (width * height), len(g.canvas))

	}
	broderSizeWant := (preComputedWidth + internal.BORDERS) * 2
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

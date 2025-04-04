package game

import (
	"testing"

	"github.com/alvinobarboza/snake/internal"
	"github.com/alvinobarboza/snake/internal/player"
)

const (
	width  int = internal.PADDING_SIDES + 3
	height int = internal.PADDING_TOP_BOTTOM + 3
)

type playerTest struct {
}

func (p playerTest) Update(bool)             {}
func (p playerTest) GrowTail()               {}
func (p playerTest) Collision(x, y int) bool { return false }
func (p playerTest) Visuals() string         { return "X" }
func (p playerTest) VisualsTail() string     { return "X" }
func (p playerTest) Index(int, int) int {
	return 1
}
func (p playerTest) NextIndex(int, int) int {
	return 1
}
func (p playerTest) LastIndex(int, int) int {
	return 1
}
func (p playerTest) GetTail() []player.Transform {
	return make([]player.Transform, 0)
}
func (p playerTest) ProcessKey(internal.InputKey) {}

type tar struct{}

func (p tar) Index() int                         { return 5 }
func (p tar) Visuals() string                    { return "X" }
func (p tar) AddSeed(int, int)                   {}
func (p tar) SpawNewLocation([]player.Transform) {}

func TestPlayerPos(t *testing.T) {
	p := playerTest{}
	ta := &tar{}
	g := NewGame(p, ta, make(chan string))

	g.h = 2
	g.w = 2
	i := g.p.Index(g.w, g.h)

	if i < 0 {
		t.Error("Expected positive, got:", i)
	}
}

func TestScreenGen(t *testing.T) {
	p := &playerTest{}
	ta := &tar{}
	g := NewGame(p, ta, make(chan string))

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
	ta := &tar{}
	g := NewGame(p, ta, make(chan string))

	g.CreateCanvas(width, height)

	want := "X"
	got := "-"
	for _, s := range g.canvas {
		if s == g.t.Visuals() {
			got = s
		}
	}

	if want != got {
		t.Errorf("Want: %s, got: %s", want, got)
	}
}

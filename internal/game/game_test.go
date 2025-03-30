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

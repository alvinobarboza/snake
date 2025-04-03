package player

import "github.com/alvinobarboza/snake/internal"

type cood struct {
	x, y int
}

type Transform struct {
	lastPos cood
	curPos  cood
}

func (t Transform) LastIndex(w, h int) int {
	return internal.NormalizedIndex(
		t.lastPos.x,
		t.lastPos.y,
		w, h,
	)
}

func (t Transform) Index(w, h int) int {
	return internal.NormalizedIndex(
		t.curPos.x,
		t.curPos.y,
		w, h,
	)
}

package player

import (
	"fmt"
	"math/rand/v2"
)

type Target interface {
	Index() int
	SpawNewLocation(tail []Transform)
	Visuals() string
	AddSeed(x, y int)
}

func NewTarget() *target {
	return &target{
		visuals: "●",
	}
}

type target struct {
	idx     int
	visuals string
	seed    cood
}

func (t *target) Index() int {
	return t.idx
}

func (t *target) SpawNewLocation(tail []Transform) {
	ix := 0
	isNotCollinding := true

	for {
		isNotCollinding = true
		ix = rand.IntN(t.seed.x * t.seed.y)
		for _, tl := range tail {
			if ix == tl.Index(t.seed.x, t.seed.y) {
				isNotCollinding = false
				break
			}
			fmt.Print(ix, tl.Index(t.seed.x, t.seed.y), "\r")
		}
		if isNotCollinding {
			break
		}
	}
	t.idx = ix
}

func (t *target) Visuals() string {
	return t.visuals
}

func (t *target) AddSeed(x, y int) {
	t.seed.x = x
	t.seed.y = y
}

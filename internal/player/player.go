package player

import (
	"sync"

	"github.com/alvinobarboza/snake/internal"
)

type Player interface {
	Index(w, h int) int
	NextIndex(w, h int) int
	LastIndex(w, h int) int
	GetTail() []Transform
	Visuals() string
	VisualsTail() string
	Update(hasGrown bool)
	ProcessKey(key internal.InputKey)
	GrowTail()
}

type player struct {
	mu        sync.Mutex
	direction cood
	head      Transform

	tail []Transform

	playerChar string
	tailChar   string
}

func NewPlayer() *player {
	return &player{
		tailChar:   "⩩",
		playerChar: "█",
		tail:       make([]Transform, 0),
	}
}

func (p *player) NextIndex(w, h int) int {
	p.mu.Lock()
	defer p.mu.Unlock()
	x := p.head.curPos.x + p.direction.x
	y := p.head.curPos.y + p.direction.y

	return internal.NormalizedIndex(x, y, w, h)
}

func (p *player) Index(w, h int) int {
	p.mu.Lock()
	defer p.mu.Unlock()
	return internal.NormalizedIndex(
		p.head.curPos.x,
		p.head.curPos.y,
		w, h,
	)
}

func (p *player) LastIndex(w, h int) int {
	p.mu.Lock()
	defer p.mu.Unlock()
	if len(p.tail) < 1 {
		return internal.NormalizedIndex(
			p.head.lastPos.x,
			p.head.lastPos.y,
			w, h,
		)
	}
	return internal.NormalizedIndex(
		p.tail[0].lastPos.x,
		p.tail[0].lastPos.y,
		w, h,
	)
}

func (p *player) GetTail() []Transform {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.tail
}

func (p *player) Visuals() string {
	return p.playerChar
}

func (p *player) VisualsTail() string {
	return p.tailChar
}

func (p *player) Update(hasGrown bool) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if !hasGrown {
		lt := len(p.tail)
		for i := range lt {
			if i == lt-1 {
				p.tail[i] = p.head
				continue
			}
			p.tail[i] = p.tail[i+1]
		}
	}
	p.head.lastPos.x = p.head.curPos.x
	p.head.lastPos.y = p.head.curPos.y

	p.head.curPos.x += p.direction.x
	p.head.curPos.y += p.direction.y
}

func (p *player) ProcessKey(key internal.InputKey) {
	p.mu.Lock()
	defer p.mu.Unlock()

	switch key {
	case internal.UP:
		p.direction.y = -1
		p.direction.x = 0
		return
	case internal.DOWN:
		p.direction.y = 1
		p.direction.x = 0
		return
	case internal.RIGHT:
		p.direction.y = 0
		p.direction.x = 1
		return
	case internal.LEFT:
		p.direction.y = 0
		p.direction.x = -1
		return
	}
}

func (p *player) GrowTail() {
	p.tail = append(p.tail, p.head)
}

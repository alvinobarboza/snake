package player

import (
	"sync"

	"github.com/alvinobarboza/snake/internal"
)

type Player interface {
	GetPosXY() (int, int)
	GetNextPosXY() (int, int)
	GetLastPosXY() (int, int)
	GetTail() []Transform
	Visuals() string
	Update()
	ProcessKey(key internal.InputKey)
	GrowTail()
}

type player struct {
	mu        sync.Mutex
	direction cood
	head      Transform

	tail []Transform

	playerChar string
}

func NewPlayer() *player {
	return &player{
		playerChar: "■",
		tail:       make([]Transform, 0),
	}
}

func (p *player) GetNextPosXY() (int, int) {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.head.curPos.x + p.direction.x,
		p.head.curPos.y + p.direction.y
}

func (p *player) GetPosXY() (int, int) {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.head.curPos.x, p.head.curPos.y
}

func (p *player) GetLastPosXY() (int, int) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if len(p.tail) < 1 {
		return p.head.lastPos.x, p.head.lastPos.y
	}
	return p.tail[0].lastPos.x, p.tail[0].lastPos.y
}

func (p *player) GetTail() []Transform {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.tail
}

func (p *player) Visuals() string {
	return p.playerChar
}

func (p *player) Update() {
	p.mu.Lock()
	defer p.mu.Unlock()

	lt := len(p.tail)
	for i := range lt {
		if i == lt-1 {
			p.tail[i] = p.head
			continue
		}
		p.tail[i] = p.tail[i+1]
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

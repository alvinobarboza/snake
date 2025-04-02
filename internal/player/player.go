package player

import (
	"sync"

	"github.com/alvinobarboza/snake/internal"
)

type Player interface {
	GetPosXY() (int, int)
	GetLastPosXY() (int, int)
	Visuals() string
	Update()
	ProcessKey(key internal.InputKey)
}

type player struct {
	mu sync.Mutex

	x int
	y int

	posX int
	posY int

	lastPosX int
	lastPosY int

	playerChar string
}

func NewPlayer() *player {
	return &player{playerChar: "■"}
}

func (p *player) GetPosXY() (int, int) {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.posX, p.posY
}

func (p *player) GetLastPosXY() (int, int) {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.lastPosX, p.lastPosY
}

func (p *player) Visuals() string {
	return p.playerChar
}

func (p *player) Update() {
	p.lastPosX = p.posX
	p.lastPosY = p.posY

	p.posX += p.x
	p.posY += p.y
}

func (p *player) ProcessKey(key internal.InputKey) {
	p.mu.Lock()
	defer p.mu.Unlock()

	switch key {
	case internal.UP:
		p.y = -1
		p.x = 0
		return
	case internal.DOWN:
		p.y = 1
		p.x = 0
		return
	case internal.RIGHT:
		p.y = 0
		p.x = 1
		return
	case internal.LEFT:
		p.y = 0
		p.x = -1
		return
	}
}

package main

type Player struct {
	x int
	y int

	posX int
	posY int
}

func (p *Player) Update() {
	p.posX += p.x
	p.posY += p.y
}

func (p *Player) ProcessKey(key InputKey) {
	switch key {
	case UP:
		p.y = 1
		p.x = 0
		return
	case DOWN:
		p.y = -1
		p.x = 0
		return
	case RIGHT:
		p.y = 0
		p.x = 1
		return
	case LEFT:
		p.y = 0
		p.x = -1
		return
	}
}

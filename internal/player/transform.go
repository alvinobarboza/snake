package player

type cood struct {
	x, y int
}

type Transform struct {
	lastPos cood
	curPos  cood
}

func (t Transform) GetLastXY() (int, int) {
	return t.lastPos.x, t.lastPos.y
}

func (t Transform) GetXY() (int, int) {
	return t.lastPos.x, t.lastPos.y
}

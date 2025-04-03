package internal

type InputKey string

const (
	UP    InputKey = "w"
	DOWN  InputKey = "s"
	LEFT  InputKey = "a"
	RIGHT InputKey = "d"
	QUIT  InputKey = "q"

	BORDERS int = 2

	PADDING_SIDES      int = 6
	PADDING_TOP_BOTTOM int = 4
)

func NormalizedIndex(posX, posY, w, h int) int {
	x := 0
	y := 0
	if posX < 0 {
		x = (w - 1) - ((posX * -1) % w)
	} else {
		x = posX % w
	}
	if posY < 0 {
		y = (h - 1) - ((posY * -1) % h)
	} else {
		y = posY % h
	}
	return y*w + x
}

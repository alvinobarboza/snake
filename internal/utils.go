package internal

type InputKey string

const (
	UP          InputKey = "w"
	DOWN        InputKey = "s"
	LEFT        InputKey = "a"
	RIGHT       InputKey = "d"
	QUIT        InputKey = "q"
	UP_ARROW    InputKey = "\x1b[A"
	DOWN_ARROW  InputKey = "\x1b[B"
	RIGHT_ARROW InputKey = "\x1b[C"
	LEFT_ARROW  InputKey = "\x1b[D"

	BORDERS int = 2

	PADDING_SIDES  int = 4
	PADDING_TOP    int = 2
	PADDING_BOTTOM int = 3

	MIN_SCREEN_WIDTH  int = 32
	MIN_SCREEN_HEIGHT int = 9
)

func NormalizedIndex(posX, posY, w, h int) int {
	x := 0
	y := 0
	if posX < 0 {
		posX++
		x = (w - 1) - ((posX * -1) % w)
	} else {
		x = posX % w
	}
	if posY < 0 {
		posY++
		y = (h - 1) - ((posY * -1) % h)
	} else {
		y = posY % h
	}
	return y*w + x
}

func HowToMessage() string {
	message := "\r\n [ UP-DOWN:W/↑  S/↓ ]   "
	message += "\n\r [ LEFT-RIGHT:A/←  D/→ ]"
	message += "\r\n [ RESTART: R QUIT: Q]\r"
	return message
}

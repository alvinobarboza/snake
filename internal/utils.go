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

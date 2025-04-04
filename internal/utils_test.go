package internal

import "testing"

func TestXxx(t *testing.T) {
	xw, yh := 5, 5

	offset := -10

	for x := range xw * 4 {
		t.Error("w", x+offset,
			"mod", (x+offset*-1)%xw,
			"i:", NormalizedIndex(x+offset, 0, xw, yh))
	}
}

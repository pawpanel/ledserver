package effects

import (
	"image/color"
	"time"
)

type Region interface {
	Count() int
	SetPixel(int, color.Color)
	SetAllPixels(color.Color)
	Apply() error
}

type Effect interface {
	Render(time.Duration, Region) time.Duration
}

type baseSolidEffect struct {
	c color.Color
}

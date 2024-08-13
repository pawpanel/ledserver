package effects

import (
	"image/color"
	"time"
)

type solidEffect struct {
	color color.Color
}

// Create a new
func NewSolidEffect(c color.Color) Effect {
	return &solidEffect{
		color: c,
	}
}

func (p *solidEffect) Render(
	elapsed time.Duration,
	region Region,
) (time.Duration, bool) {
	region.SetAllPixels(p.color)
	return 0, false
}

package effects

import (
	"image/color"
	"math"
	"time"
)

type pulseEffect struct {
	color  color.Color
	period time.Duration
}

func NewPulseEffect(c color.Color, p time.Duration) Effect {
	return &pulseEffect{
		color:  c,
		period: p,
	}
}

func (p *pulseEffect) Render(
	elapsed time.Duration,
	region Region,
) (time.Duration, bool) {
	var (
		period     = float64(p.period) / float64(time.Second)
		x          = float64(elapsed) / float64(time.Second)
		f          = math.Sin((2*x-period/2)*math.Pi/period)*.5 + .5
		r, g, b, _ = p.color.RGBA()
	)
	region.SetAllPixels(&color.RGBA{
		R: uint8(float64(r/256) * f),
		G: uint8(float64(g/256) * f),
		B: uint8(float64(b/256) * f),
	})
	return 0, true
}

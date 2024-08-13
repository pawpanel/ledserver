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
	x := float64(elapsed) / float64(time.Second)
	x -= float64(p.period) / 2
	x *= math.Pi
	x /= float64(p.period)
	x = math.Sin(x)
	x /= 2
	x += 0.5
	r, g, b, _ := p.color.RGBA()
	region.SetAllPixels(&color.RGBA64{
		R: uint16(float64(r) * x),
		G: uint16(float64(g) * x),
		B: uint16(float64(b) * x),
	})
	return 0, true
}

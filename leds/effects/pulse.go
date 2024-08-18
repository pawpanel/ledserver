package effects

import (
	"errors"
	"image/color"
	"math"
	"time"
)

type PulseEffect struct {
	Color  Color    `json:"color"`
	Period Duration `json:"period"`
	numSec float64
}

func (p *PulseEffect) Init(region Region) error {
	if p.Period == 0 {
		return errors.New("period cannot be 0")
	}
	p.numSec = float64(p.Period) / float64(time.Second)
	return nil
}

func (p *PulseEffect) Render(
	elapsed time.Duration,
	region Region,
) (time.Duration, bool) {
	var (
		x          = float64(elapsed) / float64(time.Second)
		f          = math.Sin((2*x-p.numSec/2)*math.Pi/p.numSec)*.5 + .5
		r, g, b, _ = p.Color.RGBA()
	)
	region.SetAllPixels(&color.RGBA{
		R: uint8(float64(r/256) * f),
		G: uint8(float64(g/256) * f),
		B: uint8(float64(b/256) * f),
	})
	return 0, true
}

package effects

import (
	"errors"
	"time"

	"github.com/crazy3lf/colorconv"
)

type TransitionEffect struct {
	FromColor  Color    `json:"from_color"`
	ToColor    Color    `json:"to_color"`
	Duration   Duration `json:"duration"`
	hS, sS, vS float64
	hD, sD, vD float64
}

func (t *TransitionEffect) Init(region Region) error {
	if t.Duration == 0 {
		return errors.New("duration cannot be 0")
	}
	t.hS, t.sS, t.vS = colorconv.ColorToHSV(t.FromColor)
	hE, sE, vE := colorconv.ColorToHSL(t.ToColor)
	t.hD = hE - t.hS
	t.sD = sE - t.sS
	t.vD = vE - t.vS
	return nil
}

func (t *TransitionEffect) Render(
	elapsed time.Duration,
	region Region,
) (time.Duration, bool) {
	var (
		f    = float64(elapsed) / float64(t.Duration)
		cont = f < 1
	)
	if !cont {
		f = 1
	}
	c, _ := colorconv.HSVToColor(
		t.hS+f*t.hD,
		t.sS+f*t.sD,
		t.vS+f*t.vD,
	)
	region.SetAllPixels(c)
	return 0, cont
}

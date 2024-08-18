package effects

import (
	"errors"
	"image/color"
	"time"
)

type TransitionEffect struct {
	FromColor  Color    `json:"from_color"`
	ToColor    Color    `json:"to_color"`
	Duration   Duration `json:"duration"`
	rS, gS, bS uint16
	rD, gD, bD int
}

func (t *TransitionEffect) Init(region Region) error {
	if t.Duration == 0 {
		return errors.New("duration cannot be 0")
	}
	rS, gS, bS, _ := t.FromColor.RGBA()
	rE, gE, bE, _ := t.ToColor.RGBA()
	t.rS, t.gS, t.bS = uint16(rS), uint16(gS), uint16(bS)
	t.rD = int(rE) - int(t.rS)
	t.gD = int(gE) - int(t.gS)
	t.bD = int(bE) - int(t.bS)
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
	region.SetAllPixels(color.RGBA64{
		R: uint16(float64(t.rS) + f*float64(t.rD)),
		G: uint16(float64(t.gS) + f*float64(t.gD)),
		B: uint16(float64(t.bS) + f*float64(t.bD)),
	})
	return 0, cont
}

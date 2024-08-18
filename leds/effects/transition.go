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
	rD, gD, bD uint16
}

func (t *TransitionEffect) Init(region Region) error {
	if t.Duration == 0 {
		return errors.New("duration cannot be 0")
	}
	rS, gS, bS, _ := t.FromColor.RGBA()
	rE, gE, bE, _ := t.ToColor.RGBA()
	t.rS, t.gS, t.bS = uint16(rS), uint16(gS), uint16(bS)
	t.rD = uint16(rE) - t.rS
	t.gD = uint16(gE) - t.gS
	t.bD = uint16(bE) - t.bS
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
		R: t.rS + uint16(f*float64(t.rD)),
		G: t.gS + uint16(f*float64(t.gD)),
		B: t.bS + uint16(f*float64(t.bD)),
	})
	return 0, cont
}

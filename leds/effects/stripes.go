package effects

import (
	"errors"
	"image/color"
	"math"
	"time"
)

const oneThird = float64(1) / float64(3)

type StripesEffect struct {
	Colors     []Color  `json:"colors"`
	Width      int      `json:"width"`
	Blur       float64  `json:"blur"`
	Period     Duration `json:"duration"`
	totalWidth int
}

func (s *StripesEffect) Init(region Region) error {
	if len(s.Colors) == 0 {
		return errors.New("at least one color must be specified")
	}
	if s.Width < 1 {
		return errors.New("width must be at least one")
	}
	if s.Blur < 0 || s.Blur > 1 {
		return errors.New("blur must be between 0 and 1")
	}
	if s.Period == 0 {
		return errors.New("duration cannot be 0")
	}
	s.totalWidth = len(s.Colors) * s.Width
	return nil
}

func (s *StripesEffect) Render(
	elapsed time.Duration,
	region Region,
) (time.Duration, bool) {
	var (
		fOffset = math.Mod(
			float64(elapsed),
			float64(s.Period),
		) / float64(s.Period)
		l = len(s.Colors)
	)
	for i := 0; i < region.Count(); i++ {
		var (
			offset = fOffset + float64(i%s.totalWidth)/
				float64(s.totalWidth)
			o, f    = math.Modf(offset * float64(l))
			oInt    = int(o)
			bOffset = oInt - 1
			aOffset = oInt + 1
		)
		if bOffset < 0 {
			bOffset = l - 1
		}
		if aOffset >= l {
			aOffset = 0
		}
		var (
			bR, bG, bB, _ = s.Colors[bOffset].RGBA()
			aR, aG, aB, _ = s.Colors[aOffset].RGBA()
			cR, cG, cB, _ = s.Colors[oInt].RGBA()
			e             = f*float64(2) - float64(1)
			bFactor       = s.Blur * oneThird * math.Abs(min(e, 0))
			aFactor       = s.Blur * oneThird * max(e, 0)
			cFactor       = 1 - bFactor - aFactor
		)
		region.SetPixel(i, color.RGBA64{
			R: uint16(bFactor*float64(bR) + cFactor*float64(cR) + aFactor*float64(aR)),
			G: uint16(bFactor*float64(bG) + cFactor*float64(cG) + aFactor*float64(aG)),
			B: uint16(bFactor*float64(bB) + cFactor*float64(cB) + aFactor*float64(aB)),
		})
	}
	return 0, true
}

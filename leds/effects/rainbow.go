package effects

import (
	"errors"
	"math"
	"time"

	"github.com/crazy3lf/colorconv"
)

type RainbowEffect struct {
	Width  int      `json:"width"`
	Period Duration `json:"period"`
}

func (r *RainbowEffect) Init(region Region) error {
	if r.Period == 0 {
		return errors.New("period cannot be 0")
	}
	return nil
}

func (r *RainbowEffect) Render(
	elapsed time.Duration,
	region Region,
) (time.Duration, bool) {
	for i := 0; i < region.Count(); i++ {
		o := math.Mod(
			float64(i)/float64(r.Width)+
				float64(elapsed)/float64(r.Period),
			1,
		)
		c, _ := colorconv.HSVToColor(
			math.Mod(o*360, 360),
			1,
			1,
		)
		region.SetPixel(i, c)
	}
	return 0, true
}

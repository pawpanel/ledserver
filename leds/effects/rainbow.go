package effects

import (
	"math"
	"time"

	"github.com/crazy3lf/colorconv"
)

type rainbowEffect struct {
	width  int
	period time.Duration
}

func NewRainbowEffect(w int, p time.Duration) Effect {
	return &rainbowEffect{
		width:  w,
		period: p,
	}
}

func (r *rainbowEffect) Render(
	elapsed time.Duration,
	region Region,
) (time.Duration, bool) {
	for i := 0; i < region.Count(); i++ {
		o := math.Mod(
			float64(i)/float64(r.width)+
				float64(elapsed)/float64(r.period),
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

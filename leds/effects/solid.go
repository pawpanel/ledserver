package effects

import (
	"time"
)

type SolidEffect struct {
	Color Color `json:"color"`
}

func (s *SolidEffect) Render(
	elapsed time.Duration,
	region Region,
) (time.Duration, bool) {
	region.SetAllPixels(s.Color)
	return 0, false
}

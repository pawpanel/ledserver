package effects

import (
	"image/color"
	"time"
)

type Region interface {
	Count() int
	SetPixel(int, color.Color)
	SetAllPixels(color.Color)
	Apply() error
}

type Effect interface {

	// Render accepts a duration representing the amount of time since the
	// effect was started and a region to apply the effect to. The return
	// values hint when the next render call should take place. A value of 0
	// for the duration indicates "whenever" and allows the caller to
	// determine the frame rate. If the effect can guarantee that no change
	// will occur in the specified duration, that should be specified. The
	// second return value indicates if the effect should continue or is
	// complete (no further change).
	Render(time.Duration, Region) (time.Duration, bool)
}

type EffectInit interface {

	// Init provides the effect a chance to initialize itself. This provides
	// the opportunity to validate input data and return an error if there is
	// an issue.
	Init(Region) error
}

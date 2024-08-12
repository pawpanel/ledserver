package effects

import (
	"time"
)

type pulseEffect struct {
	baseSolidEffect
}

func NewPulseEffect() {
	//...
}

func (p *pulseEffect) Render(elapsed time.Duration) time.Duration {
	return 0
}

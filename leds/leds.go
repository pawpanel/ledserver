package leds

import (
	"errors"
	"time"

	"github.com/Jon-Bright/ledctl/pixarray"
	"github.com/pawplace/ledserver/leds/effects"
	"github.com/rs/zerolog"
)

type ledRequest struct {
	region string
	effect effects.Effect
}

// Leds provides access to a ledstrip through named regions.
type Leds struct {
	frameInterval time.Duration
	regionMap     map[string]*ledRegion
	logger        zerolog.Logger
	cmdReqChan    chan *ledRequest
	cmdErrChan    chan error
	closedChan    chan any
}

func (l *Leds) run() {
	defer close(l.closedChan)
	for {
		var (
			now        = time.Now()
			nextRender time.Time
		)

		// Loop through all regions that do not have complete effects
		for _, r := range l.regionMap {
			if r.done {
				continue
			}
			if r.next.Before(now) {
				next, cont := r.effect.Render(
					now.Sub(r.start),
					r,
				)
				if cont {
					if next == 0 {
						next = l.frameInterval
					}
					r.next = now.Add(next)
				} else {
					r.done = true
				}
			} else if nextRender.IsZero() || r.next.Before(nextRender) {
				nextRender = r.next
			}
		}

		// Set a timer for the next interval (if applicable)
		var timerChan <-chan time.Time
		if !nextRender.IsZero() {
			timerChan = time.After(nextRender.Sub(now))
		}

		// Wait for either a command or the next interval
		select {
		case <-timerChan:
		case cmd, closed := <-l.cmdReqChan:
			if closed {
				return
			}
			l.cmdErrChan <- func() error {
				r, ok := l.regionMap[cmd.region]
				if !ok {
					return errors.New("invalid region")
				}
				r.start = now
				r.effect = cmd.effect
				return nil
			}()
		}
	}
}

// New creates a new Leds instance.
func New(cfg *Config) (*Leds, error) {

	// Initialize the ledstrip
	ledstrip, err := pixarray.NewWS281x(
		cfg.Count,
		3,
		pixarray.GRB,
		800000,
		10,
		[]int{cfg.Pin},
	)
	if err != nil {
		return nil, err
	}

	// Build the regions from cfg
	regionMap := make(map[string]*ledRegion)
	for _, r := range cfg.Regions {
		title := r.Title
		if title == "" {
			title = r.Name
		}
		region := &ledRegion{
			ledstrip: ledstrip,
			title:    title,
		}
		for _, b := range r.Blocks {
			if b.Offset < 0 || b.Offset+b.Count > cfg.Count {
				return nil, errors.New("offset+count outside range")
			}
			if b.Reverse {
				for i := b.Offset + b.Count; i > b.Offset; i-- {
					region.pixelMap = append(region.pixelMap, i-1)
				}
			} else {
				for i := b.Offset; i < b.Offset+b.Count; i++ {
					region.pixelMap = append(region.pixelMap, i)
				}
			}
		}
		regionMap[r.Name] = region
	}

	// If no refresh rate was provided, use 30Hz by default
	refreshRate := time.Duration(cfg.RefreshRate)
	if refreshRate == 0 {
		refreshRate = 30
	}

	// Create the Leds
	l := &Leds{
		frameInterval: time.Second / refreshRate,
		regionMap:     regionMap,
		cmdReqChan:    make(chan *ledRequest),
		cmdErrChan:    make(chan error),
		closedChan:    make(chan any),
	}

	// Start the goroutine
	go l.run()

	return l, nil
}

// Regions returns a map of region names to region titles.
func (l *Leds) Regions() map[string]string {
	r := make(map[string]string)
	for k, v := range l.regionMap {
		r[k] = v.title
	}
	return r
}

// Execute runs the specified effect on the specified region. This method
// must not be called after Close().
func (l *Leds) Execute(region string, effect effects.Effect) error {
	l.cmdReqChan <- &ledRequest{
		region: region,
		effect: effect,
	}
	return <-l.cmdErrChan
}

// Close shuts down the Leds instance.
func (l *Leds) Close() {
	close(l.cmdReqChan)
	<-l.closedChan
}

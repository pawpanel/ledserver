package leds

import (
	"time"

	"github.com/SimonWaldherr/ws2812/pixarray"
	"github.com/pawplace/ledserver/leds/effects"
	"github.com/rs/zerolog"
)

type Leds struct {
	ledstrip   pixarray.LEDStrip
	logger     zerolog.Logger
	effectChan chan effects.Effect
	closeChan  chan any
	closedChan chan any
}

func (l *Leds) run() {
	defer close(l.closedChan)
	var (
		start  time.Time
		effect effects.Effect
		timer  *time.Timer
	)
	for {
		var (
			now            = time.Now()
			nextRenderChan <-chan time.Time
		)
		if timer != nil {
			nextRenderChan = timer.C
		}
		select {
		case <-nextRenderChan:
			effect.Render(
				now.Sub(start),
				l.ledstrip,
			)
		case effect = <-l.effectChan:
			start = now
		case <-l.closeChan:
			return
		}
	}
}

func New(cfg *Config) (*Leds, error) {
	l, err := pixarray.NewWS281x(
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
	return &Leds{
		ledstrip:   l,
		effectChan: make(chan effects.Effect),
		closeChan:  make(chan any),
		closedChan: make(chan any),
	}, nil
}

func (l *Leds) Close() {
	close(l.closeChan)
	<-l.closedChan
}

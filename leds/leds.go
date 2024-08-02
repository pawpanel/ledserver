package leds

import (
	"github.com/SimonWaldherr/ws2812/pixarray"
	"github.com/rs/zerolog"
)

type Leds struct {
	ledstrip pixarray.LEDStrip
	logger   zerolog.Logger
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
		ledstrip: l,
	}, nil
}

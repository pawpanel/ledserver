package effects

import (
	"errors"
	"image/color"
	"time"
)

type ChaseEffect struct {
	Color    Color `json:"color"`
	Width    int   `json:"width"`
	Speed    int   `json:"speed"`
	interval time.Duration
	values   []color.Color
}

func (c *ChaseEffect) Init(region Region) error {
	if c.Speed == 0 {
		return errors.New("speed cannot be 0")
	}
	c.interval = time.Second / time.Duration(c.Speed)
	c.values = make([]color.Color, region.Count())
	for i := range c.values {
		if i < c.Width {
			c.values[i] = c.Color
		} else {
			c.values[i] = color.RGBA{}
		}
	}
	return nil
}

func (c *ChaseEffect) Render(
	elapsed time.Duration,
	region Region,
) (time.Duration, bool) {
	last := len(c.values) - 1
	c.values = append(
		[]color.Color{c.values[last]},
		c.values[:last]...,
	)
	for i, v := range c.values {
		region.SetPixel(i, v)
	}
	return c.interval, true
}

package leds

import (
	"image/color"
	"time"

	"github.com/SimonWaldherr/ws2812/pixarray"
	"github.com/pawplace/ledserver/leds/effects"
)

type ledRegion struct {
	ledstrip pixarray.LEDStrip
	title    string
	pixelMap []int
	start    time.Time
	next     time.Time
	done     bool
	effect   effects.Effect
}

func colorToPixel(c color.Color) pixarray.Pixel {
	r, g, b, _ := c.RGBA()
	return pixarray.Pixel{
		R: int(r / 256),
		G: int(g / 256),
		B: int(b / 256),
	}
}

func (r *ledRegion) Count() (count int) {
	return len(r.pixelMap)
}

func (r *ledRegion) SetPixel(i int, c color.Color) {
	if i >= 0 && i < len(r.pixelMap) {
		r.ledstrip.SetPixel(
			r.pixelMap[i],
			colorToPixel(c),
		)
	}
}

func (r *ledRegion) SetAllPixels(c color.Color) {
	pixelColor := colorToPixel(c)
	for _, v := range r.pixelMap {
		r.ledstrip.SetPixel(
			v,
			pixelColor,
		)
	}
}

func (r *ledRegion) Apply() error {
	return r.ledstrip.Write()
}

package effects

import (
	"encoding/json"
	"image/color"

	"github.com/crazy3lf/colorconv"
)

type Color struct {
	color.Color
}

func (c *Color) UnmarshalJSON(b []byte) error {
	var strColor string
	if err := json.Unmarshal(b, &strColor); err != nil {
		return err
	}
	v, err := colorconv.HexToColor(strColor)
	if err != nil {
		return err
	}
	c.Color = v
	return nil
}

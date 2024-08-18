package server

import (
	"fmt"
	"image/color"
	"net/http"

	"github.com/crazy3lf/colorconv"
	"github.com/gin-gonic/gin"
	"github.com/pawplace/ledserver/leds/effects"
)

func parseColor(v string) color.Color {
	c, err := colorconv.HexToColor(v)
	if err != nil {
		panic(err)
	}
	return c
}

func (s *Server) apiGetRegions(c *gin.Context) {
	c.JSON(http.StatusOK, s.leds.Regions())
}

func (s *Server) apiPostRegions(c *gin.Context) {
	var (
		regionName = c.Param("name")
		effectName = c.Param("effect")
		effect     effects.Effect
	)
	switch effectName {
	case "chase":
		effect = &effects.ChaseEffect{}
	case "pulse":
		effect = &effects.PulseEffect{}
	case "rainbow":
		effect = &effects.RainbowEffect{}
	case "solid":
		effect = &effects.SolidEffect{}
	case "transition":
		effect = &effects.TransitionEffect{}
	default:
		panic(fmt.Sprintf("invalid effect \"%s\"", effectName))
	}
	if err := c.ShouldBindJSON(effect); err != nil {
		panic(err)
	}
	if err := s.leds.Execute(regionName, effect); err != nil {
		panic(err)
	}
	c.Status(http.StatusNoContent)
}

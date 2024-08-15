package server

import (
	"fmt"
	"image/color"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/icza/gox/imagex/colorx"
	"github.com/pawplace/ledserver/leds/effects"
)

func parseColor(v string) color.Color {
	c, err := colorx.ParseHexColor(v)
	if err != nil {
		panic(err)
	}
	return c
}

func (s *Server) apiGetRegions(c *gin.Context) {
	c.JSON(http.StatusOK, s.leds.Regions())
}

type apiPostRegionsSolidParams struct {
	Color string `json:"color"`
}

type apiPostRegionsPulseParams struct {
	Color  string         `json:"color"`
	Period stringDuration `json:"period"`
}

func (s *Server) apiPostRegions(c *gin.Context) {
	var (
		regionName = c.Param("name")
		effectName = c.Param("effect")
		effect     effects.Effect
	)
	switch effectName {
	case "solid":
		v := &apiPostRegionsSolidParams{}
		if err := c.ShouldBindJSON(v); err != nil {
			panic(err)
		}
		effect = effects.NewSolidEffect(parseColor(v.Color))
	case "pulse":
		v := &apiPostRegionsPulseParams{}
		if err := c.ShouldBindJSON(v); err != nil {
			panic(err)
		}
		effect = effects.NewPulseEffect(
			parseColor(v.Color),
			time.Duration(v.Period),
		)
	default:
		panic(fmt.Sprintf("invalid effect \"%s\"", effectName))
	}
	if err := s.leds.Execute(regionName, effect); err != nil {
		panic(err)
	}
	c.Status(http.StatusNoContent)
}

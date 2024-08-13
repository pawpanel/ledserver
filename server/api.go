package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) api_get_regions(c *gin.Context) {
	c.JSON(http.StatusOK, s.leds.Regions())
}

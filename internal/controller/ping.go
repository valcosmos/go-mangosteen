package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping is the /ping router handler
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

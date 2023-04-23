package router

import (
	"mangosteen/internal/controller"

	"github.com/gin-gonic/gin"
)

// New it returns router
func New() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", controller.Ping)

	return r
}

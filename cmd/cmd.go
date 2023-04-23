package cmd

import "mangosteen/internal/router"

// RunServer is the function of run server
func RunServer() {
	r := router.New()
	r.Run()
}

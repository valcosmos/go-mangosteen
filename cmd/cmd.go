package cmd

import "mangosteen/internal/router"

func RunServer() {
	r := router.New()
	r.Run()
}

package cmd

import (
	"mangosteen/internal/database"
	"mangosteen/internal/router"
)

// RunServer is the function of run server
func RunServer() {
	database.Connect()
	database.CreateTables()
	r := router.New()
	r.Run()
}

package cmd

import (
	"fmt"
	"mangosteen/internal/router"

	"github.com/spf13/cobra"
)

func Run() {
	rootCmd := &cobra.Command{
		Use: "mangosteen",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Print("hi")
		},
	}
	rootCmd.Execute()
	// database.Connect()
	// database.CreateTables()
	// RunServer()
}

func RunDatabase() {}

// RunServer is the function of run server
func RunServer() {

	r := router.New()
	r.Run()
}

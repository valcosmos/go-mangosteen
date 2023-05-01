package cmd

import (
	"mangosteen/internal/database"
	"mangosteen/internal/router"

	"github.com/spf13/cobra"
)

func Run() {
	rootCmd := &cobra.Command{
		Use: "mangosteen",
		// Run: func(cmd *cobra.Command, args []string) {
		// 	fmt.Print("hi")
		// },
	}
	srvCmd := &cobra.Command{
		Use: "server",
		Run: func(cmd *cobra.Command, args []string) {
			RunServer()
		},
	}

	dbCmd := &cobra.Command{
		Use: "db",
		Run: func(cmd *cobra.Command, args []string) {
			database.Connect()
			// RunServer()
		},
	}

	createCmd := &cobra.Command{
		Use: "create",
		Run: func(cmd *cobra.Command, args []string) {
			database.CreateTables()

		},
	}

	rootCmd.AddCommand(srvCmd)
	rootCmd.AddCommand(dbCmd)
	dbCmd.AddCommand(createCmd)
	rootCmd.Execute()

}

func RunDatabase() {}

// RunServer is the function of run server
func RunServer() {
	r := router.New()
	r.Run()
}

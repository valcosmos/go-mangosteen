package cmd

import (
	"mangosteen/internal/database"
	"mangosteen/internal/router"

	"github.com/spf13/cobra"
)

func Run() {
	rootCmd := &cobra.Command{
		Use: "mangosteen",
	}
	srvCmd := &cobra.Command{
		Use: "server",
		Run: func(cmd *cobra.Command, args []string) {
			RunServer()
		},
	}

	dbCmd := &cobra.Command{
		Use: "db",
	}

	createCmd := &cobra.Command{
		Use: "create",
		Run: func(cmd *cobra.Command, args []string) {
			database.CreateTables()

		},
	}

	mgrtCmd := &cobra.Command{
		Use: "migrate",
		Run: func(cmd *cobra.Command, args []string) {
			database.Migrate()
		},
	}

	crudCmd := &cobra.Command{
		Use: "curd",
		Run: func(cmd *cobra.Command, args []string) {
			database.Crud()
		},
	}

	database.Connect()
	rootCmd.AddCommand(srvCmd, dbCmd)
	dbCmd.AddCommand(createCmd, mgrtCmd, crudCmd)
	defer database.Close()
	rootCmd.Execute()
}

func RunDatabase() {}

// RunServer is the function of run server
func RunServer() {
	r := router.New()
	r.Run()
}

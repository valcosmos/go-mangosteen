package cmd

import (
	"mangosteen/internal/database"
	"mangosteen/internal/email"
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

	// createCmd := &cobra.Command{
	// 	Use: "create",
	// 	Run: func(cmd *cobra.Command, args []string) {
	// 		database.CreateTables()

	// 	},
	// }

	createMigrationCmd := &cobra.Command{
		Use: "create:migration",
		Run: func(cmd *cobra.Command, args []string) {
			database.CreateMigration(args[0])
		},
	}

	mgrtCmd := &cobra.Command{
		Use: "migrate",
		Run: func(cmd *cobra.Command, args []string) {
			database.Migrate()
		},
	}

	mgrtDownCmd := &cobra.Command{
		Use: "migrate:down",
		Run: func(cmd *cobra.Command, args []string) {
			database.MigrateDown()
		},
	}

	crudCmd := &cobra.Command{
		Use: "curd",
		Run: func(cmd *cobra.Command, args []string) {
			database.Crud()
		},
	}

	emailCmd := &cobra.Command{
		Use: "email",
		Run: func(cmd *cobra.Command, args []string) {
			email.Send()
		},
	}

	database.Connect()
	rootCmd.AddCommand(srvCmd, dbCmd, emailCmd)
	dbCmd.AddCommand(mgrtCmd, crudCmd, createMigrationCmd, mgrtDownCmd)
	defer database.Close()
	rootCmd.Execute()
}

func RunDatabase() {}

// RunServer is the function of run server
func RunServer() {
	r := router.New()
	r.Run()
}

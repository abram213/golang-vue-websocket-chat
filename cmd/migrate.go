package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"chat/app"
	"chat/model"
)

var migrateCmd = &cobra.Command{
	Use: "migrate",
	RunE: func(cmd *cobra.Command, args []string) error {
		refresh, _ := cmd.Flags().GetBool("refresh")

		a, err := app.New()
		if err != nil {
			return err
		}
		defer a.Close()

		if refresh {
			fmt.Println("droping tables...")
			a.Database.DropTables(
				&model.User{},
				&model.Friendship{},
				&model.Tokens{})
		}

		a.Database.Migrate(
			&model.User{},
			&model.Friendship{},
			&model.Tokens{})
		fmt.Println("migration success!")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
	migrateCmd.Flags().Bool("refresh", false, "drop all tables before migration")
}

package cmd

import (
	"github.com/krobus00/analytics-service/internal/bootstrap"
	"github.com/spf13/cobra"
)

// migrationCmd represents the migration command.
var migrationCmd = &cobra.Command{
	Use:   "migration",
	Short: "database migration",
	Long:  `database migration`,
	Run: func(cmd *cobra.Command, args []string) {
		action, _ := cmd.Flags().GetString("action")
		migrationName, _ := cmd.Flags().GetString("name")
		step, _ := cmd.Flags().GetInt64("step")

		bootstrap.StartMigration(action, migrationName, &step)
	},
}

func init() {
	rootCmd.AddCommand(migrationCmd)
	migrationCmd.PersistentFlags().String("action", "up", "action create|up|up-by-one|up-to|down|down-to|reset|status")
	migrationCmd.PersistentFlags().String("step", "1", "step")
	migrationCmd.PersistentFlags().String("name", "", "migration name")
}

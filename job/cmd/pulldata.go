package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(pulldataCmd)
}

var pulldataCmd = &cobra.Command{
	Use:   "pull-data",
	Short: "pull data for real time box-office",
	Long:  `pull data for real time box-office`,
	Run: func(cmd *cobra.Command, args []string) {
		// todo

	},
}

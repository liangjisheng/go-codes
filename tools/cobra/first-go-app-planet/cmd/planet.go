package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dagobah",
	Short: "Dagobah is an awesome planet style RSS aggregator",
	Long: `Dagobah provides planet style RSS aggregation. It
is inspired by python planet. It has a simple YAML configuration
and provides it's own webserver.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Dagobah runs")
	},
}

// Execute ...
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

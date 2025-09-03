package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var fiberCmd = &cobra.Command{
	Use:   "fiber",
	Short: "Generates fiber boilerplate in path",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("fiber called")
	},
}

func init() {
	rootCmd.AddCommand(fiberCmd)
}

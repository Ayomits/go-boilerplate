package cmd

import (
	"github.com/Ayomits/go-boilerplate/pkg/generator"

	"github.com/spf13/cobra"
)

var fiberCmd = &cobra.Command{
	Use:   "fiber",
	Short: "Генерирует fiber проект по указанному пути",
	Run: func(cmd *cobra.Command, args []string) {
		generator.NewFiberProjectGenerator().Generate()
	},
	Aliases: []string{
		"f",
		"fi",
		"fib",
	},
}

func init() {
	rootCmd.AddCommand(fiberCmd)
}

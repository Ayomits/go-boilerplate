package cmd

import (
	"github.com/Ayomits/go-boilerplate/pkg/generator"

	"github.com/spf13/cobra"
)

var ginCmd = &cobra.Command{
	Use:   "gin",
	Short: "Генерирует gin проект по указанному пути",
	Run: func(cmd *cobra.Command, args []string) {
		generator.NewGinProjectGenerator().Generate()
	},
	Aliases: []string{
		"g",
	},
}

func init() {
	rootCmd.AddCommand(ginCmd)
}

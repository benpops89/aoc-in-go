package aoc

import (
	"aoc.benpoppy.dev/pkg/aoc"
	"github.com/spf13/cobra"
)

var setupCmd = &cobra.Command{
	Use:   "setup [year] [day]",
	Short: "Setup Advent of Code puzzle for given year and day",
	Run: func(cmd *cobra.Command, args []string) {
		aoc.Setup(args[0], args[1])
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)
}

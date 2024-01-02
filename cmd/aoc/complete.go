package aoc

import (
	"aoc.benpoppy.dev/pkg/aoc"
	"github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
	Use:   "complete [year] [day]",
	Short: "Archive completed Advent of Code puzzle for year and day",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		aoc.Complete(args[0], args[1])
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}

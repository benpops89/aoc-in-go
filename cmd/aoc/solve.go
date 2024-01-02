package aoc

import (
	"aoc.benpoppy.dev/pkg/aoc"
	"github.com/spf13/cobra"
)

var test bool
var filename string

var solveCmd = &cobra.Command{
	Use:   "solve",
	Short: "Solve Advent of Code puzzle",
	Run: func(cmd *cobra.Command, args []string) {
		if !test {
			filename = "input"
		} else {
			filename = "sample"
		}
		aoc.Solve(filename)
	},
}

func init() {
	solveCmd.Flags().BoolVarP(&test, "test", "t", false, "Run solution against sample")
	rootCmd.AddCommand(solveCmd)
}

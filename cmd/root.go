package cmd

import (
	"aoc/cmd/year2023"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "aoc",
	Short: "Advent of Code CLI",
	Long:  `A command line tool for solving Advent of Code problems across multiple years.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	// Add year-specific commands here
	rootCmd.AddCommand(year2023.Cmd)
}

package day1

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day1",
	Short: "day1",
	Long:  "day1",
	Run: func(cmd *cobra.Command, args []string) {
		execute(cmd.Parent().Name(), cmd.Name())
	},
}

func execute(parent, command string) {
	b, err := os.ReadFile(fmt.Sprintf("cmd/year%s/%s/1.txt", parent, command))
	if err != nil {
		logrus.Fatalf("error reading input: %v", err)
	}

	logrus.Infof("score part1: %d", part1(string(b)))
	// logrus.Infof("score part2: %d", part2(string(b)))
}

func readInputToArrays(input string) ([]int, []int, error) {
	// Initialize the arrays
	var left []int
	var right []int

	for _, line := range strings.Split(input, "\n") {
		// Split the line into two numbers
		parts := strings.Fields(line)
		if len(parts) != 2 {
			return nil, nil, fmt.Errorf("invalid line format: %s", line)
		}

		// Convert strings to integers
		leftNum, err1 := strconv.Atoi(parts[0])
		rightNum, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			return nil, nil, fmt.Errorf("failed to parse numbers in line: %s", line)
		}

		// Append numbers to arrays
		left = append(left, leftNum)
		right = append(right, rightNum)
	}
	return left, right, nil
}

func part1(s string) int64 {
	var score int64
	var left, right, err = readInputToArrays(s)
	if err != nil {
		fmt.Println("Error:", err)
		return score
	}
	fmt.Println(left)
	fmt.Println(right)

	return score
}

func part2(s string) int64 {
	return 0
}

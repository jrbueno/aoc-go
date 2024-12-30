/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"jrbueno/aoc-go/internal"
	"jrbueno/aoc-go/solutions/year2024/day1"
	"jrbueno/aoc-go/solutions/year2024/day2"
	"jrbueno/aoc-go/solutions/year2024/day3"
	"jrbueno/aoc-go/solutions/year2024/day4"
	"log"
	"os"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run a solution for a given day",
	Long:  `run a solution for a given day.`,
	Run: func(cmd *cobra.Command, args []string) {
		year, _ := cmd.Flags().GetInt("year")
		//log.Println(year)
		day, _ := cmd.Flags().GetInt("day")
		//log.Println(day)
		if day < 1 || day > 25 {
			log.Printf("Invalid day %d\n", day)
			return
		}
		var ds *internal.DaySolution
		switch day {
		case 1:
			ds = internal.NewDaySolution(year, day, day1.RunAll, day1.RunPartOne, day1.RunPartTwo)
		case 2:
			ds = internal.NewDaySolution(year, day, day2.RunAll, day2.RunPartOne, day2.RunPartTwo)
		case 3:
			ds = internal.NewDaySolution(year, day, day3.RunAll, day3.RunPartOne, day3.RunPartTwo)
		case 4:
			ds = internal.NewDaySolution(year, day, day4.RunAll, day4.RunPartOne, day4.RunPartTwo)
		default:
			fmt.Printf("No solution found for year %v day %v\n", year, day)
			return
		}
		partOneInput := getInputFile(year, day, 1)
		partTwoInput := getInputFile(year, day, 2)
		partOneResult, partTwoResult := ds.RunAll(year, day, partOneInput, partTwoInput)
		fmt.Printf("Part Two Result: %v\n", partTwoResult)
		fmt.Printf("Part One Result: %v\n", partOneResult)
	},
}

func getInputFile(year int, day int, part int) []string {
	fileName := fmt.Sprintf("inputs/%d/day%d-%d.txt", year, day, part)
	if _, err := os.Stat(fileName); errors.Is(err, os.ErrNotExist) {
		log.Printf("Input file %s does not exist", fileName)
		return nil
	}
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Error opening file %s: %v", fileName, err)
		panic(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file %s: %v", fileName, err)
	}
	return lines
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	runCmd.Flags().IntP("year", "y", 2024, "Year of the challenge")
	runCmd.Flags().IntP("day", "d", 0, "Day of the challenge")
	if err := runCmd.MarkFlagRequired("day"); err != nil {
		log.Fatalf("Error marking flag required: %v", err)
	}
}

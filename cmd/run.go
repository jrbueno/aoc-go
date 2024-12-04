/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"jrbueno/aoc-go/internal"
	day1 "jrbueno/aoc-go/solutions/year2024"
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
			ds = internal.NewDaySolution(year, day, day1.RunDayOneAll, day1.RunPartOne, day1.RunPartTwo)
		default:
			fmt.Printf("No solution found for year %v day %v\n", year, day)
			return
		}
		inputFile := getInputFile(year, day)
		defer inputFile.Close()
		partOneResult, partTwoResult := ds.RunAll(inputFile, year, day)
		fmt.Printf("Part Two Result: %v\n", partTwoResult)
		fmt.Printf("Part One Result: %v\n", partOneResult)
	},
}

func getInputFile(year int, day int) *os.File {
	fileName := fmt.Sprintf("inputs/%d/day%d.txt", year, day)
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Error opening file %s: %v", fileName, err)
		panic(err)
	}
	return file
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

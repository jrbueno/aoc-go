/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	day1 "jrbueno/aoc-go/internal/year2024"
	"log"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run a solution for a given day",
	Long:  `run a solution for a given day.`,
	Run: func(cmd *cobra.Command, args []string) {
		year, _ := cmd.Flags().GetInt("year")
		log.Println(year)
		day, _ := cmd.Flags().GetInt("day")
		log.Println(day)
		if day < 1 || day > 25 {
			log.Printf("Invalid day %d\n", day)
			return
		}
		result := day1.RunSolution(year, day)
		log.Println(result)
	},
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
	runCmd.MarkFlagRequired("day")
}

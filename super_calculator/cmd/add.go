/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"calculator/internals/math"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	x, y float64
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Addition of x and y",
	Long:  `Addition of x and y`,
	Run: func(cmd *cobra.Command, args []string) {
		ret := math.Addition(x, y)
		fmt.Printf("%f + %f = %f\n", x, y, ret)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().Float64VarP(&x, "x", "x", 0, "first value")
	addCmd.Flags().Float64VarP(&y, "y", "y", 0, "second value")
}

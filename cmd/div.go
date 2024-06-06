/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"calculator/internals/math"
	"fmt"

	"github.com/spf13/cobra"
)

// divCmd represents the div command
var divCmd = &cobra.Command{
	Use:   "div",
	Short: "Division of x and y",
	Long:  `Addition of x and y`,
	Run: func(cmd *cobra.Command, args []string) {
		ret := math.Division(x, y)
		fmt.Printf("%f / %f = %f\n", x, y, ret)
	},
}

func init() {
	rootCmd.AddCommand(divCmd)
	divCmd.Flags().Float64VarP(&x, "x", "x", 0, "first value")
	divCmd.Flags().Float64VarP(&y, "y", "y", 0, "second value")
}

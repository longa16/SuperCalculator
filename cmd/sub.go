/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"calculator/internals/math"
	"fmt"

	"github.com/spf13/cobra"
)

// subCmd represents the sub command
var subCmd = &cobra.Command{
	Use:   "sub",
	Short: "Subtraction of x and y",
	Long:  `Subtraction of x and y`,
	Run: func(cmd *cobra.Command, args []string) {
		ret := math.Subtraction(x, y)
		fmt.Printf("%f - %f = %f\n", x, y, ret)
	},
}

func init() {
	rootCmd.AddCommand(subCmd)
	subCmd.Flags().Float64VarP(&x, "x", "x", 0, "first value")
	subCmd.Flags().Float64VarP(&y, "y", "y", 0, "second value")
}

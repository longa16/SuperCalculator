/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"calculator/internals/math"
	"fmt"

	"github.com/spf13/cobra"
)

// mulCmd represents the mul command
var mulCmd = &cobra.Command{
	Use:   "mul",
	Short: "Multiplication of x and y",
	Long:  `Multiplication of x and y`,
	Run: func(cmd *cobra.Command, args []string) {
		ret := math.Multiplication(x, y)
		fmt.Printf("%f * %f = %f\n", x, y, ret)
	},
}

func init() {
	rootCmd.AddCommand(mulCmd)
	mulCmd.Flags().Float64VarP(&x, "x", "x", 0, "first value")
	mulCmd.Flags().Float64VarP(&y, "y", "y", 0, "second value")
}

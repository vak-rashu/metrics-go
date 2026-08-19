/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	metrics "github.com/vak-rashu/metrics-go/pkg"
)

// processCmd represents the cpu command
var processCmd = &cobra.Command{
	Use:   "process",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		list, processStat, err := metrics.ShowPerProcessData(args[0])
		if err != nil {
			return err
		}

		fmt.Println(list, processStat)

		return nil
	},
}

func init() {}

/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	metrics "github.com/vak-rashu/metrics-go/pkg"
)

// logicalCpuCmd represents the logicalCpu command
var logicalCpuCmd = &cobra.Command{
	Use:   "logicalCpu",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cpuCount, logicalCpu, err := metrics.CountCPU()
		if err != nil {
			return err
		}

		cmd.Printf("Total number of CPU are %d\nTotal number of Logical CPU are %d\n", cpuCount, logicalCpu)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(logicalCpuCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// logicalCpuCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// logicalCpuCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

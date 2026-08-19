/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	metrics "github.com/vak-rashu/metrics-go/pkg"
)

var (
	isAll      string
	logicalCpu string
)

// cpuCmd represents the cpu command
var cpuCmd = &cobra.Command{
	Use:   "cpu",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		showAll, err := cmd.Flags().GetBool("all")
		getLogical, err := cmd.Flags().GetString("logical")
		if err != nil {
			return err
		}

		if showAll {
			err := metrics.ShowPerCpuStat()
			if err != nil {
				return err
			}
		} else if getLogical == "true" {
			logicalCpu, err := metrics.LogicalCpuCount()
			if err != nil {
				return err
			}
			cmd.Printf("Total number of Logical CPU are %d\n", logicalCpu)
		} else {
			systemCpuStat, err := metrics.ShowCPUstat()
			if err != nil {
				return err
			}
			cmd.Println("System wide CPU metrics:\n", systemCpuStat)
		}

		return nil
	},
}

func init() {
	cpuCmd.PersistentFlags().StringVarP(&isAll, "all", "a", "false", "Show all cpu data")
	cpuCmd.PersistentFlags().StringVarP(&logicalCpu, "logical", "l", "false", "Toggle off logical cpu")

}

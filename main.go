/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"

	metrics "github.com/vak-rashu/metrics-go/pkg"
)

func main() {
	// cmd.Execute()
	// tui.StartTui()
	// fmt.Println(metrics.ShowCPUstat())
	// tui.CreateChart()
	if perc, err := metrics.CalculateCPUStat(); err != nil {
		panic(err)
	} else {
		fmt.Println(perc)
	}

}

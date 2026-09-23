/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"os"

	tui "github.com/vak-rashu/metrics-go/ui"
)

func main() {
	// cmd.Execute()

	cmd := os.Args
	if cmd[1] == "show" {
		tui.CreateChart()
	}
}

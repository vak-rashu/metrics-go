package tui

import (
	"fmt"
	"os"

	"github.com/NimbleMarkets/go-booba"
	"github.com/NimbleMarkets/ntcharts/v2/sparkline"
)

func CreateChart() {
	width := 25
	height := 12

	m := model{
		cpu: sparkline.New(
			width,
			height/4,
			sparkline.WithMaxValue(100.0),
			sparkline.WithStyle(cpuStyle),
		),

		diskRead: sparkline.New(
			width,
			height/4,
			sparkline.WithStyle(diskReadStyle),
		),

		diskWrite: sparkline.New(
			width,
			height/4,
			sparkline.WithStyle(diskWriteStyle),
		),

		netRX: sparkline.New(
			width,
			height/4,
			sparkline.WithStyle(netRXStyle),
		),

		netTX: sparkline.New(
			width,
			height/4,
			sparkline.WithStyle(netTXStyle),
		),
	}

	if err := booba.Run(m); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

// func CreateChart() {
// 	width := 25
// 	height := 12

// 	m := model{
// 		s5: sparkline.New(width, height/4, sparkline.WithMaxValue(100.0), sparkline.WithStyle(blockStyle4))}

// 	if err := booba.Run(m); err != nil {
// 		fmt.Println("Error running program:", err)
// 		os.Exit(1)
// 	}
// }

// old function for cli

// func StartTui() {
// p := tea.NewProgram(
// 	stat{msg: "METRICS"},
// )

// if _, err := p.Run(); err != nil {
// 	panic(err)
// }
// }

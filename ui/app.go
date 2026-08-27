package tui

import (
	"fmt"
	"os"

	"github.com/NimbleMarkets/go-booba"
	"github.com/NimbleMarkets/ntcharts/v2/sparkline"
)

// metrics "github.com/vak-rashu/metrics-go/pkg"

func StartTui() {
	// p := tea.NewProgram(
	// 	stat{msg: "METRICS"},
	// )

	// if _, err := p.Run(); err != nil {
	// 	panic(err)
	// }
}

// func CreateChart() {

// 	width := 30
// 	height := 12
// 	chart := tslc.New(width, height)

// 	zoneManager := zone.New()
// 	chart.SetZoneManager(zoneManager)
// 	chart.Focus()

// 	perc, err := metrics.CalculateCPUStat()
// 	if err != nil {
// 		panic(err)
// 	}

// 	dataSet := []float64{perc}
// 	for _, v := range dataSet {
// 		// date := time.Now().Add(time.Hour * time.Duration(24*i))
// 		chart.Push(tslc.TimePoint{Value: v})
// 	}

// 	// dataSet := []float64{0, 2, 4, 6, 8, 10, 8, 6, 4, 2, 0}
// 	// for i, v := range dataSet {
// 	// 	date := time.Now().Add(time.Hour * time.Duration(24*i))
// 	// 	chart.Push(tslc.TimePoint{Time: date, Value: v})
// 	// }

// 	chart.SetDataSetStyle("dataset2",
// 		gloss.NewStyle().
// 			Foreground(gloss.Color("10")))

// 	m := model{chart, zoneManager}
// 	if err := booba.Run(m); err != nil {
// 		fmt.Println("Error running program:", err)
// 		os.Exit(1)
// 	}
// }

func CreateChart() {
	width := 25
	height := 12
	max := 100.0
	// all sparklines contain the same values,
	// but will be scaled when displayed based on sparkline height
	// sparkline1 calls Draw with no background style
	// sparkline2 calls DrawColumnsOnly with background style
	// sparkline3 calls Draw with background style (same style as sparkline2)
	// sparkline4 calls Draw with background style
	// sparkline5 calls DrawBraille with no background style

	m := model{
		// sparkline.New(width, height, sparkline.WithMaxValue(max), sparkline.WithStyle(blockStyle)),
		// sparkline.New(width, (height/2)-1, sparkline.WithMaxValue(max), sparkline.WithStyle(blockStyle2)),
		// sparkline.New(width, (height/2)-1, sparkline.WithMaxValue(max), sparkline.WithStyle(blockStyle2)),
		// sparkline.New(width, height/4, sparkline.WithMaxValue(max), sparkline.WithStyle(blockStyle3)),
		sparkline.New(width, height/4, sparkline.WithMaxValue(max), sparkline.WithStyle(blockStyle4)),
		max}
	// booba.Run is a tea.Program substitute that dispatches to native Bubble Tea
	// or the WASM/ghostty-web bridge depending on build target.
	// See https://github.com/NimbleMarkets/go-booba-example.
	if err := booba.Run(m); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

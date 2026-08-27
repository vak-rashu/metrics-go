package tui

import (
	"fmt"
	"os"
	"time"

	gloss "charm.land/lipgloss/v2"
	booba "github.com/NimbleMarkets/go-booba"
	tslc "github.com/NimbleMarkets/ntcharts/v2/linechart/timeserieslinechart"
	zone "github.com/lrstanley/bubblezone/v2"
	metrics "github.com/vak-rashu/metrics-go/pkg"
	// metrics "github.com/vak-rashu/metrics-go/pkg"
)

func StartTui() {
	// p := tea.NewProgram(
	// 	stat{msg: "METRICS"},
	// )

	// if _, err := p.Run(); err != nil {
	// 	panic(err)
	// }
}

func CreateChart() {

	width := 30
	height := 12
	chart := tslc.New(width, height)

	zoneManager := zone.New()
	chart.SetZoneManager(zoneManager)
	chart.Focus()

	perc, err := metrics.CalculateCPUStat()
	if err != nil {
		panic(err)
	}

	dataSet := []float64{perc}
	for i, v := range dataSet {
		date := time.Now().Add(time.Hour * time.Duration(24*i))
		chart.PushDataSet("dataset2", tslc.TimePoint{Time: date, Value: v})
	}

	// dataSet := []float64{0, 2, 4, 6, 8, 10, 8, 6, 4, 2, 0}
	// for i, v := range dataSet {
	// 	date := time.Now().Add(time.Hour * time.Duration(24*i))
	// 	chart.Push(tslc.TimePoint{Time: date, Value: v})
	// }

	chart.SetDataSetStyle("dataset2",
		gloss.NewStyle().
			Foreground(gloss.Color("10")))

	for {
		m := model{chart, zoneManager}
		if err := booba.Run(m); err != nil {
			fmt.Println("Error running program:", err)
			os.Exit(1)
		}
	}
}

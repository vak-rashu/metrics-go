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

	cpu, _ := metrics.ShowCPUstat()
	// if err != nil {
	// 	return err
	// }

	dataSet := []float64{cpu.UserTime, cpu.NiceTime, cpu.SystemTime}
	for i, v := range dataSet {
		date := time.Now().Add(time.Hour * time.Duration(24*i))
		chart.PushDataSet("dataset2", tslc.TimePoint{Time: date, Value: v})
	}

	chart.SetDataSetStyle("dataset2",
		gloss.NewStyle().
			Foreground(gloss.Color("10")))

	m := model{chart, zoneManager}
	if err := booba.Run(m); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

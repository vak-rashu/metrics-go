package tui

import (
	"fmt"
	"os"
	"time"

	tea "charm.land/bubbletea/v2"
	gloss "charm.land/lipgloss/v2"
	booba "github.com/NimbleMarkets/go-booba"
	tslc "github.com/NimbleMarkets/ntcharts/v2/linechart/timeserieslinechart"
	zone "github.com/lrstanley/bubblezone/v2"
	metrics "github.com/vak-rashu/metrics-go/pkg"
)

type model struct {
	chart       tslc.Model
	zoneManager *zone.Manager
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}

	m.chart, _ = m.chart.Update(msg)
	m.chart.DrawBrailleAll()
	return m, nil
}

func (m model) View() tea.View {
	v := tea.NewView(m.zoneManager.Scan(
		gloss.NewStyle().
			BorderStyle(gloss.NormalBorder()).
			BorderForeground(gloss.Color("63")).
			Render(m.chart.View()),
	))

	v.AltScreen = true
	v.MouseMode = tea.MouseModeCellMotion

	return v
}

func createChart() {

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

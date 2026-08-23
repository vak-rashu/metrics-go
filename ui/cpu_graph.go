package tui

import (
	tea "charm.land/bubbletea/v2"
	gloss "charm.land/lipgloss/v2"
	tslc "github.com/NimbleMarkets/ntcharts/v2/linechart/timeserieslinechart"
	zone "github.com/lrstanley/bubblezone/v2"
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

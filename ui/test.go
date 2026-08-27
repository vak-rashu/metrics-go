// ntcharts - Copyright (c) 2024 Neomantra Corp.

package tui

import (
	"fmt"
	"time"

	"github.com/NimbleMarkets/ntcharts/v2/sparkline"
	metrics "github.com/vak-rashu/metrics-go/pkg"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

var perc float64

type tickMsg time.Time

var defaultStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("63")) // purple

// var titleStyle = lipgloss.NewStyle().
// 	Foreground(lipgloss.Color("3")) // yellow

// var blockStyle = lipgloss.NewStyle().
// 	Foreground(lipgloss.Color("63")) // purple

// var blockStyle2 = lipgloss.NewStyle().
// 	Foreground(lipgloss.Color("9")). // red
// 	Background(lipgloss.Color("2"))  // green

// var blockStyle3 = lipgloss.NewStyle().
// 	Foreground(lipgloss.Color("6")). // cyan
// 	Background(lipgloss.Color("3"))  // yellow

var blockStyle4 = lipgloss.NewStyle().
	Foreground(lipgloss.Color("3")) // yellow

type model struct {
	s5  sparkline.Model
	max float64
}

func doTick() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (m model) Init() tea.Cmd {
	return doTick()
}

func getStat() tea.Cmd {
	perc, err := metrics.CalculateCPUStat()
	if err != nil {
		panic(err)
	}

	return tea.Tick(1*time.Second, func(time.Time) tea.Msg {
		return perc
	})

}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	case tickMsg:
		return m, getStat()
	case float64:
		m.s5.Push(perc)

		// call different Draw functions with different Style combinations
		m.s5.DrawBraille()

	}

	// add same random value to all sparkline

	return m, getStat()
}

func (m model) View() tea.View {
	s := "press any button to push the same random value to all sparklines, `q/ctrl+c` to quit\n"
	s += lipgloss.JoinHorizontal(lipgloss.Top,
		// defaultStyle.Render("Draw() w/o background\n"+m.s1.View()),
		// defaultStyle.Render(titleStyle.Render("style w/ background")+"\nDrawColumnsOnly()\n"+m.s2.View()+"\nDraw()\n"+m.s3.View()),
		lipgloss.JoinVertical(lipgloss.Left,
			defaultStyle.Render(fmt.Sprintf("Max: %.0f, Random: %.2f", m.max, perc)),
			defaultStyle.Render("\nDrawBraille()\n"+m.s5.View()),
		),
	) + "\n"
	return tea.NewView(s)
}

// "Draw() w/ background\n"+m.s4.View()+

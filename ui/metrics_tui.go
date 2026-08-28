// cmd: metrics tui

package tui

import (
	"fmt"
	"time"

	"github.com/NimbleMarkets/ntcharts/v2/sparkline"
	metrics "github.com/vak-rashu/metrics-go/metrics"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type tickMsg time.Time

var defaultStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("63")) // purple

var blockStyle4 = lipgloss.NewStyle().
	Foreground(lipgloss.Color("3")) // yellow

type model struct {
	s5            sparkline.Model
	activeCPUTime float64
}

func doTick() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (m model) Init() tea.Cmd {
	return nil
}

func getStat() float64 {
	_, _, perc, err := metrics.CalculateCPUStat()
	if err != nil {
		panic(err)
	}
	return perc
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	case tickMsg:
		m.activeCPUTime = getStat()
		m.s5.Push(m.activeCPUTime)
		m.s5.DrawBraille()
	}
	return m, doTick()
}

func (m model) View() tea.View {
	s := "press any button to push the same random value to all sparklines, `q/ctrl+c` to quit\n"
	s += lipgloss.JoinHorizontal(lipgloss.Top,
		lipgloss.JoinVertical(lipgloss.Left,
			defaultStyle.Render(fmt.Sprintf("CPU Active Time: %.9f", m.activeCPUTime)),
			defaultStyle.Render("\nDrawBraille()\n"+m.s5.View()),
		),
	) + "\n"

	return tea.NewView(s)
}

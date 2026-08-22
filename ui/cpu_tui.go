package tui

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	metrics "github.com/vak-rashu/metrics-go/pkg"
)

// stat model follows interface Model
type stat struct {
	msg string
	val metrics.CPUStat
}

func (c stat) Init() tea.Cmd {
	// return func() tea.Msg {
	// 	gloss.Println(headLineStyle.Render(c.msg))
	// 	return nil
	// }

	return nil
}

func (c stat) View() tea.View {
	return tea.NewView(fmt.Sprintf("Values are: %v", c.val))
}

func (c stat) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.(tea.KeyMsg).String() {
		case "m":
			return c, getCPUStat()
		case "q":
			return c, tea.Quit
		}
	case metrics.CPUStat:
		return c, nil
	}
	return c, nil
}

func getCPUStat() tea.Cmd {
	return func() tea.Msg {
		cpu := metrics.ShowPerCpuStat()
		return cpu
	}
}

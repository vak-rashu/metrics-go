package tui

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	gloss "charm.land/lipgloss/v2"
	metrics "github.com/vak-rashu/metrics-go/pkg"
)

// stat model follows interface Model
type stat struct {
	msg string
	val metrics.CPUStat
	err error
}

func (c stat) Init() tea.Cmd {
	c.msg = "Metrics"
	return func() tea.Msg {
		gloss.Println(headLineStyle.Render(c.msg))
		return nil
	}
}

func (c stat) View() tea.View {
	if c.err != nil {
		return tea.NewView(fmt.Sprint(c.err))
	} else {
		return tea.NewView(fmt.Sprintf("Values are: %v", c.val))
	}
}

func (c stat) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "m":
			return c, getCPUStat()
		case "q":
			return c, tea.Quit
		}
	case metrics.CPUStat:
		c.val = msg
		return c, nil
	case error:
		c.err = msg
		return c, nil
	}
	return c, nil
}

func getCPUStat() tea.Cmd {
	return func() tea.Msg {
		if cpu, err := metrics.ShowCPUstat(); err != nil {
			return err
		} else {
			return cpu
		}
	}
}

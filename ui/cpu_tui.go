package tui

import (
	tea "charm.land/bubbletea/v2"
	gloss "charm.land/lipgloss/v2"
	metrics "github.com/vak-rashu/metrics-go/pkg"
)

// model data
type simplePage struct {
	msg string
}

var blocks = []string{
	"CPU",
	"Processes",
	"Memory",
}

func cpuTUI(msg string) simplePage {
	return simplePage{msg: msg}
}

func (s simplePage) Init() tea.Cmd {
	return func() tea.Msg {
		gloss.Println(headLineStyle.Render(s.msg))
		return nil
	}
}

func (s simplePage) View() string {

	listString := []string{}
	for _, val := range blocks {
		listString = append(listString, tabStyle.Render(val))
	}
	return gloss.JoinHorizontal(gloss.Bottom, listString...)

}

func (s simplePage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tea.KeyMsg:
		switch msg.(tea.KeyMsg).String() {
		case "m":
			return s, getCPUStat()
		case "enter":
			return s, tea.Quit
		}
	}

	return s, nil
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

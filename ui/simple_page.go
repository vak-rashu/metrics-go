package tui

import (
	gloss "charm.land/lipgloss/v2"
	tea "github.com/charmbracelet/bubbletea"
	metrics "github.com/vak-rashu/metrics-go/pkg"
)

var headLineStyle = gloss.NewStyle().
	Width(100).
	Align(gloss.Center).
	Background(gloss.Color("3")).
	Foreground(gloss.Color("12"))

var tabStyle = gloss.NewStyle().
	Border(gloss.RoundedBorder()).
	PaddingLeft(1).
	PaddingRight(1).
	BorderForeground(gloss.Color("183"))

// model data
type simplePage struct {
	msg string
}

var blocks = []string{
	"CPU",
	"Processes",
	"Memory",
}

func newSimplePage(msg string) simplePage {
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

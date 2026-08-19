package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

// model data
type simplePage struct {
	msg string
}

func newSimplePage(msg string) simplePage {
	return simplePage{msg: msg}
}

func (s simplePage) Init() tea.Cmd {
	return nil
}

func (s simplePage) View() string {
	return s.msg
}

func (s simplePage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tea.KeyMsg:
		switch msg.(tea.KeyMsg).String() {
		case "q":
			return s, tea.Quit
		}
	}

	return s, nil
}

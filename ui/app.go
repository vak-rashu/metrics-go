package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

func StartTui() {
	p := tea.NewProgram(
		newSimplePage("METRICS: Streamline your systems monitoring"),
	)

	if _, err := p.Run(); err != nil {
		panic(err)
	}

}

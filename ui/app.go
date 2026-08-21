package tui

import (
	tea "charm.land/bubbletea/v2"
)

func StartTui() {
	p := tea.NewProgram(
		cpuTUI("METRICS: Streamline your systems monitoring"),
	)

	if _, err := p.Run(); err != nil {
		panic(err)
	}

}

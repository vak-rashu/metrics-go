package tui

import tea "charm.land/bubbletea/v2"

func StartTui() {
	p := tea.NewProgram(
		stat{},
	)

	if _, err := p.Run(); err != nil {
		panic(err)
	}

}

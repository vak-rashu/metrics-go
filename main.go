/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package main

func main() {
	// cmd.Execute()
	// tui.CreateChart()
	// CreateMemoryChart()
}

// var memBoxStyle = lipgloss.NewStyle().
// 	BorderStyle(lipgloss.NormalBorder()).
// 	BorderForeground(lipgloss.Color("63"))

// type memoryModel struct {
// 	memTotal     int
// 	memFree      int
// 	memAvailable int
// }

// func (m memoryModel) Init() tea.Cmd {
// 	return nil
// }

// func (m memoryModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
// 	switch msg := msg.(type) {

// 	case tea.KeyMsg:
// 		switch msg.String() {
// 		case "q", "ctrl+c":
// 			return m, tea.Quit
// 		}

// 	}

// 	return m, nil
// }

// func (m memoryModel) View() tea.View {
// 	s := fmt.Sprintf(
// 		"Memory\n\n"+
// 			"MemTotal:     %d gB\n"+
// 			"MemFree:      %d gB\n"+
// 			"MemAvailable: %d gB",
// 		m.memTotal,
// 		m.memFree,
// 		m.memAvailable,
// 	)

// 	return tea.NewView(
// 		memBoxStyle.Render(s),
// 	)
// }

// func CreateMemoryChart() {
// 	total, free, available := metrics.GetMemStats()

// 	m := memoryModel{
// 		memTotal:     total,
// 		memFree:      free,
// 		memAvailable: available,
// 	}

// 	if _, err := tea.NewProgram(m).Run(); err != nil {
// 		fmt.Println("Error running memory TUI:", err)
// 		os.Exit(1)
// 	}
// }

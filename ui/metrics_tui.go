// // cmd: metrics tui

// package tui

// import (
// 	"fmt"
// 	"time"

// 	"github.com/NimbleMarkets/ntcharts/v2/sparkline"
// 	metrics "github.com/vak-rashu/metrics-go/metrics"

// 	tea "charm.land/bubbletea/v2"
// 	"charm.land/lipgloss/v2"
// )

// type tickMsg time.Time

// var defaultStyle = lipgloss.NewStyle().
// 	BorderStyle(lipgloss.NormalBorder()).
// 	BorderForeground(lipgloss.Color("63")) // purple

// var blockStyle4 = lipgloss.NewStyle().
// 	Foreground(lipgloss.Color("3")) // yellow

// type model struct {
// 	sec  int
// 	s5   sparkline.Model
// 	perc float64
// }

// func doTick() tea.Cmd {
// 	return tea.Tick(1*time.Second, func(t time.Time) tea.Msg {
// 		return tickMsg(t)
// 	})
// }

// func (m model) Init() tea.Cmd {
// 	return nil
// }

// func getStat() float64 {
// 	perc, err := metrics.CalculateCPUStatPerc()
// 	if err != nil {
// 		panic(err)
// 	}
// 	return perc
// }

// func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
// 	switch msg := msg.(type) {
// 	case tea.KeyMsg:
// 		switch msg.String() {
// 		case "q", "ctrl+c":
// 			return m, tea.Quit
// 		}
// 	case tickMsg:
// 		m.sec++
// 		m.perc = getStat()
// 		m.s5.Push(m.perc)
// 		m.s5.DrawBraille()
// 	}
// 	return m, doTick()
// }

// func (m model) View() tea.View {
// 	s := "press q/ctrl+c` to quit\n"
// 	s += lipgloss.JoinHorizontal(lipgloss.Top,
// 		lipgloss.JoinVertical(lipgloss.Left,
// 			defaultStyle.Render(fmt.Sprintf("CPU Active Time: %f", m.perc)),
// 			defaultStyle.Render("\nCPU\n"+m.s5.View()),
// 		),
// 	) + "\n"

// 	return tea.NewView(s)
// }

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
	BorderForeground(lipgloss.Color("63"))

var blockStyle4 = lipgloss.NewStyle().
	Foreground(lipgloss.Color("3")) // yellow

type model struct {
	cpu sparkline.Model

	// Disk
	diskRead  sparkline.Model
	diskWrite sparkline.Model

	// Network
	netRX sparkline.Model
	netTX sparkline.Model

	// Current values
	cpuPerc float64

	diskReadIOPS  float64
	diskWriteIOPS float64

	netRXPackets float64
	netTXPackets float64

	// Memory
	memTotal     float64
	memFree      float64
	memAvailable float64
}

func doTick() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (m model) Init() tea.Cmd {
	return doTick()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}

	case tickMsg:

		// ---------------- CPU ----------------
		cpuPerc, err := metrics.CalculateCPUStatPerc()
		if err != nil {
			fmt.Println("CPU:", err)
		} else {
			m.cpuPerc = cpuPerc

			m.cpu.Push(cpuPerc)
			m.cpu.DrawBraille()
		}

		// ---------------- DISK READ ----------------
		readIOPS, err := metrics.GetDiskReadIOPS()
		if err != nil {
			fmt.Println("Disk read:", err)
		} else {
			m.diskReadIOPS = readIOPS

			m.diskRead.Push(readIOPS)
			m.diskRead.DrawBraille()
		}

		// ---------------- DISK WRITE ----------------
		writeIOPS, err := metrics.GetDiskWriteIOPS()
		if err != nil {
			fmt.Println("Disk write:", err)
		} else {
			m.diskWriteIOPS = writeIOPS

			m.diskWrite.Push(writeIOPS)
			m.diskWrite.DrawBraille()
		}

		// ---------------- NETWORK ----------------
		rxPackets, txPackets, err := metrics.GetPacketsStat()
		if err != nil {
			fmt.Println("Network:", err)
		} else {
			m.netRXPackets = rxPackets
			m.netTXPackets = txPackets

			m.netRX.Push(rxPackets)
			m.netTX.Push(txPackets)

			m.netRX.DrawBraille()
			m.netTX.DrawBraille()
		}
	}

	return m, doTick()
}

func (m model) View() tea.View {
	s := ""

	// CPU
	s += defaultStyle.Render(
		fmt.Sprintf(
			"CPU Active: %.2f%%\n\nCPU\n%s",
			m.cpuPerc,
			m.cpu.View(),
		),
	)

	s += "\n\n"

	// DISK
	s += defaultStyle.Render(
		fmt.Sprintf(
			"Disk I/O\n"+
				"Read IOPS:  %.0f\n"+
				"Write IOPS: %.0f\n\n"+
				"Read\n%s\n\n"+
				"Write\n%s",
			m.diskReadIOPS,
			m.diskWriteIOPS,
			m.diskRead.View(),
			m.diskWrite.View(),
		),
	)

	s += "\n\n"

	// NETWORK
	s += defaultStyle.Render(
		fmt.Sprintf(
			"Network Packet Rate\n"+
				"RX: %.0f packets/s\n"+
				"TX: %.0f packets/s\n\n"+
				"RX\n%s\n\n"+
				"TX\n%s",
			m.netRXPackets,
			m.netTXPackets,
			m.netRX.View(),
			m.netTX.View(),
		),
	)

	s += "\n\n"

	// MEMORY
	s += defaultStyle.Render(
		fmt.Sprintf(
			"Memory\n"+
				"Total: %.2f MB\n"+
				"Free: %.2f MB\n"+
				"Available: %.2f MB",
			m.memTotal,
			m.memFree,
			m.memAvailable,
		),
	)

	return tea.NewView(s)
}

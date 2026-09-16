// cmd: metrics tui

package tui

import (
	"fmt"
	"time"

	"github.com/NimbleMarkets/ntcharts/v2/sparkline"
	metrics "github.com/vak-rashu/metrics-go/metrics"

	tea "charm.land/bubbletea/v2"
	gloss "charm.land/lipgloss/v2"
)

type tickMsg time.Time

var fetchFrequency = 1 * time.Second

type model struct {
	cpu sparkline.Model

	diskRead  sparkline.Model
	diskWrite sparkline.Model

	netRX sparkline.Model
	netTX sparkline.Model

	cpuPerc float64

	diskReadIOPS  float64
	diskWriteIOPS float64

	netRXPackets float64
	netTXPackets float64

	memTotal     int
	memFree      int
	memAvailable int
}

func NewModel() model {
	const width = 40

	return model{
		cpu: sparkline.New(width, 3, sparkline.WithStyle(cpuStyle)),

		diskRead:  sparkline.New(width, 2, sparkline.WithStyle(diskReadStyle)),
		diskWrite: sparkline.New(width, 2, sparkline.WithStyle(diskWriteStyle)),

		netRX: sparkline.New(width, 2, sparkline.WithStyle(netRXStyle)),
		netTX: sparkline.New(width, 2, sparkline.WithStyle(netTXStyle)),
	}
}

func doTick() tea.Cmd {
	return tea.Tick(fetchFrequency, func(t time.Time) tea.Msg {
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

		// ---------------- MEMORY ----------------
		t, f, a, err := metrics.GetMemStats()
		if err != nil {
			fmt.Println("Memory:", err)
		} else {
			m.memTotal = t
			m.memFree = f
			m.memAvailable = a
		}

		return m, doTick()
	}
	return m, nil
}

func (m model) View() tea.View {

	const panelWidth = 48

	// ---------------- CPU ----------------
	cpuPanel := defaultStyle.Width(panelWidth).Render(
		fmt.Sprintf(
			"CPU\n"+
				"Active: %.2f%%\n\n"+
				"%s",
			m.cpuPerc,
			m.cpu.View(),
		),
	)

	// ---------------- MEMORY ----------------
	memPanel := defaultStyle.Width(panelWidth).Render(
		fmt.Sprintf(
			"Memory\n\n"+
				"Total:     %d GB\n"+
				"Free:      %d GB\n"+
				"Available: %d GB",
			m.memTotal,
			m.memFree,
			m.memAvailable,
		),
	)

	// ---------------- DISK ----------------
	diskPanel := defaultStyle.Width(panelWidth).Render(
		fmt.Sprintf(
			"Disk I/O\n\n"+
				"Read IOPS:  %.0f\n"+
				"Write IOPS: %.0f\n\n"+
				"Read\n"+
				"%s\n\n"+
				"Write\n"+
				"%s",
			m.diskReadIOPS,
			m.diskWriteIOPS,
			m.diskRead.View(),
			m.diskWrite.View(),
		),
	)

	// ---------------- NETWORK ----------------
	networkPanel := defaultStyle.Width(panelWidth).Render(
		fmt.Sprintf(
			"Network\n\n"+
				"RX: %.0f packets/s\n"+
				"TX: %.0f packets/s\n\n"+
				"RX\n"+
				"%s\n\n"+
				"TX\n"+
				"%s",
			m.netRXPackets,
			m.netTXPackets,
			m.netRX.View(),
			m.netTX.View(),
		),
	)

	// ---------------- 2 × 2 LAYOUT ----------------
	topRow := gloss.JoinHorizontal(
		gloss.Top,
		cpuPanel,
		memPanel,
	)

	bottomRow := gloss.JoinHorizontal(
		gloss.Top,
		diskPanel,
		networkPanel,
	)

	// Put the two rows together.
	dashboard := gloss.JoinVertical(
		gloss.Left,
		topRow,
		"",
		bottomRow,
	)

	return tea.NewView(dashboard)
}

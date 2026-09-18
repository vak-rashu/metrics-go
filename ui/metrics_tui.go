// cmd: metrics tui

// package tui

// import (
// 	"fmt"
// 	"time"

// 	"github.com/NimbleMarkets/ntcharts/v2/sparkline"
// 	metrics "github.com/vak-rashu/metrics-go/metrics"

// 	tea "charm.land/bubbletea/v2"
// 	gloss "charm.land/lipgloss/v2"
// )

// type tickMsg time.Time

// var fetchFrequency = 1 * time.Second

// const (
// 	smallWidth  = 28
// 	smallHeight = 5

// 	detailWidth  = 70
// 	detailHeight = 15
// )

// type metricType int

// const (
// 	cpuMetric metricType = iota
// 	memoryMetric
// 	diskMetric
// 	networkMetric
// )

// func (m model) metricBox(
// 	title string,
// 	graph string,
// 	metric metricType,
// ) string {

// 	style := defaultStyle.
// 		Width(smallWidth).
// 		Height(smallHeight)

// 	// Highlight selected metric
// 	if m.selected == metric {
// 		style = style.Border(gloss.ThickBorder())
// 	} else {
// 		style = style.Border(gloss.NormalBorder())
// 	}

// 	return style.Render(
// 		fmt.Sprintf(
// 			"%s\n\n%s",
// 			title,
// 			graph,
// 		),
// 	)
// }

// type model struct {
// 	// Small graphs
// 	cpu       sparkline.Model
// 	memory    sparkline.Model
// 	diskRead  sparkline.Model
// 	diskWrite sparkline.Model
// 	netRX     sparkline.Model
// 	netTX     sparkline.Model

// 	// Large/detail graphs
// 	cpuDetail    sparkline.Model
// 	memoryDetail sparkline.Model
// 	diskDetail   sparkline.Model
// 	netDetail    sparkline.Model

// 	// Current values
// 	cpuPerc float64

// 	memoryPerc   float64
// 	memTotal     int
// 	memFree      int
// 	memAvailable int

// 	diskReadIOPS  float64
// 	diskWriteIOPS float64

// 	netRXPackets float64
// 	netTXPackets float64

// 	// Currently selected metric
// 	selected metricType
// }

// func NewModel() model {
// 	return model{
// 		// Small graphs
// 		cpu: sparkline.New(
// 			smallWidth,
// 			2,
// 			sparkline.WithStyle(cpuStyle),
// 		),

// 		memory: sparkline.New(
// 			smallWidth,
// 			2,
// 			sparkline.WithStyle(cpuStyle),
// 		),

// 		diskRead: sparkline.New(
// 			smallWidth,
// 			1,
// 			sparkline.WithStyle(diskReadStyle),
// 		),

// 		diskWrite: sparkline.New(
// 			smallWidth,
// 			1,
// 			sparkline.WithStyle(diskWriteStyle),
// 		),

// 		netRX: sparkline.New(
// 			smallWidth,
// 			1,
// 			sparkline.WithStyle(netRXStyle),
// 		),

// 		netTX: sparkline.New(
// 			smallWidth,
// 			1,
// 			sparkline.WithStyle(netTXStyle),
// 		),

// 		// Large graphs
// 		cpuDetail: sparkline.New(
// 			detailWidth,
// 			detailHeight,
// 			sparkline.WithStyle(cpuStyle),
// 		),

// 		memoryDetail: sparkline.New(
// 			detailWidth,
// 			detailHeight,
// 			sparkline.WithStyle(cpuStyle),
// 		),

// 		diskDetail: sparkline.New(
// 			detailWidth,
// 			detailHeight,
// 			sparkline.WithStyle(diskReadStyle),
// 		),

// 		netDetail: sparkline.New(
// 			detailWidth,
// 			detailHeight,
// 			sparkline.WithStyle(netRXStyle),
// 		),

// 		// CPU selected by default
// 		selected: cpuMetric,
// 	}
// }

// func doTick() tea.Cmd {
// 	return tea.Tick(fetchFrequency, func(t time.Time) tea.Msg {
// 		return tickMsg(t)
// 	})
// }

// func (m model) Init() tea.Cmd {
// 	return doTick()
// }

// func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

// 	switch msg := msg.(type) {

// 	case tea.KeyMsg:

// 		switch msg.String() {

// 		case "q", "ctrl+c":
// 			return m, tea.Quit

// 		case "up", "k":
// 			if m.selected > cpuMetric {
// 				m.selected--
// 			}

// 		case "down", "j":
// 			if m.selected < networkMetric {
// 				m.selected++
// 			}
// 		}

// 	case tickMsg:

// 		// ---------------- CPU ----------------

// 		cpuPerc, err := metrics.CalculateCPUStatPerc()
// 		if err != nil {
// 			fmt.Println("CPU:", err)
// 		} else {
// 			m.cpuPerc = cpuPerc

// 			// Small graph
// 			m.cpu.Push(cpuPerc)
// 			m.cpu.DrawBraille()

// 			// Large graph
// 			m.cpuDetail.Push(cpuPerc)
// 			m.cpuDetail.DrawBraille()
// 		}

// 		// ---------------- MEMORY ----------------

// 		total, free, available, err := metrics.GetMemStats()

// 		if err != nil {
// 			fmt.Println("Memory:", err)
// 		} else {
// 			m.memTotal = total
// 			m.memFree = free
// 			m.memAvailable = available

// 			if total > 0 {
// 				m.memoryPerc =
// 					float64(total-available) /
// 						float64(total) *
// 						100

// 				// Small graph
// 				m.memory.Push(m.memoryPerc)
// 				m.memory.DrawBraille()

// 				// Large graph
// 				m.memoryDetail.Push(m.memoryPerc)
// 				m.memoryDetail.DrawBraille()
// 			}
// 		}

// 		// ---------------- DISK READ ----------------

// 		readIOPS, err := metrics.GetDiskReadIOPS()

// 		if err != nil {
// 			fmt.Println("Disk read:", err)
// 		} else {
// 			m.diskReadIOPS = readIOPS

// 			// Small graph
// 			m.diskRead.Push(readIOPS)
// 			m.diskRead.DrawBraille()

// 			// Large graph
// 			m.diskDetail.Push(readIOPS)
// 			m.diskDetail.DrawBraille()
// 		}

// 		// ---------------- DISK WRITE ----------------

// 		writeIOPS, err := metrics.GetDiskWriteIOPS()

// 		if err != nil {
// 			fmt.Println("Disk write:", err)
// 		} else {
// 			m.diskWriteIOPS = writeIOPS

// 			m.diskWrite.Push(writeIOPS)
// 			m.diskWrite.DrawBraille()
// 		}

// 		// ---------------- NETWORK ----------------

// 		rxPackets, txPackets, err := metrics.GetPacketsStat()

// 		if err != nil {
// 			fmt.Println("Network:", err)
// 		} else {
// 			m.netRXPackets = rxPackets
// 			m.netTXPackets = txPackets

// 			// Small graphs
// 			m.netRX.Push(rxPackets)
// 			m.netRX.DrawBraille()

// 			m.netTX.Push(txPackets)
// 			m.netTX.DrawBraille()

// 			// Large graph
// 			m.netDetail.Push(rxPackets)
// 			m.netDetail.DrawBraille()
// 		}

// 		return m, doTick()
// 	}

// 	return m, nil
// }

// func (m model) View() tea.View {

// 	// ---------------- LEFT SIDEBAR ----------------

// 	cpuBox := m.metricBox(
// 		"CPU",
// 		m.cpu.View(),
// 		cpuMetric,
// 	)

// 	memoryBox := m.metricBox(
// 		"Memory",
// 		m.memory.View(),
// 		memoryMetric,
// 	)

// 	diskGraph := gloss.JoinVertical(
// 		gloss.Left,
// 		m.diskRead.View(),
// 		m.diskWrite.View(),
// 	)

// 	diskBox := m.metricBox(
// 		"Disk",
// 		diskGraph,
// 		diskMetric,
// 	)

// 	networkGraph := gloss.JoinVertical(
// 		gloss.Left,
// 		m.netRX.View(),
// 		m.netTX.View(),
// 	)

// 	networkBox := m.metricBox(
// 		"Network",
// 		networkGraph,
// 		networkMetric,
// 	)

// 	sidebar := gloss.JoinVertical(
// 		gloss.Left,
// 		cpuBox,
// 		"",
// 		memoryBox,
// 		"",
// 		diskBox,
// 		"",
// 		networkBox,
// 	)

// 	// ---------------- RIGHT DETAIL PANEL ----------------

// 	var detailGraph string
// 	var detailValue string
// 	var title string

// 	switch m.selected {

// 	case cpuMetric:
// 		title = "CPU"
// 		detailGraph = m.cpuDetail.View()
// 		detailValue = fmt.Sprintf(
// 			"Utilization: %.2f%%",
// 			m.cpuPerc,
// 		)

// 	case memoryMetric:
// 		title = "Memory"
// 		detailGraph = m.memoryDetail.View()
// 		detailValue = fmt.Sprintf(
// 			"Utilization: %.2f%%",
// 			m.memoryPerc,
// 		)

// 	case diskMetric:
// 		title = "Disk"
// 		detailGraph = m.diskDetail.View()
// 		detailValue = fmt.Sprintf(
// 			"Read IOPS: %.0f    Write IOPS: %.0f",
// 			m.diskReadIOPS,
// 			m.diskWriteIOPS,
// 		)

// 	case networkMetric:
// 		title = "Network"
// 		detailGraph = m.netDetail.View()
// 		detailValue = fmt.Sprintf(
// 			"RX: %.0f packets/s    TX: %.0f packets/s",
// 			m.netRXPackets,
// 			m.netTXPackets,
// 		)
// 	}

// 	detailPanel := defaultStyle.
// 		Width(detailWidth).
// 		Height(detailHeight + 5).
// 		Render(
// 			fmt.Sprintf(
// 				"%s\n\n%s\n\n%s",
// 				title,
// 				detailGraph,
// 				detailValue,
// 			),
// 		)

// 	// ---------------- FINAL LAYOUT ----------------

// 	layout := gloss.JoinHorizontal(
// 		gloss.Top,
// 		sidebar,
// 		"    ",
// 		detailPanel,
// 	)

// 	return tea.NewView(layout)
// }

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

// fetch frequency of the timer
var fetchFrequency = 1 * time.Second

// width of the screen
const width = 40

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

	selected map[int]struct{}
}

func NewModel() model {

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
	// have the init here
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
		total, free, available, err := metrics.GetMemStats()
		if err != nil {
			fmt.Println("Memory:", err)
		} else {
			m.memTotal = total
			m.memFree = free
			m.memAvailable = available
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

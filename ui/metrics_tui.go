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

const (
	smallWidth  = 28
	smallHeight = 5

	detailWidth  = 70
	detailHeight = 15
)

type metricType int

const (
	cpuMetric metricType = iota
	memoryMetric
	diskMetric
	networkMetric
)

func (m model) metricBox(
	title string,
	graph string,
	metric metricType,
) string {

	style := defaultStyle.
		Width(smallWidth).
		Height(smallHeight)

	// Highlight selected metric
	if m.selected == metric {
		style = style.BorderStyle(gloss.ThickBorder())
	} else {
		style = style.BorderStyle(gloss.NormalBorder())
	}

	return style.Render(
		fmt.Sprintf(
			"%s\n%s",
			title,
			graph,
		),
	)
}

type model struct {
	// Small graphs
	cpu       sparkline.Model
	memory    sparkline.Model
	diskRead  sparkline.Model
	diskWrite sparkline.Model
	netRX     sparkline.Model
	netTX     sparkline.Model

	// Large/detail graphs
	cpuDetail    sparkline.Model
	memoryDetail sparkline.Model
	diskDetail   sparkline.Model
	netDetail    sparkline.Model

	// Current values
	cpuPerc float64

	memoryPerc   float64
	memTotal     float64
	memFree      float64
	memAvailable float64

	diskReadIOPS  float64
	diskWriteIOPS float64

	netRXPackets float64
	netTXPackets float64

	// Currently selected metric
	selected metricType
}

func NewModel() model {
	innerSmallWidth := smallWidth - 2
	innerDetailWidth := detailWidth - 2

	return model{
		// Small graphs
		cpu: sparkline.New(
			innerSmallWidth,
			2,
			sparkline.WithMaxValue(100.0),
			sparkline.WithStyle(cpuStyle),
		),

		memory: sparkline.New(
			innerSmallWidth,
			2,
			sparkline.WithMaxValue(100.0),
			sparkline.WithStyle(cpuStyle),
		),

		diskRead: sparkline.New(
			innerSmallWidth,
			1,
			sparkline.WithStyle(diskReadStyle),
		),

		diskWrite: sparkline.New(
			innerSmallWidth,
			1,
			sparkline.WithStyle(diskWriteStyle),
		),

		netRX: sparkline.New(
			innerSmallWidth,
			1,
			sparkline.WithStyle(netRXStyle),
		),

		netTX: sparkline.New(
			innerSmallWidth,
			1,
			sparkline.WithStyle(netTXStyle),
		),

		// Large graphs
		cpuDetail: sparkline.New(
			innerDetailWidth,
			detailHeight,
			sparkline.WithMaxValue(100.0),
			sparkline.WithStyle(cpuStyle),
		),

		memoryDetail: sparkline.New(
			innerDetailWidth,
			detailHeight,
			sparkline.WithMaxValue(100.0),
			sparkline.WithStyle(cpuStyle),
		),

		diskDetail: sparkline.New(
			innerDetailWidth,
			detailHeight,
			sparkline.WithStyle(diskReadStyle),
		),

		netDetail: sparkline.New(
			innerDetailWidth,
			detailHeight,
			sparkline.WithStyle(netRXStyle),
		),

		// CPU selected by default
		selected: cpuMetric,
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

		case "up", "k":
			if m.selected > cpuMetric {
				m.selected--
			}

		case "down", "j":
			if m.selected < networkMetric {
				m.selected++
			}
		}

	case tea.MouseClickMsg:

		if msg.X >= 0 && msg.X < smallWidth {
			if msg.Y >= 0 && msg.Y <= 4 {
				m.selected = cpuMetric
			} else if msg.Y >= 6 && msg.Y <= 10 {
				m.selected = memoryMetric
			} else if msg.Y >= 12 && msg.Y <= 16 {
				m.selected = diskMetric
			} else if msg.Y >= 18 && msg.Y <= 22 {
				m.selected = networkMetric
			}
		}

	case tickMsg:

		// ---------------- CPU ----------------

		cpuPerc, err := metrics.CalculateCPUStatPerc()
		if err != nil {
			fmt.Println("CPU:", err)
		} else {
			m.cpuPerc = cpuPerc

			// Small graph
			m.cpu.Push(cpuPerc)
			m.cpu.DrawBraille()

			// Large graph
			m.cpuDetail.Push(cpuPerc)
			m.cpuDetail.DrawBraille()
		}

		// ---------------- MEMORY ----------------

		total, free, available, _, _, err := metrics.GetMemStats()

		if err != nil {
			fmt.Println("Memory:", err)
		} else {
			m.memTotal = total
			m.memFree = free
			m.memAvailable = available

			if total > 0 {
				m.memoryPerc =
					float64(total-available) /
						float64(total) *
						100

				// Small graph
				m.memory.Push(m.memoryPerc)
				m.memory.DrawBraille()

				// Large graph
				m.memoryDetail.Push(m.memoryPerc)
				m.memoryDetail.DrawBraille()
			}
		}

		// ---------------- DISK READ ----------------

		readIOPS, err := metrics.GetDiskReadIOPS("")

		if err != nil {
			fmt.Println("Disk read:", err)
		} else {
			m.diskReadIOPS = readIOPS

			// Small graph
			m.diskRead.Push(readIOPS)
			m.diskRead.DrawBraille()

			// Large graph
			m.diskDetail.Push(readIOPS)
			m.diskDetail.DrawBraille()
		}

		// ---------------- DISK WRITE ----------------

		writeIOPS, err := metrics.GetDiskWriteIOPS("")

		if err != nil {
			fmt.Println("Disk write:", err)
		} else {
			m.diskWriteIOPS = writeIOPS

			m.diskWrite.Push(writeIOPS)
			m.diskWrite.DrawBraille()
		}

		// ---------------- NETWORK ----------------

		rxPackets, txPackets, err := metrics.GetPacketsStat("")

		if err != nil {
			fmt.Println("Network:", err)
		} else {
			m.netRXPackets = rxPackets
			m.netTXPackets = txPackets

			// Small graphs
			m.netRX.Push(rxPackets)
			m.netRX.DrawBraille()

			m.netTX.Push(txPackets)
			m.netTX.DrawBraille()

			// Large graph
			m.netDetail.Push(rxPackets)
			m.netDetail.DrawBraille()
		}

		return m, doTick()
	}

	return m, nil
}

func (m model) View() tea.View {

	// ---------------- LEFT SIDEBAR ----------------

	cpuBox := m.metricBox(
		"CPU",
		m.cpu.View(),
		cpuMetric,
	)

	memoryBox := m.metricBox(
		"Memory",
		m.memory.View(),
		memoryMetric,
	)

	diskGraph := gloss.JoinVertical(
		gloss.Left,
		m.diskRead.View(),
		m.diskWrite.View(),
	)

	diskBox := m.metricBox(
		"Disk",
		diskGraph,
		diskMetric,
	)

	networkGraph := gloss.JoinVertical(
		gloss.Left,
		m.netRX.View(),
		m.netTX.View(),
	)

	networkBox := m.metricBox(
		"Network",
		networkGraph,
		networkMetric,
	)

	sidebar := gloss.JoinVertical(
		gloss.Left,
		cpuBox,
		"",
		memoryBox,
		"",
		diskBox,
		"",
		networkBox,
	)

	// ---------------- RIGHT DETAIL PANEL ----------------

	var detailGraph string
	var detailValue string
	var title string

	switch m.selected {

	case cpuMetric:
		title = "CPU"
		detailGraph = m.cpuDetail.View()
		detailValue = fmt.Sprintf(
			"Utilization: %.2f%%",
			m.cpuPerc,
		)

	case memoryMetric:
		title = "Memory"
		detailGraph = m.memoryDetail.View()
		detailValue = fmt.Sprintf(
			"Utilization: %.2f%%\nTotal: %.2f GB    Available: %.2f GB    Free: %.2f GB",
			m.memoryPerc,
			m.memTotal,
			m.memAvailable,
			m.memFree,
		)

	case diskMetric:
		title = "Disk"
		detailGraph = m.diskDetail.View()
		detailValue = fmt.Sprintf(
			"Read IOPS: %.0f    Write IOPS: %.0f",
			m.diskReadIOPS,
			m.diskWriteIOPS,
		)

	case networkMetric:
		title = "Network"
		detailGraph = m.netDetail.View()
		detailValue = fmt.Sprintf(
			"RX: %.0f packets/s    TX: %.0f packets/s",
			m.netRXPackets,
			m.netTXPackets,
		)
	}

	detailPanel := defaultStyle.
		Width(detailWidth).
		Height(detailHeight + 7).
		Render(
			fmt.Sprintf(
				"%s\n\n%s\n\n%s",
				title,
				detailGraph,
				detailValue,
			),
		)

	// ---------------- FINAL LAYOUT ----------------

	layout := gloss.JoinHorizontal(
		gloss.Top,
		sidebar,
		"    ",
		detailPanel,
	)

	// return tea.NewView(layout)

	v := tea.NewView(layout)
	v.AltScreen = true
	return v
}

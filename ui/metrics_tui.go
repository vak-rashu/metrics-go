// cmd: metrics tui

package tui

import (
	"fmt"
	"time"

	"github.com/NimbleMarkets/ntcharts/v2/sparkline"
	metrics "github.com/vak-rashu/metrics-go"

	tea "charm.land/bubbletea/v2"
	gloss "charm.land/lipgloss/v2"
)

type tickMsg time.Time

var fetchFrequency = 1 * time.Second

type metricType int

const (
	cpuMetric metricType = iota
	memoryMetric
	diskMetric
	networkMetric
)

type model struct {
	// Screen and layout dimensions
	width       int
	height      int
	navWidth    int
	metricWidth int
	detailWidth int
	bodyHeight  int

	// Active tab (0: Performance, 1: Processes, 2: Services)
	activeTab int

	// Small graphs
	cpu       sparkline.Model
	memory    sparkline.Model
	diskRead  sparkline.Model
	diskWrite sparkline.Model
	netRX     sparkline.Model
	netTX     sparkline.Model

	// Large/detail graphs
	cpuDetail       sparkline.Model
	memoryDetail    sparkline.Model
	diskReadDetail  sparkline.Model
	diskWriteDetail sparkline.Model
	netRXDetail     sparkline.Model
	netTXDetail     sparkline.Model

	// Current values
	cpuPerc float64

	memoryPerc   float64
	memTotal     float64
	memFree      float64
	memAvailable float64
	memCached    float64
	memUsed      float64

	diskReadIOPS  float64
	diskWriteIOPS float64

	netRXPackets float64
	netTXPackets float64

	// Currently selected metric
	metricSelected metricType

	//currently selected tab
	tabSelected int
}

func (m *model) recalculateSizes() {
	if m.width <= 0 || m.height <= 0 {
		return
	}

	// 3 lines reserved for top header badge & spacing
	m.bodyHeight = m.height - 3
	if m.bodyHeight < 12 {
		m.bodyHeight = 12
	}

	m.navWidth = 18
	m.metricWidth = 32

	m.detailWidth = m.width - m.navWidth - m.metricWidth - 4
	if m.detailWidth < 30 {
		m.detailWidth = 30
	}

	// 4 metric cards stacked in metricWidth
	cardOuterHeight := m.bodyHeight / 4
	if cardOuterHeight < 4 {
		cardOuterHeight = 4
	}

	smallW := m.metricWidth - 4
	if smallW < 4 {
		smallW = 4
	}
	smallH := cardOuterHeight - 3
	if smallH < 1 {
		smallH = 1
	}

	m.cpu.Resize(smallW, smallH)
	m.memory.Resize(smallW, smallH)

	diskH := smallH / 2
	if diskH < 1 {
		diskH = 1
	}
	m.diskRead.Resize(smallW, diskH)
	m.diskWrite.Resize(smallW, diskH)

	netH := smallH / 2
	if netH < 1 {
		netH = 1
	}
	m.netRX.Resize(smallW, netH)
	m.netTX.Resize(smallW, netH)

	// Detail Graph Box covers top half of detail panel
	graphBoxHeight := m.bodyHeight / 2
	if graphBoxHeight < 5 {
		graphBoxHeight = 5
	}

	detailGraphW := m.detailWidth - 4
	if detailGraphW < 10 {
		detailGraphW = 10
	}
	detailGraphH := graphBoxHeight - 3
	if detailGraphH < 2 {
		detailGraphH = 2
	}

	halfDetailH := detailGraphH / 2
	if halfDetailH < 1 {
		halfDetailH = 1
	}

	memQuarterH := detailGraphH / 4
	if memQuarterH < 1 {
		memQuarterH = 1
	}

	m.cpuDetail.Resize(detailGraphW, detailGraphH)
	m.memoryDetail.Resize(detailGraphW, detailGraphH)

	m.diskReadDetail.Resize(detailGraphW, halfDetailH)
	m.diskWriteDetail.Resize(detailGraphW, halfDetailH)

	m.netRXDetail.Resize(detailGraphW, halfDetailH)
	m.netTXDetail.Resize(detailGraphW, halfDetailH)
}

func (m model) metricBox(
	title string,
	graph string,
	metric metricType,
) string {
	cardOuterHeight := m.bodyHeight / 4
	if cardOuterHeight < 4 {
		cardOuterHeight = 4
	}

	var style gloss.Style
	if m.metricSelected == metric {
		style = cardSelectedStyle
	} else {
		style = cardNormalStyle
	}

	style = style.
		Width(m.metricWidth - 2).
		Height(cardOuterHeight - 2)

	return style.Render(
		fmt.Sprintf(
			"%s\n%s",
			title,
			graph,
		),
	)
}

func NewModel() model {
	initW := 100
	initH := 30
	initBodyH := initH - 3
	initNavW := 18
	initMetricW := 32
	initDetailW := initW - initNavW - initMetricW - 4

	cardH := initBodyH / 4
	smallW := initMetricW - 4
	smallH := cardH - 3
	if smallH < 1 {
		smallH = 1
	}

	detailGraphW := initDetailW - 4
	detailGraphH := (initBodyH / 2) - 3
	if detailGraphH < 2 {
		detailGraphH = 2
	}

	halfDetailH := detailGraphH / 2
	if halfDetailH < 1 {
		halfDetailH = 1
	}

	m := model{
		width:       initW,
		height:      initH,
		navWidth:    initNavW,
		metricWidth: initMetricW,
		detailWidth: initDetailW,
		bodyHeight:  initBodyH,
		activeTab:   0,

		// Small graphs
		cpu: sparkline.New(
			smallW,
			smallH,
			sparkline.WithMaxValue(100.0),
			sparkline.WithStyle(cpuStyle),
		),

		memory: sparkline.New(
			smallW,
			smallH,
			sparkline.WithMaxValue(100.0),
			sparkline.WithStyle(cpuStyle),
		),

		diskRead: sparkline.New(
			smallW,
			1,
			sparkline.WithStyle(diskReadStyle),
		),

		diskWrite: sparkline.New(
			smallW,
			1,
			sparkline.WithStyle(diskWriteStyle),
		),

		netRX: sparkline.New(
			smallW,
			1,
			sparkline.WithStyle(netRXStyle),
		),

		netTX: sparkline.New(
			smallW,
			1,
			sparkline.WithStyle(netTXStyle),
		),

		// Large/Detail graphs
		cpuDetail: sparkline.New(
			detailGraphW,
			detailGraphH,
			sparkline.WithMaxValue(100.0),
			sparkline.WithStyle(cpuStyle),
		),

		memoryDetail: sparkline.New(
			detailGraphW,
			detailGraphH,
			sparkline.WithMaxValue(100.0),
			sparkline.WithStyle(cpuStyle),
		),

		diskReadDetail: sparkline.New(
			detailGraphW,
			halfDetailH,
			sparkline.WithStyle(diskReadStyle),
		),

		diskWriteDetail: sparkline.New(
			detailGraphW,
			halfDetailH,
			sparkline.WithStyle(diskWriteStyle),
		),

		netRXDetail: sparkline.New(
			detailGraphW,
			halfDetailH,
			sparkline.WithStyle(netRXStyle),
		),

		netTXDetail: sparkline.New(
			detailGraphW,
			halfDetailH,
			sparkline.WithStyle(netTXStyle),
		),

		// CPU selected by default
		metricSelected: cpuMetric,
	}

	m.recalculateSizes()
	return m
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

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.recalculateSizes()

	case tea.KeyMsg:

		switch msg.String() {

		case "q", "ctrl+c":
			return m, tea.Quit

		case "up", "k":
			if m.metricSelected > cpuMetric {
				m.metricSelected--
			}

		case "down", "j":
			if m.metricSelected < networkMetric {
				m.metricSelected++
			}

		case "tab":
			m.tabSelected = (m.tabSelected + 1) % 3
			switch m.tabSelected {
			case 0:
				m.activeTab = 0
			case 1:
				m.activeTab = 1
			case 2:
				m.activeTab = 2
			}

		case "1":
			m.metricSelected = cpuMetric
		case "2":
			m.metricSelected = memoryMetric
		case "3":
			m.metricSelected = diskMetric
		case "4":
			m.metricSelected = networkMetric
		}

	// case tea.MouseClickMsg:

	// 	// Check if left sidebar clicked
	// 	if msg.X >= 0 && msg.X < m.navWidth {
	// 		if msg.Y >= 3 && msg.Y < 5 {
	// 			m.activeTab = 0
	// 		} else if msg.Y >= 5 && msg.Y < 7 {
	// 			m.activeTab = 1
	// 		} else if msg.Y >= 7 && msg.Y < 9 {
	// 			m.activeTab = 2
	// 		}
	// 	} else if msg.X >= m.navWidth && msg.X < m.navWidth+m.metricWidth {
	// 		// Metric box column clicked
	// 		yInBody := msg.Y - 3
	// 		cardH := m.bodyHeight / 4
	// 		if cardH > 0 && yInBody >= 0 {
	// 			cardIdx := yInBody / cardH
	// 			switch cardIdx {
	// 			case 0:
	// 				m.metricSelected = cpuMetric
	// 			case 1:
	// 				m.metricSelected = memoryMetric
	// 			case 2:
	// 				m.metricSelected = diskMetric
	// 			case 3:
	// 				m.metricSelected = networkMetric
	// 			}
	// 		}
	// 	}

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

		total, free, avail, cached, used, err := metrics.GetMemStats()

		if err != nil {
			fmt.Println("Memory:", err)
		} else {
			m.memTotal = total
			m.memFree = free
			m.memAvailable = avail
			m.memCached = cached
			m.memUsed = used

			if total > 0 {
				m.memoryPerc = float64(total-avail) / float64(total) * 100
			}

			// Small graph
			m.memory.Push(m.memoryPerc)
			m.memory.DrawBraille()

			// Large graph
			m.memoryDetail.Push(m.memoryPerc)
			m.memoryDetail.DrawBraille()
		}

		// ---------------- DISK READ & WRITE ----------------

		readIOPS, errR := metrics.GetDiskReadIOPS("")
		if errR != nil {
			fmt.Println("Disk read:", errR)
		} else {
			m.diskReadIOPS = readIOPS

			// Small graph
			m.diskRead.Push(readIOPS)
			m.diskRead.DrawBraille()

			// Large detail graph
			m.diskReadDetail.Push(readIOPS)
			m.diskReadDetail.DrawBraille()
		}

		writeIOPS, errW := metrics.GetDiskWriteIOPS("")
		if errW != nil {
			fmt.Println("Disk write:", errW)
		} else {
			m.diskWriteIOPS = writeIOPS

			// Small graph
			m.diskWrite.Push(writeIOPS)
			m.diskWrite.DrawBraille()

			// Large detail graph
			m.diskWriteDetail.Push(writeIOPS)
			m.diskWriteDetail.DrawBraille()
		}

		// ---------------- NETWORK ----------------

		rxPackets, txPackets, errN := metrics.GetPacketsStat("")

		if errN != nil {
			fmt.Println("Network:", errN)
		} else {
			m.netRXPackets = rxPackets
			m.netTXPackets = txPackets

			// Small graphs
			m.netRX.Push(rxPackets)
			m.netRX.DrawBraille()

			m.netTX.Push(txPackets)
			m.netTX.DrawBraille()

			// Large detail graphs
			m.netRXDetail.Push(rxPackets)
			m.netRXDetail.DrawBraille()

			m.netTXDetail.Push(txPackets)
			m.netTXDetail.DrawBraille()
		}

		return m, doTick()
	}

	return m, nil
}

func (m model) View() tea.View {

	// ---------------- TOP HEADER BADGE ----------------

	headerBadge := headerStyle.Render("Perfomace Page")

	// ---------------- 1st COLUMN: LEFT NAV SIDEBAR ----------------

	navPerf := "  Performance"
	if m.activeTab == 0 {
		navPerf = navActiveStyle.Render("> Performance")
	} else {
		navPerf = navInactiveStyle.Render("  Performance")
	}

	navProc := "  Processes"
	if m.activeTab == 1 {
		navProc = navActiveStyle.Render("> Processes")
	} else {
		navProc = navInactiveStyle.Render("  Processes")
	}

	navServ := "  Services"
	if m.activeTab == 2 {
		navServ = navActiveStyle.Render("> Services")
	} else {
		navServ = navInactiveStyle.Render("  Services")
	}

	navContent := fmt.Sprintf("\n%s\n\n%s\n\n%s", navPerf, navProc, navServ)

	navSidebar := gloss.NewStyle().
		Width(m.navWidth).
		Height(m.bodyHeight).
		BorderRight(true).
		BorderStyle(gloss.NormalBorder()).
		BorderForeground(gloss.Color("240")).
		Render(navContent)

	// ---------------- 2nd COLUMN: METRIC SELECTION BOX ----------------

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

	diskTitle := "Disk"
	diskGraph := gloss.JoinVertical(
		gloss.Left,
		m.diskRead.View(),
		m.diskWrite.View(),
	)

	diskBox := m.metricBox(
		diskTitle,
		diskGraph,
		diskMetric,
	)

	netTitle := "Net"
	networkGraph := gloss.JoinVertical(
		gloss.Left,
		m.netRX.View(),
		m.netTX.View(),
	)

	networkBox := m.metricBox(
		netTitle,
		networkGraph,
		networkMetric,
	)

	metricCards := gloss.JoinVertical(
		gloss.Left,
		cpuBox,
		memoryBox,
		diskBox,
		networkBox,
	)

	metricSidebar := gloss.NewStyle().
		Width(m.metricWidth).
		Height(m.bodyHeight).
		BorderStyle(gloss.NormalBorder()).
		BorderForeground(gloss.Color("240")).
		Render(metricCards)

	// ---------------- 3rd COLUMN: RIGHT DETAIL PANEL ----------------

	graphBoxHeight := m.bodyHeight / 2
	if graphBoxHeight < 5 {
		graphBoxHeight = 5
	}

	infoBoxHeight := m.bodyHeight - graphBoxHeight
	if infoBoxHeight < 4 {
		infoBoxHeight = 4
	}

	var detailGraph string
	var title string
	var textInfo string

	switch m.metricSelected {

	case cpuMetric:
		title = "CPU Performance"
		detailGraph = m.cpuDetail.View()
		colW := 22
		r1Labels := fmt.Sprintf("%s%s%s", infoLabelStyle.Width(colW).Render("Utilization"), infoLabelStyle.Width(colW).Render("Status"), infoLabelStyle.Width(colW).Render("Fetch Rate"))
		r1Values := fmt.Sprintf("%s%s%s", infoValueStyle.Width(colW).Render(fmt.Sprintf("%.2f%%", m.cpuPerc)), infoValueStyle.Width(colW).Render("Active"), infoValueStyle.Width(colW).Render(fetchFrequency.String()))
		r2Labels := fmt.Sprintf("%s%s%s", infoLabelStyle.Width(colW).Render("Architecture"), infoLabelStyle.Width(colW).Render("Metrics Source"), infoLabelStyle.Width(colW).Render("System Load"))
		r2Values := fmt.Sprintf("%s%s%s", infoValueStyle.Width(colW).Render("System CPU"), infoValueStyle.Width(colW).Render("/proc/stat"), infoValueStyle.Width(colW).Render("Normal"))
		textInfo = fmt.Sprintf("%s\n%s\n\n%s\n%s", r1Labels, r1Values, r2Labels, r2Values)

	case memoryMetric:
		title = "Memory Utilization"
		detailGraph = m.memoryDetail.View()
		colW := 22
		r1Labels := fmt.Sprintf("%s%s%s", infoLabelStyle.Width(colW).Render("Utilization"), infoLabelStyle.Width(colW).Render("Total Memory"), infoLabelStyle.Width(colW).Render("Used Memory"))
		r1Values := fmt.Sprintf("%s%s%s", infoValueStyle.Width(colW).Render(fmt.Sprintf("%.2f%%", m.memoryPerc)), infoValueStyle.Width(colW).Render(fmt.Sprintf("%.2f GB", m.memTotal)), infoValueStyle.Width(colW).Render(fmt.Sprintf("%.2f GB", m.memUsed)))
		r2Labels := fmt.Sprintf("%s%s%s", infoLabelStyle.Width(colW).Render("Available"), infoLabelStyle.Width(colW).Render("Free Memory"), infoLabelStyle.Width(colW).Render("Cached"))
		r2Values := fmt.Sprintf("%s%s%s", infoValueStyle.Width(colW).Render(fmt.Sprintf("%.2f GB", m.memAvailable)), infoValueStyle.Width(colW).Render(fmt.Sprintf("%.2f GB", m.memFree)), infoValueStyle.Width(colW).Render(fmt.Sprintf("%.2f GB", m.memCached)))
		textInfo = fmt.Sprintf("%s\n%s\n\n%s\n%s", r1Labels, r1Values, r2Labels, r2Values)

	case diskMetric:
		title = fmt.Sprintf(
			"Disk I/O Activity    %s    %s",
			diskReadStyle.Render("■ Read IOPS"),
			diskWriteStyle.Render("■ Write IOPS"),
		)
		detailGraph = gloss.JoinVertical(
			gloss.Left,
			m.diskReadDetail.View(),
			m.diskWriteDetail.View(),
		)
		totalIOPS := m.diskReadIOPS + m.diskWriteIOPS
		colW := 22
		r1Labels := fmt.Sprintf("%s%s%s", infoLabelStyle.Width(colW).Render("Read IOPS"), infoLabelStyle.Width(colW).Render("Write IOPS"), infoLabelStyle.Width(colW).Render("Total IOPS"))
		r1Values := fmt.Sprintf("%s%s%s", infoValueStyle.Width(colW).Render(fmt.Sprintf("%.0f", m.diskReadIOPS)), infoValueStyle.Width(colW).Render(fmt.Sprintf("%.0f", m.diskWriteIOPS)), infoValueStyle.Width(colW).Render(fmt.Sprintf("%.0f", totalIOPS)))
		r2Labels := fmt.Sprintf("%s%s%s", infoLabelStyle.Width(colW).Render("Storage Status"), infoLabelStyle.Width(colW).Render("Metrics Source"), infoLabelStyle.Width(colW).Render("Activity"))
		r2Values := fmt.Sprintf("%s%s%s", infoValueStyle.Width(colW).Render("Healthy"), infoValueStyle.Width(colW).Render("/proc/diskstats"), infoValueStyle.Width(colW).Render("Active"))
		textInfo = fmt.Sprintf("%s\n%s\n\n%s\n%s", r1Labels, r1Values, r2Labels, r2Values)

	case networkMetric:
		title = fmt.Sprintf(
			"Network Traffic    %s    %s",
			netRXStyle.Render("■ RX Received"),
			netTXStyle.Render("■ TX Transferred"),
		)
		detailGraph = gloss.JoinVertical(
			gloss.Left,
			m.netRXDetail.View(),
			m.netTXDetail.View(),
		)
		totalPkts := m.netRXPackets + m.netTXPackets
		colW := 22
		r1Labels := fmt.Sprintf("%s%s%s", infoLabelStyle.Width(colW).Render("RX Packets"), infoLabelStyle.Width(colW).Render("TX Packets"), infoLabelStyle.Width(colW).Render("Total Packets"))
		r1Values := fmt.Sprintf("%s%s%s", infoValueStyle.Width(colW).Render(fmt.Sprintf("%.0f pkts/s", m.netRXPackets)), infoValueStyle.Width(colW).Render(fmt.Sprintf("%.0f pkts/s", m.netTXPackets)), infoValueStyle.Width(colW).Render(fmt.Sprintf("%.0f pkts/s", totalPkts)))
		r2Labels := fmt.Sprintf("%s%s%s", infoLabelStyle.Width(colW).Render("Interface"), infoLabelStyle.Width(colW).Render("Link Status"), infoLabelStyle.Width(colW).Render("Flow Direction"))
		r2Values := fmt.Sprintf("%s%s%s", infoValueStyle.Width(colW).Render("All interfaces"), infoValueStyle.Width(colW).Render("Connected"), infoValueStyle.Width(colW).Render("Bi-directional"))
		textInfo = fmt.Sprintf("%s\n%s\n\n%s\n%s", r1Labels, r1Values, r2Labels, r2Values)
	}

	// Graph detail box covering top half (~50%)
	detailGraphBox := gloss.NewStyle().
		Width(m.detailWidth).
		Height(graphBoxHeight).
		BorderStyle(gloss.NormalBorder()).
		BorderForeground(gloss.Color("240")).
		Render(
			fmt.Sprintf(
				"%s\n%s",
				infoTitleStyle.Render(title),
				detailGraph,
			),
		)

	// Information text box below graph
	detailInfoBox := gloss.NewStyle().
		Width(m.detailWidth).
		Height(infoBoxHeight).
		BorderStyle(gloss.NormalBorder()).
		BorderForeground(gloss.Color("240")).
		Padding(1, 2).
		Render(textInfo)

	detailPanel := gloss.JoinVertical(
		gloss.Left,
		detailGraphBox,
		detailInfoBox,
	)

	// ---------------- FINAL FULL SCREEN LAYOUT ----------------

	mainBody := gloss.JoinHorizontal(
		gloss.Top,
		navSidebar,
		metricSidebar,
		detailPanel,
	)

	fullPage := gloss.JoinVertical(
		gloss.Left,
		headerBadge,
		"",
		mainBody,
	)

	v := tea.NewView(fullPage)
	v.AltScreen = true
	return v
}

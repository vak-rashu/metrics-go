package tui

// import (
// 	tea "charm.land/bubbletea/v2"
// 	metrics "github.com/vak-rashu/metrics-go/pkg"
// )

// // model data
// // type simplePage struct {
// // 	msg string
// // }

// type stat struct {
// 	msg string
// 	val metrics.CPUStat
// }

// // var blocks = []string{
// // 	"CPU",
// // 	"Processes",
// // 	"Memory",
// // }

// // func newSimplePage(msg string) simplePage {
// // 	return simplePage{msg: msg}
// // }

// // func (s simplePage) Init() tea.Cmd {
// // 	return func() tea.Msg {
// // 		gloss.Println(headLineStyle.Render(s.msg))
// // 		return nil
// // 	}
// // }

// func (c stat) Init() tea.Cmd {
// 	// return func() tea.Msg {
// 	// 	gloss.Println(headLineStyle.Render(c.msg))
// 	// 	return nil
// 	// }

// 	return getCPUStat()
// }

// // func (s simplePage) View() string {

// // 	// style display
// // 	listString := []string{}
// // 	for _, val := range blocks {
// // 		listString = append(listString, tabStyle.Render(val))
// // 	}
// // 	return gloss.JoinHorizontal(gloss.Bottom, listString...)

// // }

// func (c stat) View() tea.View {
// 	return tea.NewView(c.val.Cpu)
// }

// // func (s simplePage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
// // 	switch msg.(type) {
// // 	case tea.KeyMsg:
// // 		switch msg.(tea.KeyMsg).String() {
// // 		case "m":
// // 			return s, getCPUStat
// // 		case "enter":
// // 			return s, tea.Quit
// // 		}
// // 	}

// // 	return s, nil
// // }

// func (c stat) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
// 	switch msg := msg.(type) {
// 	case string:
// 		msg = string(msg)
// 		return c, tea.Quit
// 	}
// 	return c, nil
// }

// func getCPUStat() tea.Cmd {
// 	return func() tea.Msg {
// 		val := metrics.RetCPU()
// 		return val
// 	}
// }

// const url = "https://charm.sh"

// type model struct {
// 	status int
// 	err    error
// }

// func checkServer() tea.Msg {
// 	c := &http.Client{Timeout: 10 * time.Second}
// 	res, err := c.Get(url)
// 	if err != nil {
// 		return errMsg{err}
// 	}

// 	defer res.Body.Close()

// 	return statusMsg(res.StatusCode)
// }

// type statusMsg int
// type errMsg struct{ err error }

// func (e errMsg) Error() string { return e.err.Error() }

// func (m model) Init() tea.Cmd {
// 	return checkServer
// }

// func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
// 	switch msg := msg.(type) {
// 	case statusMsg:
// 		m.status = int(msg)
// 		return m, tea.Quit
// 	case errMsg:
// 		m.err = msg
// 		return m, tea.Quit

// 	case tea.KeyPressMsg:
// 		if msg.Mod == tea.ModCtrl && msg.Code == 'c' {
// 			return m, tea.Quit
// 		}
// 	}

// 	return m, nil

// }

// func (m model) View() tea.View {
// 	if m.err != nil {
// 		return tea.NewView(fmt.Sprintf("\nWe had some trouble: %v\n\n", m.err))
// 	}

// 	s := fmt.Sprintf("Checking %s ... ", url)
// 	if m.status > 0 {
// 		s += fmt.Sprintf("%d %s!", m.status, http.StatusText(m.status))
// 	}
// 	return tea.NewView("\n" + s + "\n\n")
// }

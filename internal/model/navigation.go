package model

import tea "github.com/charmbracelet/bubbletea"

// Page codes
// It denotes what page you're in.
var PAGE_CODE = map[string]int{
    "INITIAL_PAGE": 0,
    "REPORT_PAGE":  1,
    "ERROR_PAGE":  -1,
	"OAUTH_PAGE":  2,
}

// Update loop for the views where you're choosing a task.
func navigate(m Model, msg tea.Msg) int {
	tPos := len(m.GetChoices()) - 1
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j", "down":
			m.IncrementPos()
			if m.GetPos() > tPos {
				m.SetPos(tPos)
			}
		case "k", "up":
			m.DecreasePos()
			if m.GetPos() < 0 {
				m.SetPos(0)
			}
		case "enter":
			m.SetChosen(true)
			return m.GetPos()
		}
	}

	return -1
}

// Backtracking function to navigate to previous page
func backTracking(code int) (tea.Model, tea.Cmd) {
	switch code {
	case PAGE_CODE["INITIAL_PAGE"]:
		return RouterPage().Update(nil)
	case PAGE_CODE["REPORT_PAGE"]:
		return ReportPage(code).Update(nil)
	}
	return nil, nil
}
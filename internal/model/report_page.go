package model

import (
	tea "github.com/charmbracelet/bubbletea"
)



type reportPageModel struct {
	page Page
}

func ReportPage(c int) reportPageModel {
	m := reportPageModel{
		page: Page{
			choicePos: 0, 
			chosen: false, 
			quitting: false, 
			pageCode: PAGE_CODE["REPORT_PAGE"], 
			choices: []string{"Gerar relatorio - dia seguinte", "Voltar"},
			LastPageCode: c,
			},
	}
	return m
}
func (m reportPageModel) Init() tea.Cmd {
    return nil
}


func (m reportPageModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	
	if msg, ok := msg.(tea.KeyMsg); ok {
		k := msg.String()
		if k == "q" || k == "ctrl+c" {
			m.page.quitting = true
			return m, tea.Quit
		}
		if k == "esc" {
			return backTracking(m.page.LastPageCode)
		}
	}

	return m.updateChoices(msg)
}


func (m reportPageModel) View() string {
	var s string
	if m.page.quitting {
		return "\n  See you later!\n\n"
	}
	s = renderView(&m)
	return mainStyle.Render("\n" + s + "\n\n")
}

// Sub-update functions

// Update loop for the first view where you're choosing a task.
func (m reportPageModel) updateChoices(msg tea.Msg) (tea.Model, tea.Cmd) {
	i := navigate(&m, msg)
	switch i {
		case 1: // exit
			return backTracking(m.page.LastPageCode)
	}

	return m, nil
}

func (m reportPageModel) GetChoices() []string {
	return m.page.choices
}

func (m *reportPageModel) SetChosen(chosen bool) {
	m.page.chosen = chosen
}

func (m *reportPageModel) GetPos() int {
	return m.page.choicePos
}

func (m *reportPageModel) IncrementPos() {
	m.page.choicePos++
}

func (m *reportPageModel) DecreasePos() {
	m.page.choicePos--
}

func (m *reportPageModel) SetPos(pos int) {
	m.page.choicePos = pos
}
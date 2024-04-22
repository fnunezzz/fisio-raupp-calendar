package model

import (
	tea "github.com/charmbracelet/bubbletea"
)



type reportPageModel struct {
	Page Page
}

func ReportPage(c int) reportPageModel {
	m := reportPageModel{
		Page: Page{
			ChoicePos: 0, 
			Chosen: false, 
			Quitting: false, 
			PageCode: PAGE_CODE["REPORT_PAGE"], 
			Choices: []string{"Gerar relatorio - dia seguinte", "Voltar"},
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
			m.Page.Quitting = true
			return m, tea.Quit
		}
		if k == "esc" {
			return backTracking(m.Page.LastPageCode)
		}
	}

	return m.updateChoices(msg)
}


func (m reportPageModel) View() string {
	var s string
	if m.Page.Quitting {
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
			return backTracking(m.Page.LastPageCode)
	}

	return m, nil
}

func (m reportPageModel) GetChoices() []string {
	return m.Page.Choices
}

func (m *reportPageModel) SetChosen(chosen bool) {
	m.Page.Chosen = chosen
}

func (m *reportPageModel) GetPos() int {
	return m.Page.ChoicePos
}

func (m *reportPageModel) IncrementPos() {
	m.Page.ChoicePos++
}

func (m *reportPageModel) DecreasePos() {
	m.Page.ChoicePos--
}

func (m *reportPageModel) SetPos(pos int) {
	m.Page.ChoicePos = pos
}
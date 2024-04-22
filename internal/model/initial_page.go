package model

import (
	tea "github.com/charmbracelet/bubbletea"
)




type initialPageModel struct {
	page Page
}

func InitialPage() initialPageModel {
	m := initialPageModel{
		page: Page{
			choicePos: 0, 
			chosen: false, 
			quitting: false, 
			pageCode: PAGE_CODE["INITIAL_PAGE"], 
			choices: []string{"Relatorio", "Ajuda", "Sobre", "Sair"},
			},
	}
	return m
}



func (m initialPageModel) Init() tea.Cmd {
	return nil
}


func (m initialPageModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	
	if msg, ok := msg.(tea.KeyMsg); ok {
		k := msg.String()
		if k == "q" || k == "esc" || k == "ctrl+c" {
			m.page.quitting = true
			return m, tea.Quit
		}
	}


	return m.updateChoices(msg)
}

func (m initialPageModel) GetChoices() []string {
	return m.page.choices
}


func (m initialPageModel) View() string {
	var s string
	if m.page.quitting {
		return "\n  See you later!\n\n"
	}

	s = renderView(&m)
	return mainStyle.Render("\n" + s + "\n\n")
}


func (m initialPageModel) updateChoices(msg tea.Msg) (tea.Model, tea.Cmd) {
	i := navigate(&m, msg)
	switch i {
		case 0:
			return ReportPage(m.page.pageCode).Update(msg)
	}

	return m, nil
}

func (m *initialPageModel) SetChosen(chosen bool) {
	m.page.chosen = chosen
}

func (m *initialPageModel) GetPos() int {
	return m.page.choicePos
}

func (m *initialPageModel) IncrementPos() {
	m.page.choicePos += 1
}

func (m *initialPageModel) DecreasePos() {
	m.page.choicePos -= 1
}

func (m *initialPageModel) SetPos(pos int) {
	m.page.choicePos = pos
}
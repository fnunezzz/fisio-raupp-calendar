package model

import (
	tea "github.com/charmbracelet/bubbletea"
)




type routerPageModel struct {
	page Page
}

func RouterPage() routerPageModel {
	m := routerPageModel{
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



func (m routerPageModel) Init() tea.Cmd {
	return nil
}


func (m routerPageModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	
	if msg, ok := msg.(tea.KeyMsg); ok {
		k := msg.String()
		if k == "q" || k == "esc" || k == "ctrl+c" {
			m.page.quitting = true
			return m, tea.Quit
		}
	}


	return m.updateChoices(msg)
}

func (m routerPageModel) GetChoices() []string {
	return m.page.choices
}


func (m routerPageModel) View() string {
	var s string
	if m.page.quitting {
		return "\n  See you later!\n\n"
	}

	s = renderView(&m)
	return mainStyle.Render("\n" + s + "\n\n")
}


func (m routerPageModel) updateChoices(msg tea.Msg) (tea.Model, tea.Cmd) {
	i := navigate(&m, msg)
	switch i {
		case 0:
			return ReportPage(m.page.pageCode).Update(msg)
	}

	return m, nil
}

func (m *routerPageModel) SetChosen(chosen bool) {
	m.page.chosen = chosen
}

func (m *routerPageModel) GetPos() int {
	return m.page.choicePos
}

func (m *routerPageModel) IncrementPos() {
	m.page.choicePos += 1
}

func (m *routerPageModel) DecreasePos() {
	m.page.choicePos -= 1
}

func (m *routerPageModel) SetPos(pos int) {
	m.page.choicePos = pos
}
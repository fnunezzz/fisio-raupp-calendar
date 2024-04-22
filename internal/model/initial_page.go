package model

import (
	tea "github.com/charmbracelet/bubbletea"
)




type initialPageModel struct {
	Page Page
}

func InitialPage() initialPageModel {
	m := initialPageModel{
		Page: Page{
			ChoicePos: 0, 
			Chosen: false, 
			Quitting: false, 
			PageCode: PAGE_CODE["INITIAL_PAGE"], 
			Choices: []string{"Relatorio", "Ajuda", "Sobre", "Sair"},
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
			m.Page.Quitting = true
			return m, tea.Quit
		}
	}


	return m.updateChoices(msg)
}

func (m initialPageModel) GetChoices() []string {
	return m.Page.Choices
}


func (m initialPageModel) View() string {
	var s string
	if m.Page.Quitting {
		return "\n  See you later!\n\n"
	}

	s = renderView(&m)
	return mainStyle.Render("\n" + s + "\n\n")
}


func (m initialPageModel) updateChoices(msg tea.Msg) (tea.Model, tea.Cmd) {
	i := navigate(&m, msg)
	switch i {
		case 0:
			return ReportPage(m.Page.PageCode).Update(msg)
	}

	return m, nil
}

func (m *initialPageModel) SetChosen(chosen bool) {
	m.Page.Chosen = chosen
}

func (m *initialPageModel) GetPos() int {
	return m.Page.ChoicePos
}

func (m *initialPageModel) IncrementPos() {
	m.Page.ChoicePos += 1
}

func (m *initialPageModel) DecreasePos() {
	m.Page.ChoicePos -= 1
}

func (m *initialPageModel) SetPos(pos int) {
	m.Page.ChoicePos = pos
}
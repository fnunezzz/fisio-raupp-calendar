package model

import (
	tea "github.com/charmbracelet/bubbletea"
)



type tokenPageModel struct {
	page Page
	message string
}

func TokenPage(message string) tokenPageModel {
	m := tokenPageModel{
		page: Page{
			quitting: false, 
			pageCode: PAGE_CODE["ERROR_PAGE"],
			},
		message: message,
	}
	return m
}
func (m tokenPageModel) Init() tea.Cmd {
    return nil
}


func (m tokenPageModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	
	if msg, ok := msg.(tea.KeyMsg); ok {
		k := msg.String()
		if k == "q" || k == "esc" || k == "ctrl+c" {
			m.page.quitting = true
			return m, tea.Quit
		}
	}

	return m, nil
}


func (m tokenPageModel) View() string {
	if m.page.quitting {
		return "\n  See you later!\n\n"
	}

	tpl := "Error\n\n"
	tpl += m.message + "\n\n" +
		subtleStyle.Render("esc: sair") + dotStyle +
		subtleStyle.Render("q: sair")


	return mainStyle.Render("\n" + tpl + "\n\n")
}
package model

import (
	tea "github.com/charmbracelet/bubbletea"
)



type errorPageModel struct {
	page Page
	message string
}

func ErrorPage(message string) errorPageModel {
	m := errorPageModel{
		page: Page{
			quitting: false, 
			pageCode: PAGE_CODE["ERROR_PAGE"],
			},
		message: message,
	}
	return m
}
func (m errorPageModel) Init() tea.Cmd {
    return nil
}


func (m errorPageModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	
	if msg, ok := msg.(tea.KeyMsg); ok {
		k := msg.String()
		if k == "q" || k == "esc" || k == "ctrl+c" {
			m.page.quitting = true
			return m, tea.Quit
		}
	}

	return m, nil
}


func (m errorPageModel) View() string {
	if m.page.quitting {
		return "\n  See you later!\n\n"
	}

	tpl := "Error\n\n"
	tpl += m.message + "\n\n"
	tpl += subtleStyle.Render("j/k, ↑/↓: mover cursor") + dotStyle +
		subtleStyle.Render("enter: selecionar") + dotStyle +
		subtleStyle.Render("esc: voltar") + dotStyle +
		subtleStyle.Render("q: sair")


	return mainStyle.Render("\n" + tpl + "\n\n")
}
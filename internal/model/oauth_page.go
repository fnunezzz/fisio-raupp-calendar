package model

import (
	tea "github.com/charmbracelet/bubbletea"
)



type oauthPageModel struct {
	page Page
	message string
}

func OauthPage(message string) oauthPageModel {
	m := oauthPageModel{
		page: Page{
			quitting: false, 
			pageCode: PAGE_CODE["OAUTH_PAGE"],
			},
		message: message,
	}
	return m
}
func (m oauthPageModel) Init() tea.Cmd {
    return nil
}


func (m oauthPageModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	
	if msg, ok := msg.(tea.KeyMsg); ok {
		k := msg.String()
		if k == "q" || k == "esc" || k == "ctrl+c" {
			m.page.quitting = true
			return m, tea.Quit
		}
	}

	return m, nil
}


// TODO - Implement the view function for the oauthPageModel
// TODO - Should execute the entire auth process
func (m oauthPageModel) View() string {
	if m.page.quitting {
		return ciao
	}

	tpl := "Error\n\n"
	tpl += m.message + "\n\n"
	tpl += subtleStyle.Render("j/k, ↑/↓: mover cursor") + dotStyle +
		subtleStyle.Render("enter: selecionar") + dotStyle +
		subtleStyle.Render("esc: voltar") + dotStyle +
		subtleStyle.Render("q: sair")


	return mainStyle.Render("\n" + tpl + "\n\n")
}
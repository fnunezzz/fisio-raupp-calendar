package model

import (
	tea "github.com/charmbracelet/bubbletea"
)



type reportPageModel struct {
	page Page
}

func ReportPage() reportPageModel {
	m := reportPageModel{
		page: Page{
			pageCode: PAGE_CODE["REPORT_PAGE"],
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
		if k == "q" || k == "esc" || k == "ctrl+c" || k == "enter"{
			return MainPage(), nil
		}
	}

	return m, nil
}


func (m reportPageModel) View() string {
	tpl := "Relatório Gerado!\n\n"
	tpl += subtleStyle.Render("enter: voltar") + dotStyle +
		subtleStyle.Render("esc: voltar") + dotStyle +
		subtleStyle.Render("q: voltar")


	return mainStyle.Render("\n" + tpl + "\n\n")
}
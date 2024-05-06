package model

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)




type mainPageModel struct {
	page Page
}

func MainPage() mainPageModel {
	m := mainPageModel{
		page: Page{
			choicePos: 0, 
			chosen: false, 
			quitting: false, 
			pageCode: PAGE_CODE["INITIAL_PAGE"], 
			choices: []string{"Gerar Relatório", "Sair"},
			},
	}
	return m
}



func (m mainPageModel) Init() tea.Cmd {
	return nil
}


func (m mainPageModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	
	if msg, ok := msg.(tea.KeyMsg); ok {
		k := msg.String()
		if k == "q" || k == "esc" || k == "ctrl+c" {
			m.page.quitting = true
			return m, tea.Quit
		}
	}


	return m.updateChoices(msg)
}

func (m mainPageModel) GetChoices() []string {
	return m.page.choices
}


func (m mainPageModel) View() string {
	if m.page.quitting {
		return ciao
	}

    c := m.GetPos()

	tpl := "Relatório\n\n"
	tpl += "%s\n\n"
	tpl += subtleStyle.Render("j/k, ↑/↓: mover cursor") + dotStyle +
		subtleStyle.Render("enter: selecionar") + dotStyle +
		subtleStyle.Render("esc: voltar") + dotStyle +
		subtleStyle.Render("q: sair")

	choices := ""
	for i, choice := range m.GetChoices() {
		choices += fmt.Sprintf("%s\n", checkbox(choice, c == i))
	}

	s := fmt.Sprintf(tpl, choices)
	return mainStyle.Render("\n" + s + "\n\n")
}

func (m mainPageModel) createReport() error {
	// todo calendar
	// patients := map[string]string{
	// 	"Patient 1": "10:00",
	// 	"Patient 2": "10:00",
	// 	"Patient 3": "11:20",
	// 	"Patient 4": "13:40",
	// 	"Patient 5": "18:00",
	// }
	// xlsxService := service.NewXlsxService()
	// err := xlsxService.GenerateXlsxReport(patients, time.Now())
	// if err != nil {
	// 	return err
	// }


	return nil
}


func (m mainPageModel) updateChoices(msg tea.Msg) (tea.Model, tea.Cmd) {
	i := navigate(&m, msg)
	switch i {
		case 0:
			err := m.createReport()
			if err != nil {
				return ErrorPage(err.Error()), nil
			}
			return ReportPage(), nil
		case 1:
			m.page.quitting = true
			return m, tea.Quit
	}

	return m, nil
}

func (m *mainPageModel) SetChosen(chosen bool) {
	m.page.chosen = chosen
}

func (m *mainPageModel) GetPos() int {
	return m.page.choicePos
}

func (m *mainPageModel) IncrementPos() {
	m.page.choicePos += 1
}

func (m *mainPageModel) DecreasePos() {
	m.page.choicePos -= 1
}

func (m *mainPageModel) SetPos(pos int) {
	m.page.choicePos = pos
}
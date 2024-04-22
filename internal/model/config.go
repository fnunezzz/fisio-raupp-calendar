package model

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	dotChar           = " • "
)

// General stuff for styling the view
var (
	subtleStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
	checkboxStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("212"))
	dotStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("236")).Render(dotChar)
	mainStyle     = lipgloss.NewStyle().MarginLeft(2)
)

type Page struct {
	ChoicePos		int
	Chosen   		bool
	Quitting 		bool
	PageCode 		int
	Choices 		[]string
	LastPageCode 	int
}

type Model interface {
	GetPos() int
	SetPos(int)
	GetChoices() []string
	IncrementPos()
	DecreasePos()
	SetChosen(bool)
	updateChoices(msg tea.Msg) (tea.Model, tea.Cmd)
}

func renderView(m Model) string {
    c := m.GetPos()

	tpl := "Relatorio\n\n"
	tpl += "%s\n\n"
	tpl += subtleStyle.Render("j/k, ↑/↓: mover cursor") + dotStyle +
		subtleStyle.Render("enter: selecionar") + dotStyle +
		subtleStyle.Render("esc: voltar") + dotStyle +
		subtleStyle.Render("q: sair")

	choices := ""
	for i, choice := range m.GetChoices() {
		choices += fmt.Sprintf("%s\n", checkbox(choice, c == i))
	}

	return fmt.Sprintf(tpl, choices)
}

// render selected/unselected checkboxes
func checkbox(label string, checked bool) string {
	if checked {
		return checkboxStyle.Render("[x] " + label)
	}
	return fmt.Sprintf("[ ] %s", label)
}
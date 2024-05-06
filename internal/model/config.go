package model

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/oauth2"
)

const (
	dotChar           = " • "
	ciao = "\n   Ciao\n\n"
)

// General stuff for styling the view
var (
	subtleStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
	checkboxStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("212"))
	dotStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("236")).Render(dotChar)
	mainStyle     = lipgloss.NewStyle().MarginLeft(2)
	CREDENTIALS []byte = nil
	TOKEN *oauth2.Token = nil
)

type Page struct {
	choicePos		int
	chosen   		bool
	quitting 		bool
	pageCode 		int
	choices 		[]string
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

// render selected/unselected checkboxes
func checkbox(label string, checked bool) string {
	if checked {
		return checkboxStyle.Render("[x] " + label)
	}
	return fmt.Sprintf("[ ] %s", label)
}
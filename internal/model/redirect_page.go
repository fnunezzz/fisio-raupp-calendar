package model

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/fnunezzz/fisio-raupp-calendar/internal/service"
)



type redirectPageModel struct {
	page Page
	message string
	tries int
	viewport    viewport.Model
	textarea    textarea.Model
}

type tokenExists bool

func RedirectPage(message string) redirectPageModel {
	ta := textarea.New()
	ta.Prompt = ""
	ta.ShowLineNumbers = false
	ta.CharLimit = 10000
	ta.SetWidth(50)
	ta.SetHeight(10)
	vp := viewport.New(50, 10)
	vp.SetContent(message)
	m := redirectPageModel{
		page: Page{
			quitting: false, 
			pageCode: PAGE_CODE["REDIRECT_PAGE"],
			},
		message: message,
		tries: 0,
		textarea: ta,
		viewport: vp,
	}
	return m
}
func (m redirectPageModel) Init() tea.Cmd {
    return m.checkToken
}


func (m redirectPageModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

		case tea.KeyMsg:
			if (msg.String() == "esc" || msg.String() == "q") {
				m.page.quitting = true
				return m, tea.Quit
			}
			return m, m.checkToken
			
		case errMsg:
			return ErrorPage(msg.Error()).Update(nil)

		case tokenExists:
			return MainPage().Update(nil)

		default:
			if m.tries > 120 {
				return ErrorPage("Timeout").Update(nil)
			}
			return m, m.checkToken
	}
}


func (m redirectPageModel) View() string {
	if m.page.quitting {
		return ciao
	}

	tpl := fmt.Sprintf(
		"Copie o link abaixo e siga as instruções\n\n%s\n%s",
		m.viewport.View(),
		m.textarea.View(),
	) + "\n\n"


	return mainStyle.Render("\n" + tpl + "\n\n")
}


func (m *redirectPageModel) checkToken() tea.Msg {
	m.tries++
	time.Sleep(1 * time.Second)

		
	tokenService := service.NewTokenService()

	err := tokenService.CheckToken()
	if err == nil {
		return tokenExists(true)
	}
	return m
}
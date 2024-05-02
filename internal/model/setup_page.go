package model

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/fnunezzz/fisio-raupp-calendar/internal/service"
)

const (
	padding  = 2
	maxWidth = 80
)

func SetupPage() setupPageModel {
	prog := progress.New(progress.WithScaledGradient("#FF7CCB", "#FDFF8C"))

	s := setupPageModel{
		progress: prog,
		step: 0,
	}

	s.steps = []func() tea.Msg{s.createFolder, s.checkCredentials, s.checkToken}
	return s

}

type step int

type errMsg struct{ err error }

func (e errMsg) Error() string { return e.err.Error() }

type setupPageModel struct {
	percent  float64
	progress progress.Model
	step int
	steps []func() tea.Msg
}

func (m setupPageModel) Init() tea.Cmd {
	return m.steps[m.step]
}

func (m setupPageModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.progress.Width = msg.Width - padding*2 - 4
		if m.progress.Width > maxWidth {
			m.progress.Width = maxWidth
		}
		return m, nil

	case step:
		m.percent += (float64(len(m.steps)) + 1) / 10
		
		if m.percent > 1.0 {
			m.percent = 1.0
			return RouterPage().Update(msg)
		}

		return m, m.steps[msg]
		

	case errMsg:
		return ErrorPage(msg.Error()).Update(nil)

	default:
		return m, nil
	}
}

func (m setupPageModel) View() string {
	pad := strings.Repeat(" ", padding)
	return "\n" +
		pad + m.progress.ViewAs(m.percent) + "\n\n"
}

func (m *setupPageModel) checkToken() tea.Msg {
	time.Sleep(1 * time.Second)
	tokenService := service.NewTokenService()
	err := tokenService.CheckToken()
	if err != nil {

		webClient := service.NewClientService()
		go webClient.StartClient()
		defer webClient.StopClient()
		
		tokenService := service.NewTokenService()
		_, err := tokenService.GenerateToken()
		if err != nil {
			f := fmt.Sprintf("Unable to generate token: %v", err)
			err = errors.New(f)
			return errMsg{err}
		}
	}
	m.step = m.step + 1
	return step(m.step)
}

func (m *setupPageModel) createFolder() tea.Msg {
	time.Sleep(1 * time.Second)
	folderSerivce := service.NewFolderService()
	err := folderSerivce.CreateFolder()
	if err != nil {
		return errMsg{err}
	}
	m.step = m.step + 1
	return step(m.step)
}

func (m *setupPageModel) checkCredentials() tea.Msg {
	time.Sleep(1 * time.Second)
	credentialsService := service.NewCredentialsService()
	
	_, err := credentialsService.LoadCredentials()
	if err != nil {
			return errMsg{err}
	}
	m.step = m.step + 1
	return step(m.step)
}
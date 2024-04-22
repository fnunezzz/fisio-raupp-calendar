package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/fnunezzz/fisio-raupp-calendar/internal/model"
)



func main() {
        p := tea.NewProgram(model.InitialPage())
        if _, err := p.Run(); err != nil {
                fmt.Fprintf(os.Stderr, "Error: %v", err)
                os.Exit(1)
        }

}
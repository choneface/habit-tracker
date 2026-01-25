package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	Content string
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
		case tea.KeyMsg:
			if msg.String() == "q" {
				return m, tea.Quit
			}
	}

	var cmd tea.Cmd
	return m, cmd
}

func (m model) View() string {
	return m.Content
}

func main() {
	m := model{ Content: "Hello World" } 
	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

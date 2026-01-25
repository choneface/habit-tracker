package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

type habit struct {
	Title string
	Description string 
	History []byte 
}

type model struct {
	Title string 
	Habits []habit
}

func NewModel(title string) model {
	return model {
		Title: title,
		Habits: []habit{},
	}
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

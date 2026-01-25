package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

type habit struct {
	Title string
	Description string 
	History []byte 
}

type habit_view struct {
	Title string 
	Habits []habit
	Index uint8
}

type model struct {
	Title string
	HabitView habit_view
	HelpOptions string 
}

func NewModel(title string) model {
	hv := habit_view {
		Title: title,
		Habits: []habit{
			{ "Chinese Lessons", "Need passable chinese before having a kid", []byte{1,0,1,1,0,1,0,1,0,1,1,1,0,1,0},},
		},
		Index: 0,
	}

	return model {
		Title: "Habit Tracker",
		HabitView: hv,
		HelpOptions: "q to exit",
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

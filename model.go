package main

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	habit_view_mode = 0
)


type habit struct {
	Title string
	Description string 
	History []byte 
}

type habit_view struct {
	Habits []habit
	Index int
}

type model struct {
	Title string
	HabitView habit_view

	mode uint8 
	keys keyMap
}



func NewModel(title string) model {
	hv := habit_view {
		Habits: []habit{
			{ 
				"Chinese Lessons", 
				"Need passable chinese before having a kid",
				[]byte{1,0,1,1,0,1,0,1,0,1,1,1,0,1,0},
			}, {
				"Spanish Lessons", 
				"Need passable spanish before having a kid",
				[]byte{0,1,0,1,0,1,1,1,0,1,1,0,0,0,1},
			},
		},
		Index: 0,
	}

	return model {
		Title: "Habit Tracker",
		HabitView: hv,
		mode: habit_view_mode,
		keys: keys,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
		case tea.KeyMsg:
			switch {
			case key.Matches(msg, m.keys.Quit):
				return m, tea.Quit
			case key.Matches(msg, m.keys.Right):
				idx := m.HabitView.Index + 1 
				if idx >= len(m.HabitView.Habits) {
					idx = 0 
				}
				m.HabitView.Index = idx 
				break
			case key.Matches(msg, m.keys.Left): 
				idx := m.HabitView.Index - 1
				if idx < 0 {
					idx = len(m.HabitView.Habits) - 1 
				}
				m.HabitView.Index = idx 
				break
		}
	}

	var cmd tea.Cmd
	return m, cmd
}

package main

import (
	"github.com/charmbracelet/bubbles/help"
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

type habitView struct {
	Habits []habit
	Index int
}

type model struct {
	Title string
	HabitView habitView

	mode uint8 
	habitViewKeys habitViewKeyMap
	help help.Model
}



func NewModel(title string) model {
	hv := habitView {
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
		habitViewKeys: keys,
		help: help.New(),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.mode == habit_view_mode {
		return m.habitViewModeUpdate(msg)
	} 
	panic("Unsupported mode")
}
func (m model) habitViewModeUpdate(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
		case tea.KeyMsg:
			switch {
			case key.Matches(msg, m.habitViewKeys.Quit):
				return m, tea.Quit
			case key.Matches(msg, m.habitViewKeys.Right):
				idx := m.HabitView.Index + 1 
				if idx >= len(m.HabitView.Habits) {
					idx = 0 
				}
				m.HabitView.Index = idx 
				break
			case key.Matches(msg, m.habitViewKeys.Left): 
				idx := m.HabitView.Index - 1
				if idx < 0 {
					idx = len(m.HabitView.Habits) - 1 
				}
				m.HabitView.Index = idx 
				break
			case key.Matches(msg, m.habitViewKeys.Help):
				m.help.ShowAll = !m.help.ShowAll
				break
		}
	}

	var cmd tea.Cmd
	return m, cmd
}

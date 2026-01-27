package main

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
	tea "github.com/charmbracelet/bubbletea"
)

var blurredStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))

const (
	habit_view_mode = 0
	input_mode = 1
)


type habit struct {
	Title string
	Description string 
	History []byte 
}

type input struct {
	Title textinput.Model
	Description textinput.Model
	FocusIndex int
}

type habitView struct {
	Habits []habit
	Index int
}

type model struct {
	Title string
	HabitView habitView
	Input input 

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

	var t textinput.Model
	t = textinput.New()
	t.Placeholder = "Title" 
	t.CharLimit = 32
	t.PromptStyle = headerStyle
	t.TextStyle = headerStyle
	t.Focus()
	t.Width = 32

	var d textinput.Model 
	d = textinput.New()
	d.Placeholder = "Description" 
	d.CharLimit = 32
	d.PromptStyle = blurredStyle
	d.TextStyle = blurredStyle
	d.Width = 32

	i := input{
		Title: t,
		Description: d,
		FocusIndex: 0, 
	}

	return model {
		Title: "Habit Tracker",
		HabitView: hv,
		Input: i,
		mode: habit_view_mode,
		habitViewKeys: keys,
		help: help.New(),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch (m.mode) {
	case habit_view_mode:
		return m.habitViewModeUpdate(msg)
	case input_mode:
		return m.inputModeUpdate(msg)
	} 
	panic("Unsupported mode")
}

func (m model) inputModeUpdate(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, inputKeys.Up):
			m.Input.FocusIndex += 1
			if m.Input.FocusIndex > 1 {
				m.Input.FocusIndex = 0
			}
			var cmd tea.Cmd
			if m.Input.FocusIndex == 0 {
				cmd = m.Input.Title.Focus()
				m.Input.Description.Blur()

				m.Input.Title.TextStyle = headerStyle 
				m.Input.Title.PromptStyle = headerStyle

				m.Input.Description.TextStyle = blurredStyle
				m.Input.Description.PromptStyle = blurredStyle
			} else {
				cmd = m.Input.Description.Focus()
				m.Input.Title.Blur()

				m.Input.Description.TextStyle = headerStyle 
				m.Input.Description.PromptStyle = headerStyle

				m.Input.Title.TextStyle = blurredStyle
				m.Input.Title.PromptStyle = blurredStyle
			}
			return m, cmd
		case key.Matches(msg, inputKeys.Down):
			m.Input.FocusIndex -= 1
			if m.Input.FocusIndex < 0 {
				m.Input.FocusIndex = 1
			}

			var cmd tea.Cmd
			if m.Input.FocusIndex == 0 {
				cmd = m.Input.Title.Focus()
				m.Input.Description.Blur()

				m.Input.Title.TextStyle = headerStyle 
				m.Input.Title.PromptStyle = headerStyle

				m.Input.Description.TextStyle = blurredStyle
				m.Input.Description.PromptStyle = blurredStyle
			} else {
				cmd = m.Input.Description.Focus()
				m.Input.Title.Blur()

				m.Input.Description.TextStyle = headerStyle 
				m.Input.Description.PromptStyle = headerStyle

				m.Input.Title.TextStyle = blurredStyle
				m.Input.Title.PromptStyle = blurredStyle
			}
			return m, cmd
		case key.Matches(msg, inputKeys.Quit):
			m.mode = habit_view_mode
			break
		case key.Matches(msg, inputKeys.Submit):
			// TODO handle this 
			break
		case key.Matches(msg, inputKeys.Help):
			m.help.ShowAll = !m.help.ShowAll
			break
		}
	}
	var cmd1, cmd2 tea.Cmd
	m.Input.Title, cmd1 = m.Input.Title.Update(msg)
	m.Input.Description, cmd2 = m.Input.Description.Update(msg)
	
	return m, tea.Batch(cmd1, cmd2) 
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
			case key.Matches(msg, m.habitViewKeys.AddHabit):
				m.mode = input_mode
				break
		}
	}

	var cmd tea.Cmd
	return m, cmd
}

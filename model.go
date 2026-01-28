package main

import (
	"log"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var blurredStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))

const (
	HabitViewMode = 0
	InputMode = 1
)


type Habit struct {
	Title string
	Description string 
	History []byte 
}

type Input struct {
	Title textinput.Model
	Description textinput.Model
	focusIndex int
}

type habitView struct {
	Habits []Habit
	Index int
}

type Model struct {
	Title string
	HabitView habitView
	Input Input 

	Mode uint8 
	habitViewKeys habitViewKeyMap
	Help help.Model

	storage Storage
}



func NewModel(s Storage) Model {
	hv := habitView {
		Habits: s.GetHabits(), 
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

	i := Input{
		Title: t,
		Description: d,
		focusIndex: 0, 
	}

	var mode uint8 
	mode = HabitViewMode
	if len(hv.Habits) == 0 {
		mode = InputMode
	}

	return Model {
		Title: "Habit Tracker",
		HabitView: hv,
		Input: i,
		Mode: mode,
		habitViewKeys: habitViewKeys,
		Help: help.New(),
		storage: s,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch (m.Mode) {
	case HabitViewMode:
		return m.habitViewModeUpdate(msg)
	case InputMode:
		return m.inputModeUpdate(msg)
	} 
	panic("Unsupported mode")
}

func (m Model) inputModeUpdate(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, inputKeys.Up):
			m.Input.focusIndex += 1
			if m.Input.focusIndex > 1 {
				m.Input.focusIndex = 0
			}
			var cmd tea.Cmd
			if m.Input.focusIndex == 0 {
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
			m.Input.focusIndex -= 1
			if m.Input.focusIndex < 0 {
				m.Input.focusIndex = 1
			}

			var cmd tea.Cmd
			if m.Input.focusIndex == 0 {
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
			m.Mode = HabitViewMode
			break
		case key.Matches(msg, inputKeys.Submit):
			t := m.Input.Title.Value()
			d := m.Input.Description.Value()
			m.storage.SaveNewHabit(t, d)
			m.HabitView.Habits = m.storage.GetHabits()
			log.Printf("Habit length: %d", len(m.HabitView.Habits))
			break
		case key.Matches(msg, inputKeys.Help):
			m.Help.ShowAll = !m.Help.ShowAll
			break
		}
	}
	var cmd1, cmd2 tea.Cmd
	m.Input.Title, cmd1 = m.Input.Title.Update(msg)
	m.Input.Description, cmd2 = m.Input.Description.Update(msg)
	
	return m, tea.Batch(cmd1, cmd2) 
}

func (m Model) habitViewModeUpdate(msg tea.Msg) (tea.Model, tea.Cmd) {
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
				m.Help.ShowAll = !m.Help.ShowAll
				break
			case key.Matches(msg, m.habitViewKeys.AddHabit):
				m.Mode = InputMode
				break
		}
	}

	var cmd tea.Cmd
	return m, cmd
}

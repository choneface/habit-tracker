package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)
var mainStyle = lipgloss.NewStyle().
	Margin(1, 2)

var titleBarStyle = lipgloss.NewStyle().
    Bold(true).
    Foreground(lipgloss.Color("#FAFAFA"))

var headerStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#f26363"))

var containerStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("FAFAFA"))

func (m model) View() string {
	content := ""
	content += titleBarStyle.Render(m.Title) + "\n" 
	content += (containerStyle.Render(m.mainContent())) + "\n"
	content += m.helpContent()

	return mainStyle.Render(content)
}


func (m model) mainContent() string {
	switch(m.mode){
	case habit_view_mode:
		return m.habitView()
	case input_mode:
		return m.inputView() 
	}
	panic("Unsupported mode")
}

func (m model) helpContent() string {
	switch (m.mode) {
	case habit_view_mode :
		return m.help.View(m.habitViewKeys)
	case input_mode :
		return m.help.View(inputKeys) 
	}
	panic("Unsupported mode")
}

func (m model) habitView() string {
	return ( headerStyle.Render(m.HabitView.Habits[m.HabitView.Index].Title) ) + "\n\n" + 
	(renderHabitGrid(m.HabitView.Habits[m.HabitView.Index]))
}

func (m model) inputView() string {
	return m.Input.Title.View() + "\n" + m.Input.Description.View() + "\n" 
}

func renderHabitGrid(h habit) string {
	ret := ""
	for index, value := range h.History {
		block := emptyBlock()
		if value > 0 {
			block = filledBlock()
		}
		ret += block + " " 
		if index % 7 == 0 && index > 0 {
			ret += "\n\n"
		}
	}
	return ret 
}

func filledBlock() string {
	var blockStyle= lipgloss.NewStyle().
		Foreground(lipgloss.Color("#6ceb8e"))
	return blockStyle.Render("██")
}

func emptyBlock() string {
	var blockStyle= lipgloss.NewStyle().
		Foreground(lipgloss.Color("#171C24"))
	return blockStyle.Render("██")
}

func (i input) toggleInputFocus() tea.Cmd {
	var blurredStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))

	var cmd tea.Cmd
	if i.FocusIndex == 0 {
		cmd = i.Title.Focus()
		i.Description.Blur()

		i.Title.TextStyle = headerStyle 
		i.Title.PromptStyle = headerStyle

		i.Description.TextStyle = blurredStyle
		i.Description.PromptStyle = blurredStyle
	} else {
		cmd = i.Description.Focus()
		i.Title.Blur()

		i.Description.TextStyle = headerStyle 
		i.Description.PromptStyle = headerStyle

		i.Title.TextStyle = blurredStyle
		i.Title.PromptStyle = blurredStyle
	}
	return cmd 
}

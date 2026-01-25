package main

import (
	"github.com/charmbracelet/lipgloss"
)
var mainStyle = lipgloss.NewStyle().
	AlignHorizontal(lipgloss.Center).
	AlignVertical(lipgloss.Center)

var titleBarStyle = lipgloss.NewStyle().
    Bold(true).
    Foreground(lipgloss.Color("#FAFAFA"))

var headerStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#f26363"))

var helpSectionStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#757575"))

var containerStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("FAFAFA"))

func (m model) View() string {
	content := titleBarStyle.Render(m.Title) + "\n" + (containerStyle.Render(m.habitView())) + "\n" + (m.help.View(m.keys))

	return mainStyle.Render(content)
}

func (m model) habitView() string {
	return ( headerStyle.Render(m.HabitView.Habits[m.HabitView.Index].Title) ) + "\n\n" + 
	(renderHabitView(m.HabitView.Habits[m.HabitView.Index]))
}

func renderHabitView(h habit) string {
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

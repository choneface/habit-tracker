package main

import (
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

func (m Model) View() string {
	content := ""
	content += titleBarStyle.Render(m.Title) + "\n" 
	content += (containerStyle.Render(m.mainContent())) + "\n"
	content += m.helpContent()

	return mainStyle.Render(content)
}


func (m Model) mainContent() string {
	switch(m.Mode){
	case HabitViewMode:
		return m.habitView()
	case InputMode:
		return m.inputView() 
	}
	panic("Unsupported mode")
}

func (m Model) helpContent() string {
	switch (m.Mode) {
	case HabitViewMode :
		return m.Help.View(m.habitViewKeys)
	case InputMode :
		return m.Help.View(inputKeys) 
	}
	panic("Unsupported mode")
}

func (m Model) habitView() string {
	return ( headerStyle.Render(m.HabitView.Habits[m.HabitView.Index].Title) ) + "\n\n" + 
	(renderHabitGrid(m.HabitView.Habits[m.HabitView.Index]))
}

func (m Model) inputView() string {
	return m.Input.Title.View() + "\n" + m.Input.Description.View() + "\n" 
}

func renderHabitGrid(h Habit) string {
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

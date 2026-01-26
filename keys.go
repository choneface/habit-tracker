package main 

import (
	"github.com/charmbracelet/bubbles/key"
)

var keys = habitViewKeyMap {
	Left: key.NewBinding(
		key.WithKeys("left", "h"),
		key.WithHelp("←/h", "move left"),
	),
	Right : key.NewBinding(
		key.WithKeys("right", "l"),
		key.WithHelp("→/l", "move right"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q"),
		key.WithHelp("q", "quit"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	),
	AddHabit: key.NewBinding(
		key.WithKeys("n"),
		key.WithHelp("n", "add habit"),
	),
}

var inputKeys = inputKeyMap {
	Up: key.NewBinding(
		key.WithKeys("up"),
		key.WithHelp("↑", "move up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down"),
		key.WithHelp("↓", "move down"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q"),
		key.WithHelp("q", "quit"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	),
	Submit: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "submit"),
	),
}

type habitViewKeyMap struct {
	Left key.Binding
	Right key.Binding
	Quit key.Binding
	Help key.Binding
	AddHabit key.Binding
}

type inputKeyMap struct {
	Up key.Binding
	Down key.Binding
	Quit key.Binding
	Submit key.Binding
	Help key.Binding
}

func (k habitViewKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{ k.Help, k.Quit }
}

func (k habitViewKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{ k.Left, k.Right, k.AddHabit },
		{ k.Help, k.Quit }, 
	}
}

func (k inputKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{ k.Help, k.Quit }
}

func (k inputKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{ k.Up, k.Down, k.Submit},
		{ k.Help, k.Quit }, 
	}
}

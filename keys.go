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
}

type habitViewKeyMap struct {
	Left key.Binding
	Right key.Binding
	Quit key.Binding
	Help key.Binding
}

func (k habitViewKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{ k.Help, k.Quit }
}

func (k habitViewKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{ k.Left, k.Right },
		{ k.Help, k.Quit }, 
	}
}

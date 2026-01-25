package main 

import (
	"github.com/charmbracelet/bubbles/key"
)

var keys = keyMap {
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
}

type keyMap struct {
	Left key.Binding
	Right key.Binding
	Quit key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{ k.Left, k.Right, k.Quit }
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{ k.Left, k.Right },
		{ k.Quit }, 
	}
}

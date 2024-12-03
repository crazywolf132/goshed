package tui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func helpView() string {
	var s strings.Builder

	s.WriteString(titleStyle.Render("GoShed Help"))
	s.WriteString("\n\n")

	commands := [][]string{
		{"Navigation", ""},
		{"↑/↓", "Move up/down"},
		{"Tab", "Toggle preview"},
		{"Enter", "Confirm selection"},
		{"Esc", "Go back"},
		{"", ""},
		{"Global", ""},
		{"?", "Toggle help"},
		{"Ctrl+c", "Quit"},
	}

	leftWidth := 12
	for _, cmd := range commands {
		if len(cmd) == 1 {
			s.WriteString(selectedStyle.Render(cmd[0]))
			s.WriteString("\n")
			continue
		}
		key := selectedStyle.Render(cmd[0])
		desc := helpStyle.Render(cmd[1])
		padding := strings.Repeat(" ", leftWidth-lipgloss.Width(key))
		s.WriteString(key + padding + desc + "\n")
	}

	return listStyle.Render(s.String())
}

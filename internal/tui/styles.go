package tui

import "github.com/charmbracelet/lipgloss"

var (
	// Colors
	primaryColor   = lipgloss.Color("#7571F9")
	secondaryColor = lipgloss.Color("#5D5B9E")
	successColor   = lipgloss.Color("#04B575")
	warningColor   = lipgloss.Color("#FFA400")
	errorColor     = lipgloss.Color("#FF0056")

	// Base styles
	titleStyle = lipgloss.NewStyle().
			Foreground(primaryColor).
			Bold(true).
			Padding(1, 0, 1, 2).
			BorderStyle(lipgloss.RoundedBorder()).
			BorderBottom(true)

	selectedStyle = lipgloss.NewStyle().
			Foreground(primaryColor).
			Bold(true)

	// Input styles
	inputStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(primaryColor).
			Padding(0, 1)

	// List styles
	listStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(primaryColor).
			Padding(1, 2)

	listItemStyle = lipgloss.NewStyle().
			PaddingLeft(2)

	selectedListItemStyle = lipgloss.NewStyle().
				Foreground(primaryColor).
				PaddingLeft(2).
				Bold(true)

	// Help style
	helpStyle = lipgloss.NewStyle().
			Foreground(secondaryColor).
			Italic(true)

	// Status styles
	successStyle = lipgloss.NewStyle().
			Foreground(successColor).
			Bold(true)

	warningStyle = lipgloss.NewStyle().
			Foreground(warningColor).
			Bold(true)

	errorStyle = lipgloss.NewStyle().
			Foreground(errorColor).
			Bold(true)
)

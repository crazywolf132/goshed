package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/lipgloss"
)

type Preview struct {
	viewport viewport.Model
	content  string
	visible  bool
	width    int
	height   int
	style    lipgloss.Style
}

func NewPreview() Preview {
	return Preview{
		viewport: viewport.New(0, 0),
		visible:  false,
		style:    lipgloss.NewStyle().BorderStyle(lipgloss.RoundedBorder()),
	}
}

func (p *Preview) SetSize(width, height int) {
	p.width = width
	p.height = height
	p.viewport.Width = width
	p.viewport.Height = height
}

func (p *Preview) SetContent(files map[string]string) {
	var sb strings.Builder
	fileStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#7571F9")).
		Bold(true)

	for filename, content := range files {
		sb.WriteString(fileStyle.Render(fmt.Sprintf("// %s", filename)))
		sb.WriteString("\n")
		sb.WriteString(content)
		sb.WriteString("\n\n")
	}

	p.content = sb.String()
	p.viewport.SetContent(p.content)
}

func (p Preview) View() string {
	if !p.visible {
		return ""
	}

	return p.style.Copy().
		Width(p.width).
		Height(p.height).
		Render(p.viewport.View())
}

package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/crazywolf132/goshed/internal/model"
	"github.com/crazywolf132/goshed/internal/project"
	"github.com/crazywolf132/goshed/internal/template"
)

type state int

const (
	stateProjectName state = iota
	stateTemplate
	stateTags
	stateConfirm
)

type Model struct {
	state       state
	projectName textinput.Model
	templates   list.Model
	tags        textinput.Model
	err         error
	quitting    bool
	spinner     spinner.Model
	width       int
	height      int
	preview     Preview
	showHelp    bool
}

type item struct {
	name, desc string
}

func (i item) Title() string       { return i.name }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.name }

func InitialModel() Model {
	// Project name input
	pn := textinput.New()
	pn.Placeholder = "Enter project name"
	pn.Focus()
	pn.CharLimit = 50
	pn.Width = 40

	// Template selection
	templates := template.List()
	items := make([]list.Item, 0, len(templates))
	for name, tmpl := range templates {
		items = append(items, item{name: name, desc: tmpl.Description})
	}

	delegate := list.NewDefaultDelegate()
	delegate.Styles.SelectedTitle = selectedListItemStyle
	delegate.Styles.SelectedDesc = selectedListItemStyle.Copy().Italic(true)

	templateList := list.New(items, delegate, 0, 0)
	templateList.Title = "Select Template"
	templateList.SetShowHelp(false)
	templateList.SetFilteringEnabled(true)
	templateList.Styles.Title = titleStyle.Copy().MarginLeft(2)

	// Tags input
	tags := textinput.New()
	tags.Placeholder = "Enter tags (comma-separated)"
	tags.CharLimit = 100
	tags.Width = 40

	// Spinner for loading states
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("#7571F9"))

	return Model{
		state:       stateProjectName,
		projectName: pn,
		templates:   templateList,
		tags:        tags,
		spinner:     s,
		preview:     NewPreview(),
		showHelp:    false,
	}
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit
		case "enter":
			switch m.state {
			case stateProjectName:
				if m.projectName.Value() != "" {
					m.state = stateTemplate
				}
			case stateTemplate:
				m.state = stateTags
				m.tags.Focus()
			case stateTags:
				m.state = stateConfirm
			case stateConfirm:
				if msg.String() == "y" {
					return m, m.createProject
				}
				m.quitting = true
				return m, tea.Quit
			}
		case "esc":
			if m.state > stateProjectName {
				m.state--
				if m.state == stateProjectName {
					m.projectName.Focus()
					m.tags.Blur()
				}
			}
		case "tab":
			if m.state == stateTemplate {
				if selected, ok := m.templates.SelectedItem().(item); ok {
					if tmpl, err := template.Get(selected.name); err == nil {
						m.preview.SetContent(tmpl.Files)
						m.preview.visible = !m.preview.visible
					}
				}
			}
		case "?":
			m.showHelp = !m.showHelp
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		if m.preview.visible {
			m.templates.SetSize(msg.Width/2-4, msg.Height-8)
			m.preview.SetSize(msg.Width/2-4, msg.Height-8)
		} else {
			m.templates.SetSize(msg.Width-4, msg.Height-8)
		}
		return m, nil
	}

	switch m.state {
	case stateProjectName:
		m.projectName, cmd = m.projectName.Update(msg)
		return m, cmd
	case stateTemplate:
		m.templates, cmd = m.templates.Update(msg)
		return m, cmd
	case stateTags:
		m.tags, cmd = m.tags.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m Model) View() string {
	if m.quitting {
		return goodbyeView()
	}

	var s strings.Builder

	// Header
	header := titleStyle.Render("🏗  GoShed Project Creator")
	s.WriteString(lipgloss.Place(m.width, 3, lipgloss.Center, lipgloss.Center, header))
	s.WriteString("\n\n")

	// Main content with optional preview
	if m.state == stateTemplate && m.preview.visible {
		left := m.templates.View()
		right := m.preview.View()
		content := lipgloss.JoinHorizontal(lipgloss.Top, left, right)
		s.WriteString(lipgloss.Place(m.width, m.height-6, lipgloss.Center, lipgloss.Center, content))
	} else {
		content := m.getStateView()
		s.WriteString(lipgloss.Place(m.width, m.height-6, lipgloss.Center, lipgloss.Center, content))
	}

	// Error message
	if m.err != nil {
		errMsg := errorStyle.Render(fmt.Sprintf("Error: %v", m.err))
		s.WriteString("\n" + lipgloss.Place(m.width, 1, lipgloss.Center, lipgloss.Center, errMsg))
	}

	// Help text
	help := helpStyle.Render(getHelp(m.state))
	s.WriteString("\n" + lipgloss.Place(m.width, 2, lipgloss.Center, lipgloss.Bottom, help))

	return s.String()
}

func (m Model) getStateView() string {
	switch m.state {
	case stateProjectName:
		return inputStyle.Render(
			lipgloss.JoinVertical(lipgloss.Left,
				"Project Name:",
				m.projectName.View(),
			),
		)

	case stateTemplate:
		return listStyle.Render(m.templates.View())

	case stateTags:
		return inputStyle.Render(
			lipgloss.JoinVertical(lipgloss.Left,
				"Tags (comma-separated):",
				m.tags.View(),
			),
		)

	case stateConfirm:
		return confirmationView(m)

	default:
		return ""
	}
}

func confirmationView(m Model) string {
	var s strings.Builder
	s.WriteString("Create project with these settings?\n\n")

	details := []string{
		fmt.Sprintf("%s %s", selectedStyle.Render("Name:"), m.projectName.Value()),
		fmt.Sprintf("%s %s", selectedStyle.Render("Template:"), m.templates.SelectedItem().(item).name),
	}

	if m.tags.Value() != "" {
		details = append(details, fmt.Sprintf("%s %s", selectedStyle.Render("Tags:"), m.tags.Value()))
	}

	s.WriteString(lipgloss.JoinVertical(lipgloss.Left, details...))
	s.WriteString("\n\nPress 'y' to create, 'n' to cancel")

	return inputStyle.Render(s.String())
}

func goodbyeView() string {
	return lipgloss.NewStyle().
		Foreground(successColor).
		Bold(true).
		Render("Thanks for using GoShed! 👋\n")
}

func getHelp(s state) string {
	switch s {
	case stateProjectName:
		return "Enter project name • Ctrl+c to quit"
	case stateTemplate:
		return "↑/↓ to select • Tab to preview • Enter to confirm • Esc to go back • ? for help • Ctrl+c to quit"
	case stateTags:
		return "Enter tags • Enter to confirm • Esc to go back • Ctrl+c to quit"
	case stateConfirm:
		return "y/n to confirm • Esc to go back • Ctrl+c to quit"
	default:
		return ""
	}
}

func (m Model) createProject() tea.Msg {
	p := &model.Project{
		Name:     m.projectName.Value(),
		Template: m.templates.SelectedItem().(item).name,
		Tags:     strings.Split(m.tags.Value(), ","),
	}

	if err := project.Create(p); err != nil {
		m.err = err
		return nil
	}

	m.quitting = true
	return tea.Quit()
}

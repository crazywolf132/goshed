package styles

import (
	"github.com/fatih/color"
)

var (
	// Headers and titles
	Title  = color.New(color.FgHiCyan, color.Bold).SprintfFunc()
	Header = color.New(color.FgCyan).SprintfFunc()

	// Status and results
	Success = color.New(color.FgGreen).SprintfFunc()
	Warning = color.New(color.FgYellow).SprintfFunc()
	Error   = color.New(color.FgRed).SprintfFunc()

	// Project details
	ProjectName = color.New(color.FgHiBlue, color.Bold).SprintfFunc()
	FieldName   = color.New(color.FgBlue).SprintfFunc()
	TagText     = color.New(color.FgHiMagenta).SprintfFunc()
	TimeText    = color.New(color.FgHiGreen).SprintfFunc()
)

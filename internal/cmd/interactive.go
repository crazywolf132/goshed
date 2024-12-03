package cmd

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/crazywolf132/goshed/internal/tui"
	"github.com/spf13/cobra"
)

var interactiveCmd = &cobra.Command{
	Use:   "interactive",
	Short: "Start interactive mode",
	Long: `Start GoShed in interactive mode with a terminal UI.
Example: goshed interactive`,
	Aliases: []string{"i"},
	RunE: func(cmd *cobra.Command, args []string) error {
		p := tea.NewProgram(tui.InitialModel())
		if _, err := p.Run(); err != nil {
			return fmt.Errorf("failed to start interactive mode: %w", err)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(interactiveCmd)
}

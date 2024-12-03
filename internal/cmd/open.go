package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/crazywolf132/goshed/internal/project"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Open a playground in your editor",
	Long: `Open a playground in your configured editor.
Example: goshed open -n myproject`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if projectName == "" {
			return fmt.Errorf("project name is required")
		}

		// Get project
		p, err := project.Get(projectName)
		if err != nil {
			return fmt.Errorf("failed to get project: %w", err)
		}

		// Update last accessed time
		p.LastAccessed = time.Now()
		if err := project.Update(p); err != nil {
			return fmt.Errorf("failed to update project: %w", err)
		}

		// Get editor from config
		editor := viper.GetString("editor")
		if editor == "" {
			editor = "code" // Default to VS Code
		}

		// Open project in editor
		execCmd := exec.Command(editor, p.Path)
		execCmd.Stdin = os.Stdin
		execCmd.Stdout = os.Stdout
		execCmd.Stderr = os.Stderr

		if err := execCmd.Run(); err != nil {
			return fmt.Errorf("failed to open editor: %w", err)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(openCmd)
	openCmd.Flags().StringVarP(&projectName, "name", "n", "", "Name of the playground to open (required)")
	openCmd.MarkFlagRequired("name")
}

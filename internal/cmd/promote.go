package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/crazywolf132/goshed/internal/project"
	"github.com/crazywolf132/goshed/internal/styles"
	"github.com/spf13/cobra"
)

var (
	destination string
)

var promoteCmd = &cobra.Command{
	Use:   "promote",
	Short: "Promote a playground to a full project",
	Long: `Promote a playground to a full standalone project.
This will move the project out of GoShed's management.
Example: goshed promote -n myproject -d ~/projects/myproject`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if projectName == "" {
			return fmt.Errorf("project name is required")
		}

		// Get project
		p, err := project.Get(projectName)
		if err != nil {
			return fmt.Errorf("failed to get project: %w", err)
		}

		// If destination is not specified, use current directory
		if destination == "" {
			cwd, err := os.Getwd()
			if err != nil {
				return fmt.Errorf("failed to get current directory: %w", err)
			}
			destination = filepath.Join(cwd, p.Name)
		}

		// Ensure destination doesn't exist
		if _, err := os.Stat(destination); !os.IsNotExist(err) {
			return fmt.Errorf("destination %s already exists", destination)
		}

		// Create destination directory
		if err := os.MkdirAll(destination, 0755); err != nil {
			return fmt.Errorf("failed to create destination directory: %w", err)
		}

		// Copy project files to destination
		if err := project.CopyTo(p, destination); err != nil {
			return fmt.Errorf("failed to copy project: %w", err)
		}

		// Remove project from GoShed
		if err := project.Remove(p.Name); err != nil {
			return fmt.Errorf("failed to remove project from GoShed: %w", err)
		}

		fmt.Printf("%s %s to %s\n",
			styles.Success("Successfully promoted"),
			styles.ProjectName(p.Name),
			styles.Header(destination),
		)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(promoteCmd)
	promoteCmd.Flags().StringVarP(&projectName, "name", "n", "", "Name of the playground to promote (required)")
	promoteCmd.Flags().StringVarP(&destination, "destination", "d", "", "Destination directory (defaults to current directory)")
	promoteCmd.MarkFlagRequired("name")
}

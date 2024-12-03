package cmd

import (
	"fmt"

	"github.com/crazywolf132/goshed/internal/project"
	"github.com/crazywolf132/goshed/internal/styles"
	"github.com/spf13/cobra"
)

var (
	noteText string
)

var notesCmd = &cobra.Command{
	Use:   "notes",
	Short: "Manage project notes",
	Long: `Add or view notes for a project.
Example: goshed notes -n myproject [-t "My note text"]`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if projectName == "" {
			return fmt.Errorf("project name is required")
		}

		p, err := project.Get(projectName)
		if err != nil {
			return fmt.Errorf("failed to get project: %w", err)
		}

		// If note text is provided, update the notes
		if noteText != "" {
			p.Notes = noteText
			if err := project.Update(p); err != nil {
				return fmt.Errorf("failed to update project notes: %w", err)
			}
			fmt.Printf("%s %s\n", styles.Success("Updated notes for"), styles.ProjectName(p.Name))
		}

		// Display current notes
		if p.Notes != "" {
			fmt.Printf("%s %s\n", styles.FieldName("Notes:"), p.Notes)
		} else {
			fmt.Printf("%s\n", styles.Warning("No notes found"))
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(notesCmd)
	notesCmd.Flags().StringVarP(&projectName, "name", "n", "", "Name of the playground (required)")
	notesCmd.Flags().StringVarP(&noteText, "text", "t", "", "Note text to add")
	notesCmd.MarkFlagRequired("name")
}

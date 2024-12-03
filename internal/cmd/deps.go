package cmd

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/crazywolf132/goshed/internal/project"
	"github.com/crazywolf132/goshed/internal/styles"
	"github.com/spf13/cobra"
)

var depsCmd = &cobra.Command{
	Use:   "deps",
	Short: "Show project dependencies",
	Long: `Show and analyze project dependencies.
Example: goshed deps -n myproject`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if projectName == "" {
			return fmt.Errorf("project name is required")
		}

		p, err := project.Get(projectName)
		if err != nil {
			return fmt.Errorf("failed to get project: %w", err)
		}

		// Run go list -m all
		execCmd := exec.Command("go", "list", "-m", "all")
		execCmd.Dir = p.Path
		output, err := execCmd.Output()
		if err != nil {
			return fmt.Errorf("failed to list dependencies: %w", err)
		}

		// Parse and display dependencies
		deps := strings.Split(string(output), "\n")
		if len(deps) <= 1 {
			fmt.Printf("%s\n", styles.Warning("No dependencies found"))
			return nil
		}

		fmt.Printf("%s %s\n\n", styles.Title("Dependencies for"), styles.ProjectName(p.Name))
		for i, dep := range deps {
			if i == 0 || dep == "" { // Skip module name and empty lines
				continue
			}

			// Check for updates
			execCmd = exec.Command("go", "list", "-m", "-u", dep)
			execCmd.Dir = p.Path
			updateOutput, err := execCmd.Output()
			if err == nil && strings.Contains(string(updateOutput), "[") {
				fmt.Printf("%s %s\n", styles.FieldName(dep), styles.Warning("(update available)"))
			} else {
				fmt.Printf("%s\n", styles.FieldName(dep))
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(depsCmd)
	depsCmd.Flags().StringVarP(&projectName, "name", "n", "", "Name of the playground (required)")
	depsCmd.MarkFlagRequired("name")
}

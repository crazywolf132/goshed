package cmd

import (
	"fmt"
	"time"

	"github.com/crazywolf132/goshed/internal/model"
	"github.com/crazywolf132/goshed/internal/project"
	"github.com/crazywolf132/goshed/internal/styles"
	"github.com/spf13/cobra"
)

var (
	projectName  string
	templateName string
	tags         []string
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new Go playground",
	Long: `Create a new Go playground with the specified name and template.
Example: goshed create -n myproject -t basic`,
	RunE: func(cmd *cobra.Command, args []string) error {
		p := &model.Project{
			Name:         projectName,
			Created:      time.Now(),
			LastAccessed: time.Now(),
			Template:     templateName,
			Tags:         tags,
		}

		if err := project.Create(p); err != nil {
			return fmt.Errorf("%s: %w", styles.Error("%s", "Failed to create project"), err)
		}

		fmt.Printf("%s %s\n",
			styles.Success("%s", "Created new playground:"),
			styles.ProjectName("%s", projectName),
		)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.Flags().StringVarP(&projectName, "name", "n", "", "Name of the playground (required)")
	createCmd.Flags().StringVarP(&templateName, "template", "t", "basic", "Template to use (basic, web, cli)")
	createCmd.Flags().StringSliceVarP(&tags, "tags", "", nil, "Tags to categorize the playground")

	createCmd.MarkFlagRequired("name")
}

package cmd

import (
	"fmt"

	"github.com/crazywolf132/goshed/internal/styles"
	tmpl "github.com/crazywolf132/goshed/internal/template"
	"github.com/spf13/cobra"
)

var templatesCmd = &cobra.Command{
	Use:   "templates",
	Short: "List available templates",
	Long: `List all available project templates.
Example: goshed templates`,
	RunE: func(cmd *cobra.Command, args []string) error {
		templates := tmpl.List()

		fmt.Printf("%s\n\n", styles.Title("Available Templates:"))
		for name, t := range templates {
			fmt.Printf("%s %s\n", styles.ProjectName(name), styles.Header("- %s", t.Description))
			if len(t.Dependencies) > 0 {
				fmt.Printf("  %s\n", styles.FieldName("Dependencies:"))
				for _, dep := range t.Dependencies {
					fmt.Printf("    - %s\n", dep)
				}
			}
			fmt.Println()
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(templatesCmd)
}

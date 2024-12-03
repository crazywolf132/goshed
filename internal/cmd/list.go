package cmd

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/crazywolf132/goshed/internal/model"
	"github.com/crazywolf132/goshed/internal/project"
	"github.com/crazywolf132/goshed/internal/styles"
	"github.com/spf13/cobra"
)

var (
	showTags    bool
	filterTag   string
	sortBy      string
	reverseSort bool
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all playgrounds",
	Long: `List all playgrounds with their details.
Example: goshed list --tags --sort=accessed`,
	RunE: func(cmd *cobra.Command, args []string) error {
		projects, err := project.List()
		if err != nil {
			return fmt.Errorf("failed to list projects: %w", err)
		}

		// Filter by tag if specified
		if filterTag != "" {
			filtered := make([]*model.Project, 0)
			for _, p := range projects {
				for _, t := range p.Tags {
					if t == filterTag {
						filtered = append(filtered, p)
						break
					}
				}
			}
			projects = filtered
		}

		// Sort projects
		sort.Slice(projects, func(i, j int) bool {
			var result bool
			switch sortBy {
			case "name":
				result = projects[i].Name < projects[j].Name
			case "created":
				result = projects[i].Created.Before(projects[j].Created)
			case "accessed":
				result = projects[i].LastAccessed.Before(projects[j].LastAccessed)
			default:
				result = projects[i].Name < projects[j].Name
			}
			if reverseSort {
				return !result
			}
			return result
		})

		// Print projects
		if len(projects) == 0 {
			fmt.Println(styles.Warning("No playgrounds found"))
			return nil
		}

		fmt.Printf("%s\n\n", styles.Title("Found %d playgrounds:", len(projects)))
		for _, p := range projects {
			fmt.Printf("%s %s\n", styles.FieldName("Name:"), styles.ProjectName(p.Name))
			fmt.Printf("  %s %s\n", styles.FieldName("Template:"), p.Template)
			fmt.Printf("  %s %s\n", styles.FieldName("Created:"), styles.TimeText(p.Created.Format(time.RFC3339)))
			fmt.Printf("  %s %s\n", styles.FieldName("Accessed:"), styles.TimeText(p.LastAccessed.Format(time.RFC3339)))

			// Add Git status
			if status, err := project.GetGitStatus(p); err == nil {
				if status.Initialized {
					gitStatus := "Clean"
					if !status.Clean {
						gitStatus = styles.Warning("Modified")
					}
					fmt.Printf("  %s %s on %s",
						styles.FieldName("Git:"),
						gitStatus,
						styles.Header(status.Branch),
					)
					if status.Remote != "" {
						fmt.Printf(" (%s)", status.Remote)
					}
					fmt.Println()
				}
			}

			if showTags && len(p.Tags) > 0 {
				tagList := make([]string, len(p.Tags))
				for i, tag := range p.Tags {
					tagList[i] = styles.TagText("%s", tag)
				}
				fmt.Printf("  %s %s\n", styles.FieldName("Tags:"), strings.Join(tagList, ", "))
			}
			if p.Notes != "" {
				fmt.Printf("  %s %s\n", styles.FieldName("Notes:"), p.Notes)
			}
			fmt.Println()
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVar(&showTags, "tags", false, "Show project tags")
	listCmd.Flags().StringVar(&filterTag, "filter-tag", "", "Filter projects by tag")
	listCmd.Flags().StringVar(&sortBy, "sort", "name", "Sort by: name, created, accessed")
	listCmd.Flags().BoolVar(&reverseSort, "reverse", false, "Reverse sort order")
}

package cmd

import (
	"fmt"
	"time"

	"github.com/crazywolf132/goshed/internal/project"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	olderThan string
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Clean up old playgrounds",
	Long: `Remove playgrounds that haven't been accessed for a specified duration.
Example: goshed clean --older-than 720h`,
	RunE: func(cmd *cobra.Command, args []string) error {
		duration, err := time.ParseDuration(olderThan)
		if err != nil {
			return fmt.Errorf("invalid duration format: %w", err)
		}

		// Get all projects
		projects, err := project.List()
		if err != nil {
			return fmt.Errorf("failed to list projects: %w", err)
		}

		cutoff := time.Now().Add(-duration)
		cleaned := 0

		for _, p := range projects {
			if p.LastAccessed.Before(cutoff) {
				if err := project.Remove(p.Name); err != nil {
					fmt.Printf("Warning: failed to remove %s: %v\n", p.Name, err)
					continue
				}
				cleaned++
				fmt.Printf("Removed %s (last accessed: %s)\n", p.Name, p.LastAccessed.Format(time.RFC3339))
			}
		}

		fmt.Printf("Cleaned up %d projects\n", cleaned)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)
	cleanCmd.Flags().StringVar(&olderThan, "older-than", viper.GetString("cleanup.older_than"), "Remove projects older than this duration (e.g., 720h)")
}

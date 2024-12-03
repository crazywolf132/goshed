package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/crazywolf132/goshed/internal/config"
	"github.com/crazywolf132/goshed/internal/styles"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	workspaceName string
	setDefault    bool
)

var workspaceCmd = &cobra.Command{
	Use:   "workspace",
	Short: "Manage workspaces",
	Long: `Create and switch between workspaces.
Example: goshed workspace create mywork`,
}

var workspaceCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new workspace",
	RunE: func(cmd *cobra.Command, args []string) error {
		if workspaceName == "" {
			return fmt.Errorf("workspace name is required")
		}

		wsPath := filepath.Join(config.ConfigDir, "workspaces", workspaceName)
		if err := os.MkdirAll(wsPath, 0755); err != nil {
			return fmt.Errorf("failed to create workspace: %w", err)
		}

		if setDefault {
			viper.Set("workspace", workspaceName)
			viper.WriteConfig()
		}

		fmt.Printf("%s %s\n", styles.Success("Created workspace:"), styles.ProjectName(workspaceName))
		return nil
	},
}

var workspaceListCmd = &cobra.Command{
	Use:   "list",
	Short: "List available workspaces",
	RunE: func(cmd *cobra.Command, args []string) error {
		wsDir := filepath.Join(config.ConfigDir, "workspaces")
		entries, err := os.ReadDir(wsDir)
		if err != nil && !os.IsNotExist(err) {
			return fmt.Errorf("failed to read workspaces: %w", err)
		}

		current := viper.GetString("workspace")
		fmt.Printf("%s\n\n", styles.Title("Available Workspaces:"))

		for _, entry := range entries {
			if entry.IsDir() {
				name := entry.Name()
				if name == current {
					fmt.Printf("%s %s\n", styles.ProjectName(name), styles.Success("(current)"))
				} else {
					fmt.Printf("%s\n", styles.ProjectName(name))
				}
			}
		}

		return nil
	},
}

var workspaceSwitchCmd = &cobra.Command{
	Use:   "switch",
	Short: "Switch to a different workspace",
	RunE: func(cmd *cobra.Command, args []string) error {
		if workspaceName == "" {
			return fmt.Errorf("workspace name is required")
		}

		wsPath := filepath.Join(config.ConfigDir, "workspaces", workspaceName)
		if _, err := os.Stat(wsPath); os.IsNotExist(err) {
			return fmt.Errorf("workspace %s does not exist", workspaceName)
		}

		viper.Set("workspace", workspaceName)
		viper.WriteConfig()

		fmt.Printf("%s %s\n", styles.Success("Switched to workspace:"), styles.ProjectName(workspaceName))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(workspaceCmd)
	workspaceCmd.AddCommand(workspaceCreateCmd)
	workspaceCmd.AddCommand(workspaceListCmd)
	workspaceCmd.AddCommand(workspaceSwitchCmd)

	workspaceCreateCmd.Flags().StringVarP(&workspaceName, "name", "n", "", "Name of the workspace")
	workspaceCreateCmd.Flags().BoolVarP(&setDefault, "default", "d", false, "Set as default workspace")
	workspaceCreateCmd.MarkFlagRequired("name")

	workspaceSwitchCmd.Flags().StringVarP(&workspaceName, "name", "n", "", "Name of the workspace")
	workspaceSwitchCmd.MarkFlagRequired("name")
}

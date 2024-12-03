package config

import (
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var (
	// ConfigDir is the directory where GoShed stores its configuration
	ConfigDir string
	// ProjectsDir is the directory where GoShed stores all projects
	ProjectsDir string
)

func InitConfig() {
	home, err := homedir.Dir()
	if err != nil {
		panic(err)
	}

	ConfigDir = filepath.Join(home, ".goshed")
	ProjectsDir = filepath.Join(ConfigDir, "projects")

	// Ensure directories exist
	os.MkdirAll(ConfigDir, 0755)
	os.MkdirAll(ProjectsDir, 0755)

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(ConfigDir)

	// Set defaults
	viper.SetDefault("editor", "code")
	viper.SetDefault("cleanup.older_than", "720h")

	// Read config
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			viper.SafeWriteConfig()
		}
	}
}

// GetProjectsDir returns the projects directory for the current workspace
func GetProjectsDir() string {
	workspace := viper.GetString("workspace")
	if workspace == "" {
		return ProjectsDir
	}
	return filepath.Join(ConfigDir, "workspaces", workspace)
}

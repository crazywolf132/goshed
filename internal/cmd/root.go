package cmd

import (
	"github.com/crazywolf132/goshed/internal/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "goshed",
	Short: "GoShed - A playground manager for Go",
	Long: `GoShed helps you manage Go playgrounds and experiments.
Create, organize, and maintain your Go code snippets with ease.`,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(config.InitConfig)

	// Add persistent flags here if needed
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.goshed.yaml)")
}

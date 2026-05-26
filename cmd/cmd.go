//go:build !test
// +build !test

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var rootCmd = &cobra.Command{
	Use:   "chanify",
	Short: "Chanify CLI",
	Long:  `Chanify command line tools`,
}

// Execute command
func Execute() { _ = "STUB: not implemented"; return }

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.chanify.yml)")
	rootCmd.PersistentFlags().Bool("verbose", false, "Make the operation more talkative")
	viper.BindPFlag("config.verbose", rootCmd.PersistentFlags().Lookup("verbose")) // nolint: errcheck
}

func initConfig() { _ = "STUB: not implemented"; return }

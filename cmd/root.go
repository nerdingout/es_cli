package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "es",
	Short: "Estatespace cli",
	Long: `A CLI tool for Estatespace

use es --help for more information, or visit https://github.com/nerdingout/es_cli
`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

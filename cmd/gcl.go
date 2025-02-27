package cmd

import (
	"github.com/nerdingout/es_cli/descriptions"
	"github.com/nerdingout/es_cli/handlers"
	"github.com/spf13/cobra"
)

var gclCmd = &cobra.Command{
	Use:   "gcl",
	Short: descriptions.CleanLocalShort,
	Long:  descriptions.CleanLocalLong,
	Run:   handlers.GCLHandler,
}

func init() {
	rootCmd.AddCommand(gclCmd)
}

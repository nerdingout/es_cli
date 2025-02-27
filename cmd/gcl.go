package cmd

import (
	"github.com/nerdingout/es_cli/descriptions"
	"github.com/nerdingout/es_cli/handlers"
	"github.com/spf13/cobra"
)

var gclCmd = &cobra.Command{
	Use:   "gcl",
	Short: descriptions.GclShort,
	Long:  descriptions.GclLong,
	Run:   handlers.GCLHandler,
}

func init() {
	rootCmd.AddCommand(gclCmd)
}

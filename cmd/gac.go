package cmd

import (
	"github.com/nerdingout/es_cli/descriptions"
	"github.com/nerdingout/es_cli/handlers"
	"github.com/spf13/cobra"
)

var gacCmd = &cobra.Command{
	Use:   "gac",
	Short: descriptions.GacShort,
	Long:  descriptions.GacLong,
	Run:   handlers.GACHandler,
}

func init() {
	rootCmd.AddCommand(gacCmd)
}

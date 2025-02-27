package cmd

import (
	"github.com/nerdingout/es_cli/descriptions"
	"github.com/nerdingout/es_cli/handlers"
	"github.com/spf13/cobra"
)

var genCmd = &cobra.Command{
	Use:   "rhf",
	Short: descriptions.RhfShort,
	Long:  descriptions.RhfLong,
	Run:   handlers.RHFHandler,
}

func init() {
	rootCmd.AddCommand(genCmd)
}

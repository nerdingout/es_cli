package cmd

import (
	"github.com/nerdingout/es_cli/descriptions"
	"github.com/nerdingout/es_cli/handlers"
	"github.com/spf13/cobra"
)

var rfcCmd = &cobra.Command{
	Use:   "rfc",
	Short: descriptions.RfcShort,
	Long:  descriptions.RfcLong,
	Run:   handlers.RFCHandler,
}

func init() {
	rootCmd.AddCommand(rfcCmd)
}

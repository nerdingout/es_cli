package cmd

import (
	"github.com/nerdingout/es_cli/descriptions"
	"github.com/nerdingout/es_cli/handlers"

	"github.com/spf13/cobra"
)

var ttlCmd = &cobra.Command{
	Use:   "ttl",
	Short: descriptions.TtlShort,
	Long:  descriptions.TtlLong,
	Run:   handlers.TTLHandler,
}

func init() {
	rootCmd.AddCommand(ttlCmd)
}

package cmd

import (
	"github.com/nerdingout/es_cli/descriptions"
	"github.com/nerdingout/es_cli/handlers"
	"github.com/spf13/cobra"
)

var jiraCmd = &cobra.Command{
	Use:   "jtp",
	Short: descriptions.JtpShort,
	Long:  descriptions.JtpLong,
	Run:   handlers.JTPHandler,
}

func init() {
	rootCmd.AddCommand(jiraCmd)
}

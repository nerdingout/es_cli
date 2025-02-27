package cmd

import (
	"github.com/nerdingout/es_cli/descriptions"
	"github.com/nerdingout/es_cli/handlers"
	"github.com/spf13/cobra"
)

var jiraBlCmd = &cobra.Command{
	Use:   "jbl",
	Short: descriptions.JblShort,
	Long:  descriptions.JblLong,
	Run:   handlers.JBLHandler,
}

func init() {
	rootCmd.AddCommand(jiraBlCmd)
	jiraBlCmd.Flags().StringP("user", "u", "", "Backlog for specific user")
}

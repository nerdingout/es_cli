package cmd

import (
	"fmt"
	"github.com/nerdingout/es_cli/descriptions"
	"os/exec"
	"regexp"

	"github.com/spf13/cobra"
)

var jiraCmd = &cobra.Command{
	Use:   "jtp",
	Short: "jira: Will open the jira ticket for the current branch",
	Long:  descriptions.JTPDesc,
	Run: func(cmd *cobra.Command, args []string) {
		getBranchNameCmd := exec.Command("git", "symbolic-ref", "--short", "-q", "HEAD")
		branchName, err := getBranchNameCmd.Output()

		if err != nil {
			fmt.Println("❌ Error getting current directory")
			return
		}

		if err != nil {
			fmt.Println("Error getting branch name")
			return
		}

		pattern := `ES-[0-9]+`
		re := regexp.MustCompile(pattern)
		matches := re.FindAllStringSubmatch(string(branchName), -1)

		if len(matches) == 0 {
			fmt.Println("❌ Branch name does not match ES-<>")
			return
		}

		jiraTicketNumber := matches[0][0]
		openJiraTicketCmd := exec.Command("open", "https://estatespace.atlassian.net/browse/"+jiraTicketNumber)

		if err := openJiraTicketCmd.Run(); err != nil {
			fmt.Println("❌ Error opening Jira ticket")
			return
		}

		fmt.Println("✅ Opened Jira ticket in browser")
	},
}

func init() {
	rootCmd.AddCommand(jiraCmd)
}

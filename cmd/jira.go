package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"

	"github.com/spf13/cobra"
)

var jiraCmd = &cobra.Command{
	Use:   "jtp",
	Short: "jira: Will open the jira ticket for the current branch",
	Long: `
Will use the current branch name, extract the ticket number and open the ticket in the browser.

Example:

es jtp

Will open the ticket for the branch name ES-1234`,
	Run: func(cmd *cobra.Command, args []string) {
		getBranchNameCmd := exec.Command("git", "symbolic-ref", "--short", "-q", "HEAD")
		branchName, err := getBranchNameCmd.Output()

		currentDir, err := os.Getwd()

		if err != nil {
			fmt.Println("❌ Error getting current directory")
			return
		}

		fmt.Println("current dir is ==>", currentDir)

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

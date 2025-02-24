package cmd

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

var gacCmd = &cobra.Command{
	Use:   "gac",
	Short: "git: Adds and commits your files with a message",
	Long: `
Will stage all files and commit wth a message starting with the Jira ticket number. 

Example:

es gac "this is a message" 

will stage all files and commit with a message like: [ES-1234]: this is a message
`,

	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			fmt.Printf("❌ Error: No message provided\n\nExample:\n\n$ es_cli gac \"this is a message\"\n")
			return
		}

		getBranchNameCmd := exec.Command("git", "symbolic-ref", "--short", "-q", "HEAD")
		branchName, err := getBranchNameCmd.Output()

		if err != nil {
			fmt.Println("❌ Error getting branch name")
		}

		branchNameString := strings.TrimSpace(string(branchName))

		re := regexp.MustCompile(`ES-[0-9]+`)
		matches := re.FindAllStringSubmatch(branchNameString, -1)

		if len(matches) == 0 {
			fmt.Println("❌ Branch name does not match ES-<> pattern")
			return
		}

		gitMessage := matches[0][0] + ": " + args[0]
		addAllCmd := exec.Command("git", "add", ".")

		if err := addAllCmd.Run(); err != nil {
			fmt.Println("❌ Error adding all files")
			return
		}

		commitCmd := exec.Command("git", "commit", "-m", gitMessage)

		if err := commitCmd.Run(); err != nil {
			fmt.Println("❌ Error committing files")
			return
		}

		fmt.Printf("✅ Committed with message:\n\n%s\n", gitMessage)
		return
	},
}

func init() {
	rootCmd.AddCommand(gacCmd)
}

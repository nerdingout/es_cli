package cmd

import (
	"fmt"
	"github.com/nerdingout/es_cli/descriptions"
	"github.com/nerdingout/es_cli/templates"
	"os"

	"github.com/spf13/cobra"
)

var genCmd = &cobra.Command{
	Use:   "rhf",
	Short: "react: A command to generate a helper function with test.",
	Long:  descriptions.RHFDesc,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			fmt.Println("❌ Error: No function name provided, run with --help for more information")
			return
		}

		functionName := args[0]
		fileName := functionName + ".js"
		testFileName := functionName + ".test.js"

		fileContents := fmt.Sprintf(templates.Function, functionName, functionName, functionName)
		testFileContents := fmt.Sprintf(templates.FunctionTest, functionName, fileName, functionName)

		if err := os.WriteFile("index.js", []byte(fileContents), 0644); err != nil {
			fmt.Println("❌ Error creating function file")
		}

		if err := os.WriteFile(testFileName, []byte(testFileContents), 0644); err != nil {
			fmt.Println("❌ Error creating test file:", err)
		}

		fmt.Println("✅ File name:", fileName)
		fmt.Println("✅ Test file name:", testFileName)
	},
}

func init() {
	rootCmd.AddCommand(genCmd)
}

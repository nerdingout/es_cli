package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var genCmd = &cobra.Command{
	Use:   "rhf",
	Short: "A command to generate a helper function with test.",
	Long: `
Can use the gen command to generate starter / boilerplate code for a helper function.
The helper function and test will be created in the folder you run the command in.

Example:

es rhf "countVowels"

Creates an index.js with boilerplate code for a countVowels function, 
and a test file that passes the boilerplate.
`,
	Run: func(cmd *cobra.Command, args []string) {
		functionName := args[0]
		fileName := functionName + ".js"
		testFileName := functionName + ".test.js"

		fileContents := fmt.Sprintf(`/**
 * @function %s
 * @param {string} arg
 * @returns {string}
 */

const %s = (arg) => {
  return arg;
};

export default %s;
`, functionName, functionName, functionName)

		testFileContents := fmt.Sprintf(`import %s from ".";

describe("%s", () => {
  test("tests function", () => {
    expect(%s("hello")).toBe("hello");
  });
});
`, functionName, fileName, functionName)

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

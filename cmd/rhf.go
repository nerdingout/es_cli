package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "rhf",
	Short: "A command to generate a helper function.",
	Long: `
Can use the gen command to generate starter / boilerplate code.

Example:

es rhf "countVowels"

Creates an index.js with boilerplate code for a countVowels function, 
and a test file that passes the boilerplate.
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gen called", args)
		functionName := args[0]
		fileName := functionName + ".js"
		testFileName := functionName + ".test.js"

		fileContents := fmt.Sprintf(`
const %s = (string) => {
	console.log(string)
}

export default %s
`, functionName, functionName)

		testFileContents := fmt.Sprintf(`
import %s from "./"

describe("%s", () => {
  test("tests function", () => {
    expect($s("hello")).toBe("hello");
  });
});
`, functionName, fileName)

		fmt.Println("fileName:", fileName)
		fmt.Println("testFileName:", testFileName)

		if err := os.WriteFile("index.js", []byte(fileContents), 0644); err != nil {
			fmt.Println(err)
		}

		if err := os.WriteFile(testFileName, []byte(testFileContents), 0644); err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(genCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// genCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// genCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

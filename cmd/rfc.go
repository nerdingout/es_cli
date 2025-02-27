package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

func CountStoriesDepth(s string) int {
	parts := strings.Split(s, "/")
	for i, part := range parts {
		if part == "src" {
			return len(parts) - i
		}
	}
	return 0
}

func AddDots(depth int) string {
	var result string

	for i := 1; i < depth; i++ {
		result += "../"
	}

	return fmt.Sprintf("%sstories/tests/setupTests", result)
}

var rfcCmd = &cobra.Command{
	Use:   "rfc",
	Short: "react: Creates a React component, test file, storybook, and presentation hook",
	Long: `Creates a React component, Jest test file, storybook, and presentation hook

example:

es rfc NewComponent
`,
	Run: func(cmd *cobra.Command, args []string) {
		componentName := args[0]
		presentationHookFileName := fmt.Sprintf("use%sData.js", componentName)
		storyFilename := fmt.Sprintf("%s.stories.jsx", componentName)
		testFileName := fmt.Sprintf("%s.test.js", componentName)

		componentContents := fmt.Sprintf(`import React from "react";
import PropTypes from "prop-types";
import use%sData from "./use%sData";

/**
 *
 * @param props
 * @returns {Element}
 * @constructor
 */
const %s = ({ text }) => {
  const { incrementCount, count } = use%sData();

  return (
    <div>
      <p>{text}</p>
      <p>{count}</p>
      <button type="button" onClick={incrementCount}>
        Increment
      </button>
    </div>
  );
};

%s.propTypes = {
  text: PropTypes.string.isRequired,
};

export default %s;
`, componentName, componentName, componentName, componentName, componentName, componentName)

		presentationHookContents := fmt.Sprintf(`import { useCallback, useState } from "react";

const use%sData = () => {
  const [count, setCount] = useState(0);

  const incrementCount = useCallback(() => {
    setCount(count + 1);
  }, [count, setCount]);

  return { incrementCount, count };
};

export default use%sData;
`, componentName, componentName)

		storyContents := fmt.Sprintf(`import React from "react";
import %s from ".";

export default {
  title: "Tests/%s",
  component: %s,
  args: {
    text: "Test Component",
  },
};

const Template = (args) => <%s {...args} />;

export const Default = Template.bind({});
`, componentName, componentName, componentName, componentName)

		getPathCmd := exec.Command("pwd")
		pathCmdOutput, err := getPathCmd.Output()

		if err != nil {
			fmt.Println("❌ Cannot get path:", err)
			return
		}

		storiesDepth := CountStoriesDepth(string(pathCmdOutput))

		testFileContents := fmt.Sprintf(`/* eslint-disable react/jsx-filename-extension */

import React from "react";
import { render } from "%s";

import %s from ".";

describe("<%s />", () => {
  it("snapshot of %s", () => {
    const tree = render(<%s text="test text" />);
    expect(tree).toMatchSnapshot();
  });
});
`, AddDots(storiesDepth), componentName, componentName, componentName, componentName)

		if err := os.WriteFile("index.jsx", []byte(componentContents), 0644); err != nil {
			fmt.Println("❌ Error creating component file:", err)
		}

		if err := os.WriteFile(presentationHookFileName, []byte(presentationHookContents), 0644); err != nil {
			fmt.Println("❌ Error creating presentation hook file:", err)
		}

		if err := os.WriteFile(storyFilename, []byte(storyContents), 0644); err != nil {
			fmt.Println("❌ Error creating component story file:", err)
		}

		if err := os.WriteFile(testFileName, []byte(testFileContents), 0644); err != nil {
			fmt.Println("❌ Error creating component test file:", err)
		}

		fmt.Println("✅ Component with story and presentation hook created")
	},
}

func init() {
	rootCmd.AddCommand(rfcCmd)
}

/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rfcCmd represents the rfc command
var rfcCmd = &cobra.Command{
	Use:   "rfc",
	Short: "Creates a React component, test file, storybook, and presentation hook",
	Long: `Creates a React component, Jest test file, storybook, and presentation hook

example:

es rfc NewComponent
`,
	Run: func(cmd *cobra.Command, args []string) {
		componentName := args[0]
		presentationHookFileName := fmt.Sprintf("use%sData.js", componentName)
		storyFilename := fmt.Sprintf("%s.stories.js", componentName)

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

		if err := os.WriteFile("index.jsx", []byte(componentContents), 0644); err != nil {
			fmt.Println("❌ Error creating component file:", err)
		}

		if err := os.WriteFile(presentationHookFileName, []byte(presentationHookContents), 0644); err != nil {
			fmt.Println("❌ Error creating component file:", err)
		}

		if err := os.WriteFile(storyFilename, []byte(storyContents), 0644); err != nil {
			fmt.Println("❌ Error creating component file:", err)
		}

		fmt.Println("✅ Component with story and presentation hook created")
	},
}

func init() {
	rootCmd.AddCommand(rfcCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rfcCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rfcCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

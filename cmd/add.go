package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addTask)
}

var addTask = &cobra.Command{
	Use:   "add",
	Short: "Add a new task to your TODO list",
	Long:  `Add a new task to your TODO list`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		inputTask := strings.Join(args, " ")
		fmt.Println(inputTask)
	},
}

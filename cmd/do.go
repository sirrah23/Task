package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(doTask)
}

var doTask = &cobra.Command{
	Use:   "do",
	Short: "Mark a task on your TODO list as complete",
	Long:  `Mark a task on your TODO list as complete`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		inputTaskId, _ := strconv.ParseInt(args[0], 10, 64)
		fmt.Println(inputTaskId)
	},
}

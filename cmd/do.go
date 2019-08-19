package cmd

import (
	"log"
	"strconv"

	"github.com/sirrah23/task/db"
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
		inputTaskID, err := strconv.ParseInt(args[0], 10, 0)
		if err != nil {
			log.Fatal(err)
			return
		}

		if inputTaskID < 1 {
			log.Fatalf("Invalid task ID %d", inputTaskID)
			return
		}

		d, err := db.NewConnection()
		if err != nil {
			log.Fatal(err)
			return
		}
		defer d.Close()
		err = d.DeleteTask(int(inputTaskID - 1))
		if err != nil {
			log.Fatal(err)
			return
		}
	},
}

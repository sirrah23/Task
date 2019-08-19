package cmd

import (
	"log"
	"strings"

	"github.com/sirrah23/task/db"
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
		d, err := db.NewConnection()
		if err != nil {
			log.Fatal(err)
			return
		}
		defer d.Close()
		err = d.AddTask(inputTask)
		if err != nil {
			log.Fatal(err)
			return
		}
	},
}

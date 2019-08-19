package cmd

import (
	"fmt"
	"log"

	"github.com/sirrah23/task/db"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listTasks)
}

var listTasks = &cobra.Command{
	Use:   "list",
	Short: "List all of your incomplete tasks",
	Long:  `List all of your incomplete tasks`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		d, err := db.NewConnection()
		if err != nil {
			log.Fatal(err)
			return
		}
		defer d.Close()
		tasks, err := d.ListTasks()
		if err != nil {
			log.Fatal(err)
			return
		}
		for idx, task := range tasks {
			fmt.Printf("%d. %s\n", idx+1, task)
		}
	},
}

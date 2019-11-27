package cmd

import (
	"codes/clitask/db"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

//Gettask ..
var Gettask = db.AllTasks
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all your Tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := Gettask()
		if err != nil {
			fmt.Println("Something went wrong")
			os.Exit(1)
		}
		if len(tasks) == 0 {
			fmt.Println("You have no tasks pending")
			return
		}
		fmt.Println("You have the following tasks pending")
		for i, task := range tasks {
			fmt.Printf("%d. %s \n", i+1, task.Value)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}

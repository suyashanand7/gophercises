package cmd

import (
	"codes/clitask/db"
	"fmt"

	"github.com/spf13/cobra"

	"strconv"
)

//Delfunc ..
var Delfunc = db.DeleteTask

//Alltask ..
var Alltask = db.AllTasks

//Deletestatus ..
var Deletestatus = false
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to get the ID of the task")
			} else {
				ids = append(ids, id)
			}
		}
		tasks, err := Alltask()
		if err != nil {
			fmt.Println("Oops! Something went wrong")
			return
		}
		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Println("Invalid Task Number :", id)
				continue
			}
			task := tasks[id-1]
			err := Delfunc(task.Key)
			if err != nil {
				fmt.Println("Failed to remove the task", id)
			} else {
				fmt.Println("Completed the task : ", id)
				Deletestatus = true
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(deleteCmd)
}

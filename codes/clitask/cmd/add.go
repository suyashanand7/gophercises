package cmd

import (
	"codes/clitask/db"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

//Createtaskfunc ..
var Createtaskfunc = db.CreateTask

//Statusrecord ...
var Statusrecord = false

//Task ..
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a New Task",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, "")
		_, err := Createtaskfunc(task)
		if err != nil {
			fmt.Println("Failed to add task to the list")
			return
		}
		fmt.Printf("Added \"%s\" to your list. \n", task)
		Statusrecord = true
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}

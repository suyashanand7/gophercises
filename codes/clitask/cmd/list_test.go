package cmd

import (
	"codes/clitask/db"
	"testing"
)

func TestList(t *testing.T) {
	Gettask = func() ([]db.Task, error) {
		var abc []db.Task
		abc = append(abc, db.Task{
			Key:   1,
			Value: "Dummy Task",
		})
		return abc, nil
	}
	listCmd.Run(listCmd, []string{""})
	
}

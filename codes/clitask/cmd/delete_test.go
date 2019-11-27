package cmd

import (
	"codes/clitask/db"
	"errors"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestDelete(t *testing.T) {
	Alltask = func() ([]db.Task, error) {
		taskdemo := []db.Task{
			{Key: 1, Value: "DemoTask"},
		}
		return taskdemo, nil
	}
	Delfunc = func(key int) error {
		if key < 0 {
			return errors.New("Invalid Key")
		}
		return nil
	}
	deleteCmd.Run(deleteCmd, []string{"1"})
	assert.Equal(t, Deletestatus, true)
}

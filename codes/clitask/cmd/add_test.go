package cmd

import (
	"errors"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestAdd(t *testing.T) {
	Createtaskfunc = func(task string) (id int, err error) {
		if task == "" {
			return -1, errors.New("FAILED")
		}
		return 0, nil
	}
	addCmd.Run(addCmd, []string{"DummyTask"})
	assert.Equal(t, Statusrecord, true)
}

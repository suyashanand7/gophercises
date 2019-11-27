package db

import (
	"path/filepath"
	"testing"

	"github.com/magiconair/properties/assert"
	"github.com/mitchellh/go-homedir"
)

func TestInit(t *testing.T) {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "demo.db")
	err := Init(dbPath)
	flag := (err == nil)
	db.Close()
	assert.Equal(t, flag, true)
}

func TestCreate(t *testing.T) {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "demo.db")
	Init(dbPath)
	task := "dummy task"
	_, err := CreateTask(task)
	flag := (err == nil)
	db.Close()
	assert.Equal(t, flag, true)
}

func TestAllTasks(t *testing.T) {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "demo.db")
	Init(dbPath)
	task := "dummy task"
	abcd, _ := CreateTask(task)
	_ = abcd
	_, errrn := AllTasks()
	if errrn == nil {
		t.Logf("Pass")
		return
	}
	t.Error("Fail")
	db.Close()
}

func TestDeleteTask(t *testing.T) {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "demo.db")
	Init(dbPath)
	task := "dummy task"
	CreateTask(task)
	err := DeleteTask(1)
	flag := (err == nil)
	assert.Equal(t, flag, true)
}

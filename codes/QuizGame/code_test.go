package main

import (
	"flag"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	flag.Parse()
	os.Exit(code)
}

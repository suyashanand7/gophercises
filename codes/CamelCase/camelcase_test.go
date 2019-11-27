package main

import (
	"fmt"
	"testing"
)

func TestCount(t *testing.T) {
	cnt := count("iAmSuyash")
	if cnt != 3 {
		t.Error()
	} else {
		fmt.Println(cnt)
	}
}

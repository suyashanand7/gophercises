package main

import (
	"fmt"
	"testing"
)

func TestCipher(t *testing.T) {
	testvar := ciph("ABC", 1)
	if len(testvar) == 0 {
		t.Error("Enter a valid string")
	} else if testvar != "BCD" {
		fmt.Println("Test String = ABC  and  Rotate = 1")
		t.Errorf("FAILED --- Expected -  %v  Got - %v", "BCD", testvar)
	} else {
		t.Logf("SUCCESS --- Expected -  %v  Got - %v", "BCD", testvar)
	}
}
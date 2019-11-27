package parse

import (
	"testing"
)

//TestParse ..
func TestParse(t *testing.T) {
	xyz := [][]string{{"2+2", "4"}}
	abc := GiveLines(xyz)
	for _, line := range abc {
		if line.Q == "2+2" && line.A == "4" {
			t.Logf("PARSED")
			return
		}
		t.Errorf("Failed to parse")
	}
}

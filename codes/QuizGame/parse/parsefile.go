package parse

//Prob ..
type Prob struct {
	Q string
	A string
}

//GiveLines ..
func GiveLines(lines [][]string) []Prob {
	abc := make([]Prob, len(lines))
	for i, line := range lines {
		abc[i] = Prob{
			Q: line[0],
			A: line[1],
		}
	}
	return abc
}

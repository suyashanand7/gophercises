package main

import (
	"codes/QuizGame/parse"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	filename := flag.String("csv", "problems.csv", "Enter a CSV file to start the quiz")
	quiztime := flag.Int("time", 7, "Enter time for quiz in seconds")
	flag.Parse()
	file, err := os.Open(*filename)
	if err != nil {
		fmt.Println("Failed to Load the CSV File")
		os.Exit(1)
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		fmt.Println("Failed to parse the file")
		os.Exit(1)
	}
	problems := parse.GiveLines(lines)
	corr := 0
	timer := time.NewTimer(time.Duration(*quiztime) * time.Second)
	for i, line := range problems {
		fmt.Printf("Problem %d - %s = \n", i+1, line.Q)
		anschan := make(chan string)
		go func() {
			answer := ""
			fmt.Scanf("%s\n", &answer)
			anschan <- answer
		}()
		select {
		case <-timer.C:
			fmt.Printf("You scored %d out of %d \n", corr, len(problems))
			return
		case ans := <-anschan:
			if ans == line.A {
				corr++
			}
		}
	}
}

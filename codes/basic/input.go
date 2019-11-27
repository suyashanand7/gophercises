package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	abc := flag.String("flag1", "default.csv", "Enter File")
	flag.Parse()
	file, err := os.Open(*abc)
	if err != nil {
		fmt.Println("Failed to open: ", *abc)
		os.Exit(1)
	}
	_ = file
}

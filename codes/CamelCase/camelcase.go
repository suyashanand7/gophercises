package main

import (
	"fmt"
)

func count(str string) int {
	a := len(str)
	count := 1
	for i := 0; i < a; i++ {
		abc := (byte)(str[i])
		if abc >= 64 && abc <= 90 {
			count++
		}
	}
	return (count)
}
func main() {
	str := ""
	fmt.Println("Enter the string to count words")
	fmt.Scanln(&str)
	a := count(str)
	fmt.Println(a)
}

package main

import "fmt"

func fibonacci() func() int {
	last1 := 0
	last2 := 1
	var t int
	return func() int {
		t = last1
		last1 = last2
		last2 = last1 + t
		return t
	}
}
func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

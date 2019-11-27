package main

import (
	"fmt"
)

func ciph(str string, n int) string {
	newstr := ""
	for i := 0; i < len(str); i++ {
		ch := (byte)(str[i])
		var c byte = 0
		if ch <= 122 {
			if ch <= 90 {
				c = byte(ch - 65)
				c = (c + byte(n)) % 26
				c = c + 65
			} else {
				c = ch - 97
				c = (c + byte(n)) % 26
				c = c + 97
			}
		}
		newstr = newstr + string(c)
	}
	return (newstr)
}

func main() {
	fmt.Println("Enter a string to rotate")
	var str string
	fmt.Scanln(&str)
	fmt.Println("Enter the rotate count")
	var n int
	fmt.Scanln(&n)
	res := ciph(str, n)
	fmt.Println(res)
}

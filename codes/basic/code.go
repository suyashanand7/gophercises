package main

import "fmt"

func swap(a,b string)(string,string)  {
	return b,a
}
func main()  {
	a,b := swap("HI","Suyash")
	fmt.Println(a,b)
}
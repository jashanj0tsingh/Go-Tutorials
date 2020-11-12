package main

import "fmt"

func reverseString(str string) string {

	// convert string to a slice of runes
	cs := []rune(str)
	end := len(cs)
	for i := 0; i < end; i++ {
		cs[i], cs[end-1] = cs[end-1], cs[i]
		end--
	}
	return string(cs)
}

func main() {
	fmt.Println(reverseString("golang"))
}
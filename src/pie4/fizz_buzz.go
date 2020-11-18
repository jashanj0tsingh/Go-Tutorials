package main

import "fmt"

func isMultiple(i, j int) bool {
	return i%j == 0
}
func main() {
	fizzbuzz := make(map[int]string)
	fizzbuzz[3] = "fizz"
	fizzbuzz[5] = "buzz"
	for i := 1; i < 101; i++ {
		str := ""
		for k, v := range fizzbuzz {
			if isMultiple(i, k) == true {
				str += v
			}
		}
		fmt.Println(i, ":", str)
	}
}

// eof

package main

import "fmt"

// Given an unsorted list of integers, write a function that returns a new list
// with all duplicate values removed.

// input list {2,1,4,5,6,2,1,9,8,5}
// output list {2,1,4,5,6,9,8}

type list []int

func printList(l list) {
	for _, v := range l {
		fmt.Print(v, " ")
	}
}

func isPresentIn(value int, l list) bool {
	for _, v := range l {
		if v == value {
			return true
		} else {
			continue
		}
	}
	return false
}

func removeDuplicates(listWithDuplicates list) list {
	listWithUniques := make(list, 0)
	for _, v := range listWithDuplicates {
		if isPresentIn(v, listWithUniques) {
			continue
		} else {
			listWithUniques = append(listWithUniques, v)
		}
	}
	return listWithUniques
}

func main() {
	var l = list{2, 1, 4, 5, 6, 2, 1, 9, 8, 5}
	foo := removeDuplicates(l)
	printList(foo)
}

package main

import "fmt"

var a = []int{10, 3, 240, 34, 100, 2}

func insertSort(a []int) []int {
	for j := 1; j < len(a); j++ {
		key := a[j]
		i := j - 1
		for (i >= 0) && (a[i] > key) {
			a[i+1] = a[i]
			i -= 1
		}
		a[i+1] = key
	}
	return a
}

func main() {
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", insertSort(a))
}

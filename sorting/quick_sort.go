package Sorting

import (
	"math/rand"
)

func quickSort(n []int) []int {
	if len(n) < 2 {
		return n
	}
	left, right := 0, len(n) - 1
	pivot := rand.Int() % len(n)

	n[pivot], n[right] = n[right], n[pivot]
	for i, _ := range n {
		if n[i] < n[right] {
			n[left], n[i] = n[i], n[left]
			left++
		}
	}
	n[left], n[right] = n[right],n[left]
	quickSort(n[:left])
	quickSort(n[left+1:])

	return n
}

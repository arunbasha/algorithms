package Sorting

func insertionSort(n []int) []int {
	len := len(n)
	for i:=1; i<len; i++ {
		tmp := n[i]
		j := i
		for j = i; j > 0 && n[j-1] > tmp; j-- {
			n[j] = n[j-1]
		}
		n[j] = tmp
	}
	return n
}
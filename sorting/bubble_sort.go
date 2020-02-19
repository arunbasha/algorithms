package Sorting

func bubbleSort(n []int) []int {
	len := len(n)
	for i:=0; i<len - 1; i++ {
		for j:= i + 1; j<len; j++ {
			if n[i] > n[j] {
				n[i], n[j] = n[j], n[i]
			}
		}
	}
	return n
}
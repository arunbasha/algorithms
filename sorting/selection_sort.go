package Sorting

func selectionSort(n []int) []int {
	len := len(n)
	for i:=0; i<len; i++ {
		minIndex := i
		for j:= i+1; j<len; j++ {
			if n[j] < n[minIndex] {
				minIndex = j
			}
		}
		n[i], n[minIndex] = n[minIndex], n[i]
	}
	return n
}

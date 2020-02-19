package Algorithm

func reverseIntSlice(n []int) []int {
	for i, j := 0, len(n)-1; i <j; i,j = i+1, j-1 {
		n[i], n[j] = n[j], n[i]
	}
	return n
}


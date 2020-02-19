package Sorting

func mergeSort(n []int) []int {
	len := len(n)
	if len == 1 {
		return  n
	}

	mid := len / 2

	var (
		left = make([]int, mid)
		right = make([]int, len-mid)
	)
	for i := 0; i<len; i++{
		if i< mid {
			left[i] = n[i]
		} else {
			right[i-mid] = n[i]
		}
	}

	return merge(mergeSort(left), mergeSort(right))
}


func merge(left, right []int) []int {
	res := make([]int, len(left) + len(right))
	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			res[i] = left[0]
			left = left[1:]
		} else {
			res[i] = right[0]
			right = right[1:]
		}
		i++
	}
	for j:=0; j< len(left); j++ {
		res[i] = left[j]
		i++
	}

	for j:=0; j<len(right); j++{
		res[i] = right[j]
		i++
	}
	return 	res
}

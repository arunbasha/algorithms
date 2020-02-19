package Algorithm

func Sum2InList(list []int, sum int) (int, int){

	for index, i := range  list {
		for j := index + 1; j< len(list) ; j++ {
			if i + list[j] == sum {
				return i, list[j]
			}
		}
	}
	return 0, 0
}

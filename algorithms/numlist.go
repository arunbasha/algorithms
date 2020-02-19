package Algorithm

func NumInList(list []int, num int) bool {

	//if list != nil {
	//	return false
	//}

	for _, val := range list {
		if val == num {
			return true
		}
	}
	return false
}
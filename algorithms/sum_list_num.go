package Algorithm

import "fmt"

func SumListNum(list []int) (int, error) {
	if list == nil {
		return 0, fmt.Errorf("Invalid list")
	}
	sum := 0
	for _, val :=  range list {
		sum = sum + val
	}
	return sum, nil
}
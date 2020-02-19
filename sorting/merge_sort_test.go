package Sorting

import (
	"fmt"
	"testing"
)


func TestMergeSort(t *testing.T) {
	tests := []struct{
		input []int
		output []int
	}{
		{[]int{1,4,3,4,5,3,7,8,9}, []int{1,3,3,4,4,5,7,8,9}},
		{[]int{9,8,7,6,5,4,3,2,1,0}, []int{0,1,2,3,4,5,6,7,8,9}},
		{[]int{7, 5, 2, 4, 3, 9}, []int{2,3,4,5,7,9}},
		{[]int{29, 64, 73, 34, 20}, []int{20,29,34,64,73}},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("BubbleSort_%v", test.input), func(t *testing.T) {
			res:= mergeSort(test.input)
			i := 0
			for _, v := range res {
				if v != test.output[i] {
					t.Fatalf("expected %v, got %v", test.output, res)
					return
				}
				i++
			}
		})
	}
}

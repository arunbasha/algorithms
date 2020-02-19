package Algorithm

import (
	"fmt"
	"testing"
)

func TestSum2InList(t *testing.T) {
	tests := []struct{
		list []int
		sum int
		output int
		output2 int
	} {
		{[]int{1,2,3,4,5,6}, 4, 1,3},
		{[]int{1,2,3,4,5,6}, 7, 1,6},
		{[]int{1,2,3,4,5,6}, 10, 4,6},
		{[]int{1,2,3,4,5,6}, 6, 1,5},
		{[]int{1,2,3,4,5,6}, 3, 1,2},
		{[]int{1,2,3,4,5,6}, 11, 5,6},
		{[]int{1,2,3,4,5,6}, 100, 0,0},
		{[]int{0, -1, 1}, 0, -1,1},

	}

	for _,test := range tests {
		t.Run(fmt.Sprintf("Sum2InList(%v %d)", test.list, test.sum), func(t *testing.T){
			if val1, val2 := Sum2InList(test.list, test.sum); val1 != test.output && val2 != test.output2 {
				t.Fatalf("Expected (%d, %d), got (%d,%d)", test.output, test.output2, val1, val2)
			}
		})
	}
}
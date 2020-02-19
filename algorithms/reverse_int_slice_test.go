package Algorithm

import (
	"fmt"
	"testing"
)

func TestReverseInSlice(t *testing.T) {
	tests := []struct {
		input []int
		output []int
	}{
		{[]int{1,2,3,4},[]int{4,3,2,1}},
		{[]int{34,23,213,4,5,76},[]int{76,5,4,213,23,34}},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("ReverseIntSlice_%v", test.input), func(t *testing.T) {
			res:=reverseIntSlice(test.input)
			for i, v := range res {
				if v != test.output[i] {
					t.Fatalf("Expected, %+v, got %+v", test.output, res)
				}
			}
		})
	}
}

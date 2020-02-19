package Algorithm

import (
	"fmt"
	"testing"
)

func TestFactor(t *testing.T) {
	tests := []struct{
		primes []int
		sum int
		res []int
	}{
		{[]int{2,3,5}, 10, []int{2,5}},
		{[]int{2,3,5,7}, 70, []int{2,5,7}},
		{[]int{2,3,5,7}, 14, []int{2,7}},
		{[]int{2,3,5,7,13,17,23,29}, 8671, []int{13,23,29}},
		{[]int{2,3,5,7,13,17,23,29,31,37}, 9945637, []int{13,23,29, 31, 37}},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Factorial_(%v,%d)", test.primes,test.sum), func(t *testing.T){
			res := factor(test.primes, test.sum)
			for i, val  := range test.res {
				if res[i] != val{
					t.Fatalf("expect %v, got %v", test.res, res)
				}
			}
		})
	}
}
package Algorithm

import (
	"fmt"
	"testing"
)

func TestFibonacci(t *testing.T) {
	tests := []struct{
		num int
		res int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{8, 21},
		{9, 34},
		{10, 55},
		{22, 17711},
		{41, 165580141},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Fibonacci_%d", test.num), func(t *testing.T){
			if res := fibonacci(test.num); res != test.res {
				t.Fatalf("Expected %d, got %d", test.res, res)
			}
		})
	}

}


func TestFibonacci_2(t *testing.T) {
	tests := []struct{
		num int
		res int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{8, 21},
		{9, 34},
		{10, 55},
		{22, 17711},
		{41, 165580141},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%d", test.num), func(t *testing.T){
			if res := fibonacci_2(test.num); res != test.res {
				t.Fatalf("Expected %d, got %d", test.res, res)
			}
		})
	}

}
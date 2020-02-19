package Algorithm

import (
	"fmt"
	"testing"
)

func TestGCD(t *testing.T) {
	tests := []struct {
		num1, num2 int
		output int
	}{
		{10,5,5},
		{12, 18,6},
		{-4,14,2},
		{9,28,1},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Gcd_(%d,%d)", test.num1,test.num2), func(t *testing.T) {
			if res := GCD(test.num1, test.num2); res != test.output {
				t.Fatalf("expected %d, got %d", test.output, res)
			}
		})
	}
}

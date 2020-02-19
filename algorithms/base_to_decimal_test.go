package Algorithm

import (
	"fmt"
	"testing"
)

func TestBaseToDec(t *testing.T) {
	tests := []struct{
		value string
		base int
		output int
	}{
		{"1", 1, 1},
		{"10", 2, 2},
		{"11", 2, 3},
		{"11",  3, 4},
		{"12", 3, 5},
		{"21", 3, 7},
		{"12345", 10, 12345},
		{"A",16, 10},
		{"B",16, 11},
		{"C",16, 12,},
		{"D",16, 13},
		{"E",16, 14},
		{"F",16, 15},
		{"10",16, 16},
		{"11",16, 17},
		{"12",16, 18},
		{"1F",16, 31},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("BaseToDec_(%s,%d)_%d", test.value, test.base, i+1), func(t *testing.T) {
			if res := BaseToDec(test.value, test.base); res != test.output {
				t.Fatalf("Expected %d, got %d", test.output, res)
			}
		})
	}

}

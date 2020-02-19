package Algorithm

import (
	"fmt"
	"testing"
)

func TestDecToBase (t *testing.T) {
	tests := []struct{
		dec, base int
		output string
	} {
		{1, 2, "1"},
		{2, 2, "10"},
		{3, 2, "11"},
		{4, 3, "11"},
		{5, 3, "12"},
		{7, 3, "21"},
		{12345, 10, "12345"},
		{10, 16, "A"},
		{11, 16, "B"},
		{12, 16, "C"},
		{13, 16, "D"},
		{14, 16, "E"},
		{15, 16, "F"},
		{16, 16, "10"},
		{17, 16, "11"},
		{18, 16, "12"},
		{31, 16, "1F"},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("DecToBase_(%d,%d)", test.dec, test.base), func (t *testing.T){
			if res := decToBase(test.dec, test.base); res != test.output {
				t.Fatalf("expected %s, got %s", test.output, res)
			}
		} )
	}
}

func TestDecToBase_2(t *testing.T) {
	tests := []struct{
		dec, base int
		output string
	} {
		{1, 2, "1"},
		{2, 2, "10"},
		{3, 2, "11"},
		{4, 3, "11"},
		{5, 3, "12"},
		{7, 3, "21"},
		{12345, 10, "12345"},
		{10, 16, "A"},
		{11, 16, "B"},
		{12, 16, "C"},
		{13, 16, "D"},
		{14, 16, "E"},
		{15, 16, "F"},
		{16, 16, "10"},
		{17, 16, "11"},
		{18, 16, "12"},
		{31, 16, "1F"},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("DecToBase_(%d,%d)", test.dec, test.base), func (t *testing.T){
			if res := decToBase_2(test.dec, test.base); res != test.output {
				t.Fatalf("Expected %s, got %s", test.output, res)
			}
		} )
	}
}

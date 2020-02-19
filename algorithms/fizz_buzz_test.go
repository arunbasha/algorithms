package Algorithm

import (
	"testing"
)

func TestFizzBuzz( t *testing.T) {
	tests := []struct {
		input int
		output string
	} {
		{10,"Buzz"},
		{9,"Fizz"},
		{15,"Fizz Buzz"},
		{300,"Fizz Buzz"},
		{99,"Fizz"},
		{1000,"Buzz"},
	}
	
	for _, test := range tests {
		t.Run("test", func(t *testing.T) {
			if res := fizzBuzz(test.input); res != test.output {
				t.Fatalf("expected %s, got %s", test.output, res)
			}
		})
	}
}
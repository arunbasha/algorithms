package Algorithm

import (
	"fmt"
	"testing"
)

func TestCheckPoint(t *testing.T) {
	tests := []struct{
		workers  []string
		numAsseblies int
		output bool
	} {
		{[]string{"a", "b", "c", "d", "e"}, 5, true},
		{[]string{"a", "b", "c", "d", "e"}, 10, true},
		{[]string{"a"}, 5, true},
		{[]string{"a", "b"}, 10,true},
		{[]string{"a", "b", "c", "d", "e", "f", "g"}, 2,true},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("CheckPoint_%v", test.workers), func(t *testing.T) {
			if res := checkPoint(test.workers, test.numAsseblies); res != test.output {
				t.Fatalf("expected %v, got %v", res, test.output)
			}
		})
	}
}

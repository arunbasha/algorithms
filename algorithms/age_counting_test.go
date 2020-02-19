package Algorithm

import (
	"fmt"
	"testing"
	"time"
)

func TestAgeCounting(t *testing.T) {
	tests := []struct {
		input time.Time
		output int
	}{
		{time.Date(2010, 3, 14, 0, 0, 0, 0, time.UTC), 10},
		{time.Date(2009, 3, 14, 0, 0, 0, 0, time.UTC), 11},
		{time.Date(2005, 5, 12, 0, 0, 0, 0, time.UTC), 15},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("AgeTesting_%d", i + +1), func(t *testing.T) {
			if v := ageCounting(test.input); v != test.output {
				t.Fatalf("expected %d, got %d", test.output, v)
			}
		})
	}
}

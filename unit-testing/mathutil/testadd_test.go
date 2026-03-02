package mathutil

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	test := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positive", 2, 3, 5},
		{"zero", 0, 5, 5},
		{"negatiive", -1, -1, -2},
	}

	for i, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			fmt.Printf("start =  %d name = %v", i, tt.name)

			result := Add(tt.a, tt.b)

			fmt.Printf("end =  %d name = %v", i, tt.name)

			if result != tt.expected {
				t.Fatalf("expect %d got %d", tt.expected, result)
			}
		})
	}

}

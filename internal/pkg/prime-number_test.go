package pkg

import (
	"testing"
)

func TestHighestPrime(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  int
	}{
		{name: "1. input lower than 0", input: -1, want: 0},
		{name: "2. input = 0", input: 0, want: 0},
		{name: "3. input lower than 3", input: 2, want: 0},
		{name: "4. input is prime number", input: 5, want: 3},
		{name: "5. input = 10", input: 10, want: 7},
		{name: "5. input = 100", input: 100, want: 97},
		{name: "6. input is very large number", input: 1000000000, want: 999999937},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HighestPrime(tt.input); got != tt.want {
				t.Errorf("HighestPrime(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

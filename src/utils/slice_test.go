package utils

import (
	"testing"
)

func TestContains(t *testing.T) {
	tests := []struct {
		name       string
		sl         []string
		to_contain string
		result     bool
	}{
		{
			name:       "When slice does contain element, then returns True.",
			sl:         []string{"one", "two", "three"},
			to_contain: "two",
			result:     true,
		},
		{
			name:       "When slice does not contain element, then returns false.",
			sl:         []string{"one", "two", "three"},
			to_contain: "four",
			result:     false,
		},
		{
			name:       "When slice is empty, then returns false.",
			sl:         []string{},
			to_contain: "won't be checked",
			result:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contains(tt.sl, tt.to_contain); got != tt.result {
				t.Errorf("Contains() = %v, want %v", got, tt.result)
			}
		})
	}
}

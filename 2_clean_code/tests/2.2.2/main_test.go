package main

import (
	"testing"
)

func Sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func TestSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"zero", []int{}, 0},
		{"one", []int{1}, 1},
		{"two", []int{1, 2}, 3},
		{"many", []int{1, 2, 3, 4, 5}, 15},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Sum(test.nums...)
			if got != test.want {
				t.Errorf("%s: got %d, want %d", test.name, got, test.want)
			}
		})
	}
}

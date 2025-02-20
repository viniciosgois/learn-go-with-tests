package main

import "testing"

func TestSum(t *testing.T) {
	t.Run("should return the sum of slice with any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		output := Sum(numbers)
		expected := 6

		if output != expected {
			t.Errorf("got %d expected %d given, %v", output, expected, numbers)
		}
	})
}

package main

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	input := []int{1, 2, 3}
	output := Sum(input)
	expected := 6

	if output != expected {
		t.Errorf("got %d expected %d given, %v", output, expected, input)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, expected []int) {
		t.Helper()
		if !slices.Equal(got, expected) {
			t.Errorf(" got %v expected %v", got, expected)
		}
	}

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		expected := []int{2, 9}
		checkSums(t, got, expected)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		expected := []int{0, 9}
		checkSums(t, got, expected)
	})
}

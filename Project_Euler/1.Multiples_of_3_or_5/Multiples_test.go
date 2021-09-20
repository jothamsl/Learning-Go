package main

import "testing"

func TestTwoSum(t *testing.T) {
	t.Run("Checking for multiples of 3 and 5 below 10", func(t *testing.T) {
		got := MultipleSum(3, 5, 10)
		want := 23

		if got != want {
			t.Errorf("Got %d, wanted %d, given 3 and 5", got, want)
		}
	})

	t.Run("Checking for multiples of 12 and 16 below 33", func(t *testing.T) {
		got := MultipleSum(12, 16, 32)
		want := 52

		if got != want {
			t.Errorf("Got %d, wanted %d, given 12 and 16", got, want)
		}
	})
}

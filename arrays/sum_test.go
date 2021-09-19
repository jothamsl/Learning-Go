package main

import "testing"

func TestSum(t *testing.T) {
	t.Run("Collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 4}

		got := Sum(numbers)
		wanted := 7

		if got != wanted {
			t.Errorf("Got %d want %d given, %v", got, wanted, numbers)
		}
	})
}

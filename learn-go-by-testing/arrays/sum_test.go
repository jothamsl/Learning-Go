package main

import (
	"reflect"
	"testing"
)

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

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %v want %v", got, want)
	}
}

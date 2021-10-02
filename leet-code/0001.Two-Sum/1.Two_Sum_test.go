package leetcode

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {

	t.Run("Sum of two digits in array that returns 9", func(t *testing.T) {
		nums := []int{2, 7, 11, 15}

		got := TwoSum(nums, 9)
		want := []int{0, 1}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %v, wanted %v, from %v", got, want, nums)
		}
	})

	t.Run("Sum of two digits in array that returns 6", func(t *testing.T) {
		nums := []int{3, 2, 4}

		got := TwoSum(nums, 6)
		want := []int{1, 2}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %v, wanted %v, from %v", got, want, nums)
		}
	})

}

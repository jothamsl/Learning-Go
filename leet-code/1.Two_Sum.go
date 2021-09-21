package leetcode

import "fmt"

func TwoSum(arr []int, val int) (res []int) {
	arrlength := len(arr)
	fmt.Println(arrlength)

	for j := arrlength - 1; j > 0; j-- {

		fmt.Println(arr[j])
	}

	for i := 0; i < arrlength; i++ {
		fmt.Println(arr[i])
	}
	return
}

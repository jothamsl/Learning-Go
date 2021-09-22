package leetcode

func TwoSum(arr []int, target int) []int {

	res := make([]int, 2)

	// Loop Through values in array
	for idx, v := range arr {
		// Using the value 'v', loop through other values of array:
		for i := 0; i < len(arr); i++ {
			// Find the difference between the target and 'v'
			diff := target - v
			// If the differece value is the same as the indexed value:
			if diff == arr[i] {
				// assign to index of array
				res[1] = idx
				res[0] = i
			}
		}
	}
	return res
}

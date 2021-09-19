package main

func Sum(array []int) (cumulator int) {
	for _, num := range array {
		cumulator += num
	}
	return cumulator
}

func SumAll(numSum ...[]int) []int { // This is a variadic function
	var sums []int
	for _, numbers := range numSum {
		sums[i] = append(sums, Sum(numbers))
	}
	return sums
}

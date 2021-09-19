package main

func Sum(array []int) (cumulator int) {
	for _, num := range array {
		cumulator += num
	}
	return cumulator
}

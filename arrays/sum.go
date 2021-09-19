package main

func Sum(array [5]int) (cumulator int) {
	for _, num := range array {
		cumulator += num
	}
	return cumulator
}

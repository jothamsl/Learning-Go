package main

func MultipleSum(x, y, z int) (sum int) {
	f := func(v int) int {
		var p = z / v
		return v * (p * (p + 1)) / 2
	}

	return f(x) + f(y) - f(x*y)
}

package main

import (
	"fmt"
	"math"
)

func main() {
	var a = "Timers"
	fmt.Println(a)

	var b, c int = 1, 3
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int 
	fmt.Println(e)

	train := "Lonely"
	fmt.Println(train)

	// Constants
	const s string = "ROBOTS IN DISGUISE"
	fmt.Println(s)

	const f = 3e20 / 500000000
	fmt.Println(f)

	fmt.Println(int64(f))
	fmt.Println(math.Sin(f))
}
package main

import "fmt"

func main() {
	i := 1
	for i < 20 {
		fmt.Println(i)
		i++
	}
	for j := 0; j <= 12; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	for i := 0; i < 4; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

 }
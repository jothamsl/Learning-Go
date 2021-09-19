package main

import (
	"fmt"
	"math/rand"
	"time"
)

var spaceLine = ""
var tripType = ""

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Printf("%-18v %-5v %-11v %v", "Spaceline", "Days", "Trip-type", "Price")
	fmt.Printf("\n==========================================\n")

	for i := 0; i < 10; i++ {
		speed := rand.Intn(15) + 16
		duration := 62100000 / speed / 86400
		price := 20.0 + speed 
		switch rand.Intn(3) + 1 {
		case 1:
			spaceLine = "SpaceX"
		case 2:
			spaceLine = "Virgin Galactic"
		case 3:
			spaceLine = "Space Adventures"
		}

		if rand.Intn(2) == 1 {
			tripType = "Round-trip"
			price = price * 2
		} else {
			tripType = "One-way"
		}


		fmt.Printf("%-18v %-5v %-11v %v %v\n", spaceLine, duration, tripType, "$", price)
	}
}

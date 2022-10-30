package main

import (
	"fmt"
	"math"
)

func main() {
	var speedHourHand float64 = 1.0 / 120.0
	var speedMinuteHand float64 = 0.1

	var hourHand float64
	var minuteHand float64

	const TWELVE_HOURS = 43200
	const TOLERANCE_VALUE = 0.09

	var hour, minute, second int

	// for each second in 12 hours
	for i := 0; i <= TWELVE_HOURS; i++ {

		// calculate hand position by multiplying speed with second
		// and reset hand after one cycle (360 degree) using math.Mod
		minuteHand = math.Mod(speedMinuteHand*float64(i), 360)
		hourHand = math.Mod(speedHourHand*float64(i), 360)

		// compare hand position of hour and minute, but with TOLERANCE_VALUE
		if hourHand-minuteHand < TOLERANCE_VALUE && minuteHand <= hourHand {
			hour = i / 3600
			minute = (i % 3600) / 60
			second = i % 60

			// fmt.Println("Hour Hand: ", hourHand)
			// fmt.Println("Minute Hand: ", minuteHand)

			fmt.Printf("Jarum jam dan menit berimpit pada jam %02d:%02d:%02d\n", hour, minute, second)
		}
	}

}

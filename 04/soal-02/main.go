package main

import (
	"fmt"
	"time"
)

func main() {
	var speedAli float64
	var speedBadu float64

	var distance = 1000.0
	var duration = 0

	var distanceAli = 0.0
	var distanceBadu = 0.0
	var durationBadu = 0

	speedAli = 2.0
	speedBadu = 3.0

	// calculate duration from A point to the B point
	for distanceAli < distance {
		duration++

		// we start to calculate durationBadu and distanceBadu after 20 seconds
		if duration > 20 {
			distanceBadu = distanceBadu + speedBadu
			durationBadu++
		}

		distanceAli = distanceAli + speedAli
		speedAli = speedAli + 0.1

	}

	fmt.Println("Durasi setelah Ali mencapai titik B adalah", duration, "detik")

	// at this moment, Ali has reached the B point and will go back
	// to the A point with a constant speed of 5m/s

	speedAli = 5.0
	distanceAli = 0.0 // reset distanceAli

	// keep going until they meet at some point in the middle of the distance
	for distanceAli+distanceBadu < distance {
		durationBadu++
		distanceAli = distanceAli + speedAli
		distanceBadu = distanceBadu + speedBadu
		// fmt.Println("Distance Ali ", distanceAli)
		// fmt.Println("Distance Badu ", distanceBadu)
		time.Sleep(500 * time.Millisecond)
	}

	fmt.Println("Ali dan Badu akan bertemu setelah detik ke", durationBadu, "sejak Badu berangkat dari titik A")

}
